package main

import "fmt"

func main() {
	file, fileErr := LoadFile("../.env")

	if fileErr != nil {
		fmt.Println(fileErr.Error())
	}

	input, inputErr := ParseJson(file)

	if inputErr != nil {
		fmt.Println(inputErr.Error())
	}

	fmt.Println(input.Data)

	data, err := StingifyData(*input)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(data))
}
