//go:build windows
// +build windows

package griagent

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// this will be called service is started
func (s *GriAgent) Start() {

	// // monitored children
	// s.wev.SetMonitoredChildren(s.monitoredChild)

	// prepare chan variables
	s.exit = make(chan struct{})
	s.err = make(chan error)
	s.done = make(chan bool)

	// prepare server address that this app
	// should connect to
	deployer := fmt.Sprintf("%s:%d",
		s.Config.DeployerAddress,
		s.Config.DeployerPort,
	)

	log.Printf("dialing tcp on %s\n", deployer)

	// try to connect to deployer server
	// using TCP connection to send it's event
	// to that server
	con, err := net.Dial("tcp", deployer)
	if err != nil {
		log.Println("could not connect to tcp: ", deployer)
		os.Exit(int(ERR_UNIX_CLIENT_SOCKET))
	}

	// initialize this host for the server
	rq := ClientServerReqResp{
		Agent:     true,
		Command:   CMD_INIT,
		HostID:    s.Config.HostID,
		Host:      s.Config.Host,
		RequestID: fmt.Sprintf("req-%v", time.Now().Unix()),
		TaskID:    s.Config.TaskID,
	}

	log.Printf("initializing tcp connection with init command on %s\n", deployer)

	// convert struct to json string
	bts, err := rq.JSONToStringClientServerReqResp()
	if err != nil {
		log.Println("could not marshal to json: ", err)
		os.Exit(int(ERR_UNIX_CLIENT_SOCKET_MARSHAL))
	}

	// initialize socket client
	_, err = con.Write([]byte(fmt.Sprintf("%s\n", bts)))
	if err != nil {
		if err != nil {
			fmt.Println("could not initialize tcp client: ", err)
			os.Exit(int(ERR_UNIX_CLIENT_SOCKET_INIT))
		}
	}

	go func() {
		// do the monitoring job
		err, _ := s.wev.Do()
		if err != nil {
			s.err <- err
			// s.killThemAll(children)
		}
		// nothing more to do
		s.done <- true
	}()

	// sleep for some seconds
	// to wait wmi event listener
	time.Sleep(5 * time.Second)

	// handle tcp requests
	s.Conn = con
	go s.handleTcpReqests()

	s.wg.Add(1)
	// s.task <- task_type(1)
	go s.DoTheJob()

	//
	for {
		select {
		case <-s.sigs:
			s.killThemAll(s.wev.monitoredchild)
			s.ExecError(
				errors.New("SIGINT recieved from the OS or the user"),
			)
		case err := <-s.err:
			s.killThemAll(s.wev.monitoredchild)
			s.ExecError(err)
		case d := <-s.done:
			if d {
				s.ExecSuccess()
			}

		case tm := <-s.timeoutCh:
			if tm {
				children := "this app remain open on the target machines: "
				for i, mc := range s.wev.monitoredchild {
					children += fmt.Sprintf("%s (%d)", mc.Name, mc.ProcessID)
					if i < len(s.wev.monitoredchild)-1 {
						children += ", "
					}
				}
				if s.wev.lastExitStatus == 0 {
					// last child exit status is ok
					// so we could send a success custome message
					s._rq(CMD_BOOTSTRAP_FINISH_DONE, children)
				} else {
					// last child exit status is not ok
					// so we should send a error message
					s.killThemAll(s.wev.monitoredchild)
					s.ExecError(err)
				}
			}
		}
	}

}
