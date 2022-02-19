//go:build windows
// +build windows

package griagent

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

// this fucntion will start all the
// execution steps for our service
func (s *GriAgent) ExecStarting() error {

	bootstrapPath := fmt.Sprintf("%s\\%s\\%s",
		os.Getenv("systemroot"),
		s.Config.Dir,
		s.Config.Bootstrap,
	)
	s.bootstrapPath = bootstrapPath

	s.bootstrap = exec.Command(bootstrapPath, s.Config.Params...)

	log.Printf("starting bootstrap: %s %s\n", bootstrapPath, s.Config.Params)

	err := s._rq(CMD_BOOTSTRAP_START, s.bootstrapPath)
	if err != nil {
		log.Println("Error:", err)
		s.ExecError(err)
		// os.Exit(1)
		return err
	}

	go func() {

		//starting cmd
		if err := s.bootstrap.Start(); err != nil {
			log.Printf("starting bootstrap failed: %v\n", err)
			s.err <- err
			return
		}

		log.Printf("bootstrap started: %s %s\n", bootstrapPath, s.Config.Params)

		// started
		err := s._rq(CMD_BOOTSTRAP_STARTED, s.bootstrapPath)
		if err != nil {
			s.err <- err
			return
		}

		log.Printf("wait bootstrap (%v) to finish: %s %s\n", s.bootstrap.Process.Pid, bootstrapPath, s.Config.Params)

		// check status
		if err := s.bootstrap.Wait(); err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				// This code copied from: https://stackoverflow.com/questions/10385551/get-exit-code-go/55055100
				// answer by https://stackoverflow.com/users/82219/tux21b
				//
				// The program has exited with an exit code != 0

				// This works on both Unix and Windows. Although package
				// syscall is generally platform dependent, WaitStatus is
				// defined for both Unix and Windows and in both cases has
				// an ExitStatus() method with the same signature.

				if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {

					log.Printf("[2]bootstrap (%v) failed: %s %s\n", s.bootstrap.Process.Pid, bootstrapPath, s.Config.Params)

					s.err <- errors.New(fmt.Sprintf("Exit Status: %d", status.ExitStatus()))
					return
				}
			} else {

				log.Printf("[1]bootstrap (%v) failed: %s %s\n", s.bootstrap.Process.Pid, bootstrapPath, s.Config.Params)

				s.err <- err
				return
			}
		}
		// cmd finished ok
		// s.done <- true

	}()

	return nil

}
