package main

import (
	"errors"
	"fmt"
	"hello/hello"
	"os"
)

func main_class2() {

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1) // ← exit honnête, APRÈS les defer de run()
	}
}

func run() error {

	if len(os.Args) > 1 {
		fmt.Println(hello.Say(os.Args[1:]))
		return nil
	}
	return errors.New("usage: hello <name>...")

}
