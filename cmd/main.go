package main

import (
	"fmt"
	"log"

	griagent "github.com/javadmohebbi/griAgent/v1"
)

func main() {

	wp := griagent.WinAPI{}

	procType, err := wp.GetProcessorInfo()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("cpu arch: ", procType)

	// major, minor, build := wp.GetWinVersion()

	// fmt.Print("windows version ", major, ".", minor, " (Build ", build, ")\n")

	pids, err := wp.EnumProcesses()
	if err != nil {
		log.Fatalln(err)
	}
	for i, pid := range pids {
		fmt.Printf("%d) %d \t", i, pid)
	}
	fmt.Println()
}
