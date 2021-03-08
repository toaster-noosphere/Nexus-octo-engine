package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/TheTh1rt33nth/Nexus-octo-engine/pkg/server"
)

func main() {
	s := server.New("127.0.0.1:1337")
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go s.Run()

	<-done
	signal.Stop(done)
}
