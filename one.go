package gobagoeslib

import (
	"fmt"
	"strings"
)

func PrintMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
