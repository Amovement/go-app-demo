package app_core

import (
	"fmt"

	config "github.com/Amovement/go-app-demo/pkg/config"
)

func getMode(options ...interface{}) {
	cfg := config.GetConfig()
	fmt.Printf("App Mode: %s\n", cfg.App.Mode)
}

func getAppHelpInfo(options ...interface{}) {
	fmt.Printf("%s\n", GetAppHelpInfo())
}
