//go:build windows
// +build windows

package griagent

import (
	"log"
	"os"
)

// this fucntion will be called in case
// our service succeeded
func (s *GriAgent) ExecSuccess() {

	log.Printf("[4] bootstrap (%v) done\n", s.bootstrap.Process.Pid)

	s._rq(CMD_BOOTSTRAP_FINISH_DONE, s.bootstrapPath)

	s.wev.Stop()

	os.Exit(0)
}
