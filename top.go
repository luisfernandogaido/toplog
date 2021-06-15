package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	reUp          = regexp.MustCompile(`up ((\d+) min|(\d+) days?, (([\d ]\d:\d\d)|(\d+) min)| ?(\d{1,2}:\d\d)),`)
	reLoadAverage = regexp.MustCompile(`load average: ([\d.]+), ([\d.]+), ([\d.]+)`)
	reTasks       = regexp.MustCompile(
		`Tasks:([\d ]+) total,([\d ]+) running,([\d ]+)sleeping,([\d ]+)stopped,([\d ]+)zombie`,
	)
	reCpuPercs = regexp.MustCompile(
		`%Cpu\(s\):([ .\d]+) us,([ .\d]+) sy,([ .\d]+) ni,([ .\d]+) ` +
			`id,([ .\d]+) wa,([ .\d]+) hi,([ .\d]+) si,([ .\d]+) st`,
	)
	reMem  = regexp.MustCompile(`MiB Mem :([ .\d]+) total,([ .\d]+) free,([ .\d]+) used,([ .\d]+) buff/cache`)
	reSwap = regexp.MustCompile(`MiB Swap:([ .\d]+) total,([ .\d]+) free,([ .\d]+) used\.([ .\d]+) avail Mem`)
)

type Top struct {
	Time               time.Time `json:"time"`
	UpSeconds          int       `json:"up_seconds"`
	LoadAverageOne     float64   `json:"load_average_one"`
	LoadAverageFive    float64   `json:"load_average_five"`
	LoadAverageFifteen float64   `json:"load_average_fifteen"`
	TasksTotal         int       `json:"tasks_total"`
	TasksRunning       int       `json:"tasks_running"`
	TasksSleeping      int       `json:"tasks_sleeping"`
	TasksStopped       int       `json:"tasks_stopped"`
	TasksZombie        int       `json:"tasks_zombie"`
	CpusUs             float64   `json:"cpus_us"`
	CpusSy             float64   `json:"cpus_sy"`
	CpusNi             float64   `json:"cpus_ni"`
	CpusId             float64   `json:"cpus_id"`
	CpusWa             float64   `json:"cpus_wa"`
	CpusHi             float64   `json:"cpus_hi"`
	CpusSi             float64   `json:"cpus_si"`
	CpusSt             float64   `json:"cpus_st"`
	MemTotal           float64   `json:"mem_total"`
	MemFree            float64   `json:"mem_free"`
	MemUsed            float64   `json:"mem_used"`
	MemBuffCache       float64   `json:"mem_buff_cache"`
	SwapTotal          float64   `json:"swap_total"`
	SwapFree           float64   `json:"swap_free"`
	SwapUsed           float64   `json:"swap_used"`
	SwapAvail          float64   `json:"swap_avail"`
}

func getTop(t time.Time, lines []string) (Top, error) {
	top := Top{
		Time: t,
	}
	uptime, err := getUp(lines[0])
	if err != nil {
		log.Fatal(err)
	}
	top.UpSeconds = uptime
	avgOne, avgFive, avgFifteen, err := getLoadAverage(lines[0])
	if err != nil {
		return Top{}, fmt.Errorf("get top: %w", err)
	}
	top.LoadAverageOne = avgOne
	top.LoadAverageFive = avgFive
	top.LoadAverageFifteen = avgFifteen
	tasksTotal, tasksRunning, tasksSleeping, tasksStopped, tasksZombie, err := getTasks(lines[1])
	if err != nil {
		return Top{}, fmt.Errorf("get top: %w", err)
	}
	top.TasksTotal = tasksTotal
	top.TasksRunning = tasksRunning
	top.TasksSleeping = tasksSleeping
	top.TasksStopped = tasksStopped
	top.TasksZombie = tasksZombie
	cpusUs, cpusSy, cpusNi, cpusId, cpusWa, cpusHi, cpusSi, cpusSt, err := getCpuPercs(lines[2])
	if err != nil {
		return Top{}, fmt.Errorf("get top: %w", err)
	}
	top.CpusUs = cpusUs
	top.CpusSy = cpusSy
	top.CpusNi = cpusNi
	top.CpusId = cpusId
	top.CpusWa = cpusWa
	top.CpusHi = cpusHi
	top.CpusSi = cpusSi
	top.CpusSt = cpusSt
	memTotal, memFree, memUsed, memBuffCache, err := getMem(lines[3])
	if err != nil {
		return Top{}, fmt.Errorf("get top: %w", err)
	}
	top.MemTotal = memTotal
	top.MemFree = memFree
	top.MemUsed = memUsed
	top.MemBuffCache = memBuffCache
	swapTotal, swapFree, swapUsed, swapAvail, err := getSwap(lines[4])
	if err != nil {
		return Top{}, fmt.Errorf("get top: %w", err)
	}
	top.SwapTotal = swapTotal
	top.SwapFree = swapFree
	top.SwapUsed = swapUsed
	top.SwapAvail = swapAvail
	return top, nil
}

