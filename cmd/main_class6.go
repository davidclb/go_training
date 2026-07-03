package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main_class7() {

	for _, fname := range os.Args[1:] {

		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		fmt.Println("The file has", len(data), "bytes")
		file.Close()
	}

}

func custom_wc() {

	//Return how many lines, words and how many chars

	var loc, wd, cc int

	for _, fname := range os.Args[1:] {

		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()
			wd += len(strings.Fields(s))
			cc += len(s)
			loc++
		}

		cc, err := io.ReadAll(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		fmt.Println("%7d %7d %7d %s \n", loc, wd, cc, fname)
		file.Close()
	}

}
