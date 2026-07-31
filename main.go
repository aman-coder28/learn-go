package main

import (
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

	password := fns.ReadPassword()

	fns.ClearScreen()

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
