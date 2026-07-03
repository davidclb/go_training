package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main_class4() {

	/*Explain what is a rune and relation with a string */

	v := "公共汽车"
	fmt.Printf("%8T %[1]v\n", v)

	fmt.Printf("%8T %[1]v\n", []rune(v))
	fmt.Printf("%8T %[1]v\n", []byte(v))

	if len(os.Args) < 3 {

		fmt.Fprintln(os.Stderr, " not enough args")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)

		fmt.Println(t)
	}
}
