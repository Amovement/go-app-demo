package app_core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command map[string]func(...interface{})

type Core struct {
	Cmd  Command
	Exit chan os.Signal
}

func NewApp() (Core, error) {
	app := Core{
		Cmd: Command{
			"help": getAppHelpInfo,
			"exit": func(options ...interface{}) {
				// TODO
			},
			"mode": getMode,
		},
	}

	return app, nil
}

func (app Core) Run() {
	app.Cmd["help"]("")
	fmt.Print("/")

	buf := bufio.NewScanner(os.Stdin)
	for buf.Scan() {
		input := strings.Split(buf.Text(), " ")
		inputCommand := input[0]
		inputOption := input[1:]
		if _, ok := app.Cmd[inputCommand]; !ok {
			fmt.Printf("[Error] Unknown Command: %s\n", inputCommand)
			app.Cmd["help"]("")
		} else {
			app.Cmd[inputCommand](inputOption)
		}
		fmt.Print("/")
	}
}
