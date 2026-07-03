package main

import "fmt"

func main_class9() {

	// Exemple de c'est le code qui bug pas le language : I don't understand the point underlined by this example
	s := make([]func(), 4)
	for i := 0; i <= 4; i++ {
		// solution i2:=i and print i2 &i2 as closure capture
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}

	}

	for i := 0; i <= 4; i++ {

		s[i]()
	}

}
