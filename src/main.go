package main

import (
	"fmt"
)

func main() {
	var BasicCalc = NewCalcualtor()

	var result, error = BasicCalc.Calculate()

	if error != nil {
		fmt.Printf("\n%s\n", error.Error())

		Restart()
	}

	fmt.Printf("%s\n\n", result)
	Restart()
}
