package main

import (
	"fmt"
	"github.com/chzyer/readline"
)

func main() {
	fmt.Println("Lispy Version 0.0.1")
	fmt.Println("Press Ctrl+c to Exit\n")

		rl, err := readline.New("glispy> ")
		if err != nil {
			panic(err)
		}
		defer rl.Close()
	for{
		line, err := rl.Readline()
		if err != nil {
			break
		}
		fmt.Println(line)
	}

}
