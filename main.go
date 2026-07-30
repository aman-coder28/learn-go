package main

import (
	"fmt"
	fns "learning/functions"
)

func main() {
	file, fileErr := fns.LoadFile(".env")

	if fileErr != nil {
		fmt.Println(fileErr.Error())

		return
	}

	input, inputErr := fns.ParseJson(file)

	if inputErr != nil {
		fmt.Println(inputErr.Error())

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
