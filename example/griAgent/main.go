//go:build windows
// +build windows

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	griagent "github.com/javadmohebbi/griAgent/v1"
)

var conf *griagent.GriConfig
var dir string

func main() {
	// dir of current exec path
	// dir = getCurrentDir()
	dir = "C:\\Windows\\Temp\\YARMA"

	conf = getConf()

	_ = prepareLogPathAndOutput()

	// prepare SIGINT channel variable
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	// make chan var
	err := make(chan error)
	done := make(chan bool)
	tm := make(chan bool)

	// create new instance of griAgent
	ga := griagent.New(
		conf,
		dir,
		sigs,
		done,
		err,
		tm,
		uint32(os.Getpid()),
	)
	log.Println("starting the service from main func")
	ga.Start()

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
