package main

import (
	"bufio"
	"fmt"
	fns "learning/functions"
	"os"
)

func main() {
	fileName := fns.GetArgs()

	content, err := fns.LoadFile(fileName)

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Fprintf(os.Stdin, "Enter Password: ")
	password := fns.ReadLine(scanner)

	if json, ok := fns.IsJson(string(content)); ok == true {
		decrypted := fns.Decrypt(password, json.Data)

		os.Stdout.WriteString(decrypted)
	} else {
		encrypted := fns.Encrypt(password, string(content))

		input := fns.Input{
			Data: encrypted,
		}

		data, err := fns.StringifyData(input)

		if err != nil {
			fmt.Println(err.Error())

			return
		}

		fns.WriteFile(string(data))
	}
}
