package main

import (
	"fmt"
)

func second () {
	fmt.Printf("I am second!\n")
}

func first() {
	defer func() {
	fmt.Printf("I am second!\n")
}()
	fmt.Printf("I am First!\n")
}

func main() {
	defer second()
	first()
	first()
	fmt.Printf("I am at the end!\n")

}
