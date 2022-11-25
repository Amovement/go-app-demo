package app_core

import (
	"bufio"
	"fmt"
	"github.com/Amovement/go-app-demo/pkg/printer"
	"os"
	"strings"
)

type Command map[string]func(Core, ...interface{})

type Core struct {
	Cmd  Command
	Exit chan os.Signal
}

func NewApp() (Core, error) {
	app := Core{
		Cmd: Command{
			"help": getAppHelpInfo,
			"exit": exitApp,
			"mode": getMode,
		},
		Exit: make(chan os.Signal),
	}

	return app, nil
}

func (app Core) Run() {
	app.Cmd["help"](app, "")
	fmt.Print(">")

	buf := bufio.NewScanner(os.Stdin)
	for buf.Scan() {
		input := strings.Split(buf.Text(), " ")
		inputCommand := input[0]
		inputOption := input[1:]
		if _, ok := app.Cmd[inputCommand]; !ok {
			app.Cmd["help"](app, "")
			printer.PrintMessage("APP", "Error", "Unknown Command: "+inputCommand)
		} else {
			app.Cmd[inputCommand](app, inputOption)
			printer.PrintMessage("APP", "Info", "Executing the Command: "+inputCommand)
		}
	}
}
