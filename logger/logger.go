package logger

import "fmt"

func Log(message string) {
	fmt.Println("[LOG] " + message)
}
