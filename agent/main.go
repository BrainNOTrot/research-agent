package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	agent := NewAgent(
		"artificial intelligence breakthroughs",
	)

	agent.Run()

	fmt.Println("Agent running...")

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-stop

	agent.Stop()

	fmt.Println("Shutdown complete")
}
