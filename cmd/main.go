package main

import (
	"fmt"
	"go-web-challenge/internal/application"
	"os"
)

func main() {
	// env
	os.Setenv("SERVER_ADDR", "localhost:8080")
	os.Setenv("DB_FILE", "./docs/db/tickets.csv")
	// ...

	// application
	// - config
	cfg := &application.ConfigAppDefault{
		ServerAddr: os.Getenv("SERVER_ADDR"),
		DbFile:     os.Getenv("DB_FILE"),
	}
	app := application.NewApplicationDefault(cfg)

	// - setup
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}

	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