func getUp(line string) (int, error) {
	matches := reUp.FindStringSubmatch(line)
	if len(matches) != 8 {
		return 0, fmt.Errorf("getup: len(maches) != 8 em %v", line)
	}
	if matches[2] != "" {
		fmt.Println("1")
		minutes, err := hourToMinutes("00:" + matches[2])
		if err != nil {
			return 0, fmt.Errorf("getup: %w", err)
		}
		return 60 * minutes, nil
	}
	if matches[3] != "" && matches[4] != "" && matches[6] == "" {
		fmt.Println("2")
		days, err := strconv.Atoi(matches[3])
		if err != nil {
			return 0, fmt.Errorf("getup: %w", err)
		}
		minutes, err := hourToMinutes(matches[4])
		if err != nil {
			return 0, fmt.Errorf("getup: %w", err)
		}
		return 86400*days + 60*minutes, nil
	}
	if matches[3] != "" && matches[6] != "" {
		fmt.Println("3")
		days, err := strconv.Atoi(matches[3])
		if err != nil {
			return 0, fmt.Errorf("getup: %w", err)
		}
		minutes, err := strconv.Atoi(matches[6])
		if err != nil {
			return 0, fmt.Errorf("getup: %w", err)
		}
		return 86400*days + 60*minutes, nil
	}
	if matches[7] != "" {
		fmt.Println("4")
		minutes, err := hourToMinutes(matches[4])
		if err != nil {
			return 0, fmt.Errorf("getup: %w", err)
		}
		return 60 * minutes, nil
	}
	return 0, fmt.Errorf("getup: sem matches previstos")
}

func hourToMinutes(s string) (int, error) {
	s = strings.TrimSpace(s)
	partes := strings.Split(s, ":")
	if len(partes) != 2 {
		return 0, fmt.Errorf("hourtominutes: len(partes) != 2 em %v", s)
	}
	h, err := strconv.Atoi(partes[0])
	if err != nil {
		return 0, fmt.Errorf("hourToMinutes: %w", err)
	}
	m, err := strconv.Atoi(partes[1])
	if err != nil {
		return 0, fmt.Errorf("hourToMinutes: %w", err)
	}
	return 60*h + m, nil
}

func getLoadAverage(line string) (one float64, five float64, fifteen float64, err error) {
	matches := reLoadAverage.FindStringSubmatch(line)
	if len(matches) != 4 {
		return 0, 0, 0, fmt.Errorf("get load average: len(matches) != 4")
	}
	one, err = strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("get load average: %w", err)
	}
	five, err = strconv.ParseFloat(matches[2], 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("get load average: %w", err)
	}
	fifteen, err = strconv.ParseFloat(matches[3], 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("get load average: %w", err)
	}
	return one, five, fifteen, nil
}

