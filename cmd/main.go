package main

import (
	"time"

	"github.com/VandiKond/Calculator.go.git/internal/application"
)

func main() {
	app := application.New(time.Second * 10)
	app.Run()
}
