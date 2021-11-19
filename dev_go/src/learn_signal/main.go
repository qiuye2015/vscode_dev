package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("start")
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	sig := <-sigChan
	log.Println("exit sig", sig)
}
