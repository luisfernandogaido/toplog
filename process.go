package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Process struct {
	Time        time.Time `json:"time"`
	Pid         int       `json:"pid"`
	User        string    `json:"user"`
	Pr          string    `json:"pr"`
	Ni          int       `json:"ni"`
	Virt        float64   `json:"virt"`
	Res         float64   `json:"res"`
	Shr         float64   `json:"shr"`
	S           string    `json:"s"`
	PercCpu     float64   `json:"perc_cpu"`
	PercMem     float64   `json:"perc_mem"`
	TimeSeconds float64   `json:"time_seconds"`
	Command     string    `json:"command"`
}

func getProcess(t time.Time, line string) (Process, error) {
	allWords := strings.Split(line, " ")
	words := make([]string, 0, len(allWords))
	for _, w := range allWords {
		if w == "" {
			continue
		}
		words = append(words, w)
	}
	if len(words) != 12 {
		fmt.Println(line, len(words))
		return Process{}, fmt.Errorf("getprocess: len(words) != 12")
	}
	p := Process{Time: t}
	pid, err := strconv.Atoi(words[0])
	if err != nil {
		return Process{}, fmt.Errorf("get process: %w", err)
	}
	p.Pid = pid
	p.User = words[1]
	p.Pr = words[2]
	ni, err := strconv.Atoi(words[3])
	if err != nil {
		return Process{}, fmt.Errorf("get process: %w", err)
	}
	p.Ni = ni
	virt, err := myParseFloat(words[4])
	if err != nil {
		return Process{}, fmt.Errorf("get process: %w", err)
	}
	p.Virt = virt
	res, err := myParseFloat(words[5])
	if err != nil {
		return Process{}, fmt.Errorf("get process: %w", err)
	}
	p.Res = res
	shr, err := myParseFloat(words[6])
	if err != nil {
		return Process{}, fmt.Errorf("get process: %w", err)
	}
	p.Shr = shr
	p.S = words[7]
	percCpu, err := myParseFloat(words[8])
	if err != nil {
		return Process{}, fmt.Errorf("get process: %w", err)
	}
	p.PercCpu = percCpu
	percMem, err := myParseFloat(words[9])
	if err != nil {
		return Process{}, fmt.Errorf("get process: %w", err)
	}
	p.PercMem = percMem
	timeSeconds, err := getTimePlus(words[10])
	if err != nil {
		return Process{}, fmt.Errorf("get process: %w", err)
	}
	p.TimeSeconds = timeSeconds
	p.Command = words[11]
	return p, nil
}

func myParseFloat(s string) (float64, error) {
	g := false
	if strings.Contains(s, "g") {
		g = true
		s = strings.Replace(s, "g", "", 1)
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("myparsefloat: %w", err)
	}
	if g {
		f *= 1024 * 1024
	}
	return f, nil
}

func getTimePlus(s string) (float64, error) {
	partes := strings.Split(s, ":")
	minutos, err := strconv.ParseFloat(partes[0], 64)
	if err != nil {
		return 0, fmt.Errorf("gettimeplus: %w", err)
	}
	segundos, err := strconv.ParseFloat(partes[1], 64)
	if err != nil {
		return 0, fmt.Errorf("gettimeplus: %w", err)
	}
	return minutos*60 + segundos, nil
}
