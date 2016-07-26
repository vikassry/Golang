package main

import (
	"fmt"
)

func second () {
	r := recover()
	fmt.Printf("I am second! %s\n", r)
}

func first() {
// 	defer func() {
// 	fmt.Printf("I am second!\n")
// }()
	fmt.Printf("I am First!\n")
}

func main() {
	first()
	first()
	fmt.Printf("I am at the end!\n")
	defer second()

	panic("panicccccccccccccccccc")
}
