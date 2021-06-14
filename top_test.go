package main

import (
	"fmt"
	"testing"
)

func TestGetUp(t *testing.T) {
	fmt.Println(getUp("up 27 min,"))
	fmt.Println(getUp("top - 10:29:19 up 1 day, 21:49,  2 users,  load average: 0.00, 0.00, 0.00"))
	fmt.Println(getUp("top - 10:28:56 up 210 days, 23:41,  1 user,  load average: 0.09, 0.07, 0.08"))
	fmt.Println(getUp("top - 10:49:48 up  1:03,  1 user,  load average: 0.26, 0.59, 0.76"))
	fmt.Println(getUp("top - 13:14:26 up 2 days, 34 min,  1 user,  load average: 0.03, 0.08, 0.08"))
}

func TestGetLoadAverage(t *testing.T) {
	fmt.Println(getLoadAverage("top - 10:29:19 up 1 day, 21:49,  2 users,  load average: 0.00, 0.00, 0.00"))
	fmt.Println(getLoadAverage("top - 10:28:56 up 210 days, 23:41,  1 user,  load average: 0.09, 0.07, 0.08"))
	fmt.Println(getLoadAverage("top - 10:49:48 up  1:03,  1 user,  load average: 0.26, 0.59, 0.76"))
	fmt.Println(getLoadAverage("top - 13:14:26 up 2 days, 34 min,  1 user,  load average: 0.03, 0.08, 0.08"))
}

func TestGetTasks(t *testing.T) {
	fmt.Println(getTasks("Tasks: 125 total,   1 running, 124 sleeping,   0 stopped,   0 zombie"))
}

func TestGetCpuPercs(t *testing.T) {
	fmt.Println(getCpuPercs("%Cpu(s):  5.9 us,  2.9 sy,  0.0 ni, 88.2 id,  0.0 wa,  0.0 hi,  2.9 si,  0.0 st"))
}

func TestGetMem(t *testing.T) {
	fmt.Println(getSwap("MiB Swap:      0.0 total,      0.0 free,      0.0 used.   1640.5 avail Mem"))
}
