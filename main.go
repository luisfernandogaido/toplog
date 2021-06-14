package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

type TopPro struct {
}

func main() {
	tops := make([]Top, 0)
	for {
		cmd := exec.Command("top", "-i", "-n 1", "-b", "-o", "%MEM")
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err, stdoutStderr)
		}
		out := string(stdoutStderr)
		lines := strings.Split(out, "\n")
		fmt.Println(lines)
		top := Top{
			Time: time.Now(),
		}
		tops = append(tops, top)
		if len(tops) == 2 {
			tops = nil
			fmt.Println(tops)
		}
		time.Sleep(time.Second * 10)
	}
}
