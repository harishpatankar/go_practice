package main

import (
	"fmt"

	"github.com/harishpatankar/go_practice/tree/master/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
