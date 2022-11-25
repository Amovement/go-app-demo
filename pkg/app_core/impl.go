package app_core

import (
	"fmt"
	"os"

	config "github.com/Amovement/go-app-demo/pkg/config"
)

func getMode(_ Core, _ ...interface{}) {
	cfg := config.GetConfig()
	fmt.Printf("App Mode: %s\n", cfg.App.Mode)
}

func getAppHelpInfo(_ Core, _ ...interface{}) {
	fmt.Printf("%s\n", GetAppHelpInfo())
}

func exitApp(app Core, _ ...interface{}) {
	app.Exit <- os.Interrupt
}
