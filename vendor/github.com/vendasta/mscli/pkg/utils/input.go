package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetUserInput(message string, defaultValue string) string {
	fmt.Printf("%s: (%s) ", message, defaultValue)
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	value := reader.Text()
	if value == "" {
		value = defaultValue
	}
	return value
}
