package main

import (
	"log"
	"os"
)

const pid string = "/proc/1/cmdline"

func main() {
	_, err := os.ReadFile(pid)
	if err != nil {
		log.Fatalf("failed to read %s: %v", pid, err)
	}
}