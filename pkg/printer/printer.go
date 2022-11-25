package printer

import "fmt"

func PrintMessage(source string, level string, message string) {
	fmt.Printf("\n[%s] [%s] -- %s \n", source, level, message)
	fmt.Print(">")
}
