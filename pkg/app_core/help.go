package app_core

var help_info string

func init() {
	help_info = `

A demo go app

COMMAND:
	help		View the help documentation
	exit		Exit the program
	mode		Show the app mode information

To get more help with this demo, visit https://github.com/Amovement/go-app-demo

	`
}

func GetAppHelpInfo() string {
	return help_info
}
