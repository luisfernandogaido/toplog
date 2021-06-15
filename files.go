package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func writeTops(tops []Top) error {
	fileName := time.Now().Format("tops-20060102.jsonl")
	f, err := os.OpenFile("../"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("writetops: %w", err)
	}
	defer f.Close()
	for _, top := range tops {
		b, err := json.Marshal(top)
		if err != nil {
			return fmt.Errorf("writetops: %w", err)
		}
		f.Write(b)
		f.WriteString("\n")
	}
	return nil
}

func writeProcs(tops []Process) error {
	fileName := time.Now().Format("procs-20060102.jsonl")
	f, err := os.OpenFile("../"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("writeprocs: %w", err)
	}
	defer f.Close()
	for _, top := range tops {
		b, err := json.Marshal(top)
		if err != nil {
			return fmt.Errorf("writeprocs: %w", err)
		}
		f.Write(b)
		f.WriteString("\n")
	}
	return nil
}
