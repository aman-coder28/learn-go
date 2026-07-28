package main

import (
	"fmt"
)

func main() {
	var BasicCalc = NewCalculator()

	var result, err = BasicCalc.Calculate()

	if err != nil {
		fmt.Printf("\n%s\n", err.Error())

		Restart()
	} else {
		fmt.Printf("%s\n\n", result)
		Restart()
	}
}
