//go:build windows
// +build windows

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	griagent "github.com/javadmohebbi/griAgent/v1"
	svc "github.com/judwhite/go-svc"
)

// this type will implement svc.service
// and all Stop and Start functions are implemented here
// in main.go
type program struct {
	// GriAgent app
	svr *griagent.GriAgent

	// context which uses for timeout
	ctx context.Context
}

// config variable
var conf *griagent.GriConfig

// working dir variable
var dir string

// service program which is an instance of
// program struct
var prg *program

// service context
func (p *program) Context() context.Context {
	return p.ctx
}

func main() {
	// get current dir
	dir = getCurrentDir()

	// get conf
	conf = getConf()

	// prepare log file
	_ = prepareLogPathAndOutput()

	// prepare SIGINT channel variable
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	// make chan var
	err := make(chan error)
	done := make(chan bool)
	tme := make(chan bool)

	// create new instance of griAgent
	ga := griagent.New(
		conf,
		dir,
		sigs,
		done,
		err,
		tme,
		uint32(os.Getpid()),
	)

	// prepare windwos service context with timeout
	// timeout defines in json config file
	// and the unit is minutes
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(conf.Timeout*uint(time.Minute)),
	)

	// run whole process in
	// a gorutine
	go func() {

		// cancel context by force, assuming the whole process is complete
		defer cancel()

		// prepare windwos service with our program
		prg = &program{
			svr: ga,
			ctx: ctx,
		}

		// svc.Run will call Init, Start, and Stop
		if err := svc.Run(prg); err != nil {
			log.Fatal(err)
		}

	}()

	select {
	case <-ctx.Done():
		switch ctx.Err() {
		case context.DeadlineExceeded:

			// send signal to service
			log.Println("timeout exceeded for this task")
			tme <- true

		case context.Canceled:

			// send signal to service
			log.Println("task context canceled by force")
			tme <- true
		}
	}

}

// its needed by go-svc
func (p *program) Init(env svc.Environment) error {
	return nil
}

// this will start windows service
func (p *program) Start() error {
	log.Printf("Starting griAgent service...\n")
	go p.svr.Start()
	return nil
}

// this will stop windows service
func (p *program) Stop() error {

	log.Printf("Stopping griAgent service...\n")
	if err := p.svr.Stop(); err != nil {
		return err
	}
	log.Printf("The griAgent service Stopped.\n")
	return nil
}

// prepare loggin system
func prepareLogPathAndOutput() *os.File {
	logPath := filepath.Join(dir, fmt.Sprintf("griAgent_%v_%v.log",
		conf.TaskID, conf.Time,
	))

	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	// set logfile to log on terminal and file
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)

	// return file for further usage
	return f

}

// get conig
func getConf() *griagent.GriConfig {
	conf, err := griagent.NewConfig(dir)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}

// get executable dir
func getCurrentDir() string {
	d, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
	}
	return d
}
