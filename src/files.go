package main

import (
	"errors"
	"os"
)

func LoadFile(fileName string) (string, error) {
	if file, err := os.ReadFile(fileName); err != nil {
		return "", errors.New(err.Error())
	} else {
		return string(file), nil
	}
}