func getTasks(line string) (total int, running int, sleeping int, stopped int, zombie int, err error) {
	matches := reTasks.FindStringSubmatch(line)
	if len(matches) != 6 {
		return 0, 0, 0, 0, 0, fmt.Errorf("get tasks: len(matches) != 6")
	}
	total, err = strconv.Atoi(strings.TrimSpace(matches[1]))
	if err != nil {
		return 0, 0, 0, 0, 0, fmt.Errorf("get tasks: %w", err)
	}
	running, err = strconv.Atoi(strings.TrimSpace(matches[2]))
	if err != nil {
		return 0, 0, 0, 0, 0, fmt.Errorf("get tasks: %w", err)
	}
	sleeping, err = strconv.Atoi(strings.TrimSpace(matches[3]))
	if err != nil {
		return 0, 0, 0, 0, 0, fmt.Errorf("get tasks: %w", err)
	}
	stopped, err = strconv.Atoi(strings.TrimSpace(matches[4]))
	if err != nil {
		return 0, 0, 0, 0, 0, fmt.Errorf("get tasks: %w", err)
	}
	zombie, err = strconv.Atoi(strings.TrimSpace(matches[5]))
	if err != nil {
		return 0, 0, 0, 0, 0, fmt.Errorf("get tasks: %w", err)
	}
	return total, running, sleeping, stopped, zombie, err
}

func getCpuPercs(line string) (us float64, sy float64, ni float64, id float64, wa float64, hi float64, si float64, st float64, err error) {
	matches := reCpuPercs.FindStringSubmatch(line)
	if len(matches) != 9 {
		return 0, 0, 0, 0, 0, 0, 0, 0, fmt.Errorf("get cpu percs: len(matches) != 9")
	}
	us, err = strconv.ParseFloat(strings.TrimSpace(matches[1]), 64)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, fmt.Errorf("get cpu percs: %w", err)
	}
	sy, err = strconv.ParseFloat(strings.TrimSpace(matches[2]), 64)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, fmt.Errorf("get cpu percs: %w", err)
	}
	ni, err = strconv.ParseFloat(strings.TrimSpace(matches[3]), 64)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, fmt.Errorf("get cpu percs: %w", err)
	}
	id, err = strconv.ParseFloat(strings.TrimSpace(matches[4]), 64)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, fmt.Errorf("get cpu percs: %w", err)
	}
	wa, err = strconv.ParseFloat(strings.TrimSpace(matches[5]), 64)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, fmt.Errorf("get cpu percs: %w", err)
	}
	hi, err = strconv.ParseFloat(strings.TrimSpace(matches[6]), 64)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, fmt.Errorf("get cpu percs: %w", err)
	}
	si, err = strconv.ParseFloat(strings.TrimSpace(matches[7]), 64)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, fmt.Errorf("get cpu percs: %w", err)
	}
	st, err = strconv.ParseFloat(strings.TrimSpace(matches[8]), 64)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, fmt.Errorf("get cpu percs: %w", err)
	}
	return us, sy, ni, id, wa, hi, si, st, nil
}

func getMem(line string) (total float64, free float64, used float64, buffCache float64, err error) {
	matches := reMem.FindStringSubmatch(line)
	if len(matches) != 5 {
		return 0, 0, 0, 0, fmt.Errorf("get mem: len(matches) != 5")
	}
	total, err = strconv.ParseFloat(strings.TrimSpace(matches[1]), 64)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("get mem: %w", err)
	}
	free, err = strconv.ParseFloat(strings.TrimSpace(matches[2]), 64)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("get mem: %w", err)
	}
	used, err = strconv.ParseFloat(strings.TrimSpace(matches[3]), 64)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("get mem: %w", err)
	}
	buffCache, err = strconv.ParseFloat(strings.TrimSpace(matches[4]), 64)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("get mem: %w", err)
	}
	return total, free, used, buffCache, err
}

func getSwap(line string) (total float64, free float64, used float64, avail float64, err error) {
	matches := reSwap.FindStringSubmatch(line)
	if len(matches) != 5 {
		return 0, 0, 0, 0, fmt.Errorf("get swap: len(matches) != 5")
	}
	total, err = strconv.ParseFloat(strings.TrimSpace(matches[1]), 64)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("get swap: %w", err)
	}
	free, err = strconv.ParseFloat(strings.TrimSpace(matches[2]), 64)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("get swap: %w", err)
	}
	used, err = strconv.ParseFloat(strings.TrimSpace(matches[3]), 64)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("get swap: %w", err)
	}
	avail, err = strconv.ParseFloat(strings.TrimSpace(matches[4]), 64)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("get swap: %w", err)
	}
	return total, free, used, avail, err
}
