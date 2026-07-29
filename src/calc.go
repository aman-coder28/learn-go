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
<<<<<<< HEAD
		return fmt.Sprintf("%s%.2f", c.String(), Divide(c.first, c.second)), nil
=======
		if c.second == 0 {
			return fmt.Sprintf("%s%.2f", c.String(), float32(0)), nil
		}

		return fmt.Sprintf("%s%.2f", c.String(), c.first/c.second), nil
>>>>>>> f1ba937 (Pass Scanner from Main Func)
	default:
		return "0", errors.New("error: invalid operator")
	}
}

func NewCalculator(scanner *bufio.Scanner) Calculator {
	fmt.Println("Enter a First Number: ")
	var f = GetFloat(ReadLine(scanner))

	var opr rune
	fmt.Println("Enter an Operator (+, -, /, x): ")

	if o := ReadLine(scanner); len(o) != 0 {
		opr = rune(strings.TrimSpace(o)[0])
	} else {
		opr = rune("*"[0])
	}

	fmt.Println("Enter a Second Number: ")
	var s = GetFloat(ReadLine(scanner))

	return Calculator{f, opr, s}
}

<<<<<<< HEAD
func Divide(f float64, s float64) float64 {
	if s == 0 {
		return float64(0)
	}

	return f / s
}

=======
>>>>>>> f1ba937 (Pass Scanner from Main Func)
func ReadLine(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}

func GetFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)

	if err != nil {
		return float64(0)
	}

	return f
}

func RunCalculator(scanner *bufio.Scanner) {
	var calculator = NewCalculator(scanner)
	var result, err = calculator.Calculate()

	if err != nil {
		fmt.Printf("\n%s\n", err.Error())
	} else {
		fmt.Printf("%s\n\n", result)
	}

	fmt.Println("To Restart, Press r.")
	fmt.Println("To Quit, Press q or any character.")

	res := ReadLine(scanner)

	if len(strings.TrimSpace(res)) != 0 {
		switch res[0] {
		case 'r':
			ClearScreen()

			RunCalculator(scanner)
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
