//go:build windows
// +build windows

package griagent

import (
	"fmt"
	"log"
	"os/exec"
)

// will kill all remaining children
func (s *GriAgent) killThemAll(children []Monitoredchild) {
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
