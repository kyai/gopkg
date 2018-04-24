package proc

import (
	"fmt"
	"gopkg/file"
	"os"
	"os/signal"
	"syscall"
)

var (
	FileName = "run.pid"
)

func init() {
	start()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		stop()
		os.Exit(0)
	}()
}

func start() {
	file.Writer("./"+FileName, fmt.Sprintf("%d", os.Getpid()))
}

func stop() {
	file.Remove("./" + FileName)
}
