package main

import (
	"fmt"
	fns "learning/functions"
)

func main() {
	file, err := fns.LoadFile(".env")

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	input, err := fns.ParseJson(file)

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Println(input.Data)

	data, err := fns.StringifyData(*input)

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Println(string(data))
}
