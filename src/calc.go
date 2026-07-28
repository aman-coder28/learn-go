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
	return fmt.Sprintf("%.2f %c %.2f = ", c.first, c.opr, c.second)
}

func (c Calculator) Calculate() (string, error) {
	switch c.opr {
	case '+':
		return fmt.Sprintf("\n%s%.2f", c.String(), c.first+c.second), nil
	case '-':
		return fmt.Sprintf("\n%s%.2f", c.String(), c.first-c.second), nil
	case 'x':
		return fmt.Sprintf("\n%s%.2f", c.String(), c.first*c.second), nil
	case '/':
		return fmt.Sprintf("\n%s%.2f", c.String(), c.first/c.second), nil
	default:
		return "0", errors.New("Error: Invalid Operator")
	}
}

func NewCalcualtor() Calculator {
	fmt.Println("Enter a First Number: ")
	var f = get_float(read_line())

	var opr rune
	fmt.Println("Enter an Operator (+, -, /, x): ")

	if o := read_line(); len(o) != 0 {
		opr = rune(strings.TrimSpace(o)[0])
	} else {
		opr = rune("*"[0])
	}

	fmt.Println("Enter a Second Number: ")
	var s = get_float(read_line())

	return Calculator{f, opr, s}
}

func read_line() string {
	scanner := bufio.NewScanner(os.Stdin)

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
	fmt.Println("To Restart, Press r.")
	fmt.Println("To Quit, Press q or any character.")

	res := read_line()

	if len(strings.TrimSpace(res)) != 0 {
		switch res[0] {
		case 'r':
			ClearScreen()

			NewCalcualtor().Calculate()
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
