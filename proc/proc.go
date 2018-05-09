package proc

import (
	"fmt"
	"io/ioutil"
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
	ioutil.WriteFile("./"+FileName, []byte(fmt.Sprintf("%d", os.Getpid())), 0666)
}

func stop() {
	os.Remove("./" + FileName)
}
