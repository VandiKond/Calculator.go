package application

import (
	"log"
	"os"
	"time"
)

type Application struct {
	Duration time.Duration
}

func New(d time.Duration) *Application {
	return &Application{
		Duration: d,
	}
}

func (a *Application) Run() error {
	// Exiting in duration
	defer log.Printf("application stopped before timeout")
	go a.ExitTimeOut()

	// The program

	// The program end

	// Returning without error
	return nil
}

func (a *Application) ExitTimeOut() {
	time.Sleep(a.Duration)
	log.Printf("timeout %s has passed. Ending the program", a.Duration)
	os.Exit(418)
}
