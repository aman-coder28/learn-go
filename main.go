package main

import (
	"fmt"
	fns "learning/functions"
)

func main() {
	file, fileErr := fns.LoadFile(".env")

	if fileErr != nil {
		fmt.Println(fileErr.Error())
	}

	input, inputErr := fns.ParseJson(file)

	if inputErr != nil {
		fmt.Println(inputErr.Error())
	}

	fmt.Println(input.Data)

	data, err := fns.StingifyData(*input)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(data))
}
