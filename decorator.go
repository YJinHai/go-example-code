package main

import (
	"fmt"
	"log"
)

func logging(f func() int) func()  {
	return func() {
		log.Println("Println user")
		f()
	}
}

func User() int {
	fmt.Println("i am yangjinhai")
}

func main() {
	User := logging(User)
	User()
}
