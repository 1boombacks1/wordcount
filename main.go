package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Print(0)
			return
		}
	}()
	count_words := len(strings.Split(os.Args[1], " "))
	fmt.Print(count_words)

}
