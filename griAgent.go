//go:build windows
// +build windows

package griagent

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

// creates a new instance of GriAgent that should
// do the task
func New(conf *GriConfig, dir string, sigs chan os.Signal,
	procDone chan bool, procErr chan error,
	tmc chan bool,
	pid uint32,
) *GriAgent {

	// check proc version
	// this is for future use that
	// our task.yaml could have a specific bootstrap for
	// 64 or 32 bits CPU architectures
	// currently we don't use this check
	cpu, err := getProcessorInfo()
	if err != nil {
		log.Println(err)
		cpu = PROC_ARCH_UNSUPPORTED
	}

	log.Println("griAgent started with pid: ", pid)

	wev := NewWMIWmiMonitorProcessEvents(
		// parent taht we want to monitor
		// its child
		uint32(pid),

		// within time in seconds
		"0.1",

		// sig int channel variable
		sigs,
	)

	// return a *griAgent instance
	return &GriAgent{
		arch:     cpu,
		Config:   conf,
		SvcDir:   dir,
		sigs:     sigs,
		procDone: procDone,
		procErr:  procErr,
		pid:      pid,
		wev:      wev,
	}

}

// build the client request and return string
// to send to server
func (s *GriAgent) _rq(cmd Command, descPaylaod string) error {
	rq := ClientServerReqResp{
		Agent:       true,
		Command:     cmd,
		HostID:      s.Config.HostID,
		Host:        s.Config.Host,
		RequestID:   fmt.Sprintf("req-%v", time.Now().Unix()),
		DescPayload: descPaylaod,
		TaskID:      s.Config.TaskID,
	}

	bts, err := rq.JSONToStringClientServerReqResp()
	if err != nil {
		return errors.New(fmt.Sprintf("could not marshal to error: %v", err))
	}

	_, err = s.Conn.Write([]byte(fmt.Sprintf("%s\n", bts)))
	if err != nil {
		if err != nil {
			return errors.New(fmt.Sprintf("[%d] could not send tcp request: %v", ERR_TCP_CLIENT_AGENT_ERROR, err))
		}
	}

	return err

}
