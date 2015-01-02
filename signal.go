package main

import (
	"github.com/henrysher/packer/packer"
	"github.com/henrysher/packer/packer/plugin"
	"log"
	"os"
	"os/signal"
)

// Prepares the signal handlers so that we handle interrupts properly.
// The signal handler exists in a goroutine.
func setupSignalHandlers(env packer.Environment) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	go func() {
		// First interrupt. We mostly ignore this because it allows the
		// plugins time to cleanup.
		<-ch
		log.Println("First interrupt. Ignoring to allow plugins to clean up.")

		env.Ui().Error("Interrupt signal received. Cleaning up...")

		// Second interrupt. Go down hard.
		<-ch
		log.Println("Second interrupt. Exiting now.")

		env.Ui().Error("Interrupt signal received twice. Forcefully exiting now.")

		// Force kill all the plugins, but mark that we're killing them
		// first so that we don't get panics everywhere.
		plugin.CleanupClients()
		os.Exit(1)
	}()
}
