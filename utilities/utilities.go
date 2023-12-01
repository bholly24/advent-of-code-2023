package utilities

import (
	"fmt"
	"os"
	"strings"
)

func ReadInLines(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}
