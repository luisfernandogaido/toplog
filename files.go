package main

import (
	"fmt"
	"os"
	"time"
)

func writeTops(tops []Top) error {
	fileName := time.Now().Format("tops-20060102.jsonl")
	f, err := os.OpenFile("./"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("writetops: %w", err)
	}
	defer f.Close()
	return nil
}
