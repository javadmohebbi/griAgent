//go:build windows
// +build windows

package griagent

import (
	"net"
	"os"
	"os/exec"
	"sync"
)

// struct that use for the service to
// do the installation task
type GriAgent struct {
	monitoredChild *[]Monitoredchild

	// PID of griAgent
	pid uint32

	// this will check if task timeout
	// is exceeded
	timeoutCh chan bool

	// os signals like CTRL + D
	sigs chan os.Signal

	// will be used if wmievent proce mon done the job
	procDone chan bool

	// will be used if wmievent proce mon stopped with error
	procErr chan error

	// CPU architecture type
	arch procArchType

	// error channel, it will be filled if any error
	// is happening on the task
	err chan error

	// done will be true if the task completed successfully
	done chan bool

	// it uses to handle connection between this app and the
	// server, if is not able to connect to server,
	// it will generate and error and exit
	Conn net.Conn

	// GriAgent configuration, normally will be read from a json file
	Config *GriConfig

	// Command to run
	// this is the very first executable that GriAgent will
	// execute. usually somthing like installer.exe, installer.msi , ...
	bootstrap *exec.Cmd

	// command line options
	// an array of command line parameters that Bootstrap
	// needs. Usually /silent, /quite or maybe /unattended , ...
	// check yout installation package manual to find the unattended commandline arguments
	bootstrapPath string

	//
	exit chan struct{}

	// it uses to manage concurrent goroutines
	wg sync.WaitGroup

	// service directory
	SvcDir string

	// an instance of WmiMonitorProcessEvents for monitoring
	// installation process and child processes
	wev *WmiMonitorProcessEvents
}
