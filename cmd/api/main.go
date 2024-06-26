package main

import "applicationDesignTest/internal/api"

func main() {
	cfg := api.NewConfig()
	app := NewApp(cfg)
	app.Run()
}
