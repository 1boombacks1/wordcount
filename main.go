package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if os.Args[1] == "" {
		fmt.Print(0)
	} else {
		count_words := len(strings.Split(os.Args[1], " "))
		fmt.Print(count_words)
	}

}
