package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	griagent "github.com/javadmohebbi/griAgent/v1"
)

func main() {

	// get pid from command line arg
	pid := flag.Uint("p", 0, "PID to monitor")
	flag.Parse()

	// simple check
	if *pid == 0 {
		log.Println("invalid pid: ")
		flag.PrintDefaults()
	}

	// prepare SIGINT channel variable
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	wev := griagent.NewWMIWmiMonitorProcessEvents(
		// parent taht we want to monitor
		// its child
		uint32(*pid),

		// within time in seconds
		"0.1",

		// sig int channel variable
		sigs,
	)

	// do the monitoring job
	err, children := wev.Do()

	// sleep for some seconds
	// to wait wmi event listener
	time.Sleep(5 * time.Second)

	if err != nil {

		log.Println(err)

		// we could decide to kill a process
		// or leave them. But usually we should kill
		// them because their parents are killed
		killThemAll(children)
	}

	log.Println("Exiting....!")

}

// will kill all remaining children
func killThemAll(children []griagent.Monitoredchild) {
	for _, child := range children {
		pid := fmt.Sprintf("%d", child.ProcessID)
		cmd := exec.Command("taskkill", "/F", "/PID", pid)
		if err := cmd.Start(); err != nil {
			log.Printf("can not kill %s(%d) due to error: %v\n",
				child.Name, child.ProcessID, err,
			)
		}
		log.Printf("%s(%d) killed",
			child.Name, child.ProcessID,
		)
	}
}
