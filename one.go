package gobagoeslib

import (
	"fmt"
	"strings"
)

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}