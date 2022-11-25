package main

import (
	core "github.com/Amovement/go-app-demo/pkg/app_core"
)

func main() {
	app, _ := core.NewApp()
	go app.Run()
	<-app.Exit
}
