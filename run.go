package main

import (
	"mydocker/container"
	"os"

	log "github.com/sirupsen/logrus"
)

/*Run starts a container*/
func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
