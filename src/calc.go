package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Calculator struct {
	first  float64
	opr    rune
	second float64
}

func (c Calculator) String() string {
	return fmt.Sprintf("\n%.2f %c %.2f = ", c.first, c.opr, c.second)
}

func (c Calculator) Calculate() (string, error) {
	switch c.opr {
	case '+':
		return fmt.Sprintf("%s%.2f", c.String(), c.first+c.second), nil
	case '-':
		return fmt.Sprintf("%s%.2f", c.String(), c.first-c.second), nil
	case 'x':
		return fmt.Sprintf("%s%.2f", c.String(), c.first*c.second), nil
	case '/':
		return fmt.Sprintf("%s%.2f", c.String(), divide(c.first, c.second)), nil
	default:
		return "0", errors.New("Error: Invalid Operator")
	}
}

func NewCalcualtor() Calculator {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a First Number: ")
	var f = get_float(read_line(scanner))

	var opr rune
	fmt.Println("Enter an Operator (+, -, /, x): ")

	if o := read_line(scanner); len(o) != 0 {
		opr = rune(strings.TrimSpace(o)[0])
	} else {
		opr = rune("*"[0])
	}

	fmt.Println("Enter a Second Number: ")
	var s = get_float(read_line(scanner))

	return Calculator{f, opr, s}
}

func divide(f float64, s float64) float64 {
	if s == 0 {
		return float64(0)
	}

	return f / s
}

func read_line(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}

func get_float(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)

	if err != nil {
		return float64(0)
	}

	return f
}

func Restart() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("To Restart, Press r.")
	fmt.Println("To Quit, Press q or any character.")

	res := read_line(scanner)

	if len(strings.TrimSpace(res)) != 0 {
		switch res[0] {
		case 'r':
			ClearScreen()

			var result, error = NewCalcualtor().Calculate()

			if error != nil {
				fmt.Printf("\n%s\n", error.Error())

				Restart()
			}

			fmt.Printf("%s\n\n", result)

			Restart()

		default:
			os.Exit(1)
		}
	}
}

func ClearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
