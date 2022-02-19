//go:build windows
// +build windows

package griagent

import (
	"fmt"
	"log"
	"os"
)

// this fucntion will be called in case
// our service faces an error
func (s *GriAgent) ExecError(err error) {
	msg := fmt.Sprintf("%v, err: %v", s.bootstrap, err)

	log.Printf("[3] bootstrap (%v) failed: %v\n", s.bootstrap.Process.Pid, msg)

	s._rq(CMD_BOOTSTRAP_FINISH_ERROR, msg)

	s.wev.Stop()

	os.Exit(int(CMD_BOOTSTRAP_FINISH_ERROR))
}
