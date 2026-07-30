package fns

import (
	"encoding/json"
	"os"
)

func LoadFile(fileName string) ([]byte, error) {
	if file, err := os.ReadFile(fileName); err != nil {
		return []byte(""), err
	} else {
		return file, nil
	}
}

type Input struct {
	Data string `json:"data"`
}

func ParseJson(data []byte) (*Input, error) {
	var input Input

	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	} else {
		return &input, nil
	}
}

func StingifyData(input Input) ([]byte, error) {
	if data, err := json.Marshal(input); err != nil {
		return []byte(""), err
	} else {
		return data, nil
	}
}
