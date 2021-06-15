package main

import (
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	tops := make([]Top, 0)
	procs := make([]Process, 0)
AquiO:
	for {
		cmd := exec.Command("top", "-i", "-n 1", "-b", "-w", "512")
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err, string(stdoutStderr))
		}
		now := time.Now()
		out := string(stdoutStderr)
		lines := strings.Split(out, "\n")
		top, err := getTop(now, lines)
		if err != nil {
			log.Println(err)
			continue
		}
		tops = append(tops, top)
		lines = lines[7:]
		for _, line := range lines {
			if line == "" {
				continue
			}
			proc, err := getProcess(now, line)
			if err != nil {
				log.Println(err)
				continue AquiO
			}
			procs = append(procs, proc)
		}
		if len(tops) == 60 {
			if err := writeTops(tops); err != nil {
				log.Println(err)
				continue AquiO
			}
			if err := writeProcs(procs); err != nil {
				log.Println(err)
				continue AquiO
			}
			tops = nil
			procs = nil
		}
		time.Sleep(time.Second * 5)
	}
}
