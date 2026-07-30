package main

import (
	"fmt"
	fns "learning/functions"
)

func main() {
	encryped := fns.Encrypt("password", "hello, how are you?")

	fmt.Println(encryped)
}
