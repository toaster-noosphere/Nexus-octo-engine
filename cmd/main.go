package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/TheTh1rt33nth/Nexus-octo-engine/internal/api"
	"github.com/TheTh1rt33nth/Nexus-octo-engine/internal/app"
)

func main() {
	s := api.NewServer("127.0.0.1:1337", &app.Engine{})
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go s.Run()

	<-done
	signal.Stop(done)
}
