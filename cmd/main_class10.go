package main

import (
	"fmt"
)

func main_class10() {

	/*Explain what is a rune and relation with a string */

	a := [3]int{1, 2, 3}

	// a := [3]{1}, a :=[...]{1}
	b := a[0:1]
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	c := b[0:2]
	fmt.Println("c=", c)

	fmt.Println(len(b))
	fmt.Println(cap(b))

	fmt.Println(len(c))
	fmt.Println(cap(c))

	d := a[0:1:1]
	fmt.Println("d=", d)

	fmt.Println(len(d))
	fmt.Println(cap(d))

	e := d[0:2]
	fmt.Println("e=", e)

}
