package fns

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/argon2"
)

func LoadFile(fileName string) ([]byte, error) {
	if file, err := os.ReadFile(fileName); err != nil {
		return []byte(""), err
	} else {
		return file, nil
	}
}

func GenerateKey(password string) []byte {
	salt := make([]byte, 16)

	rand.Read(salt)

	return argon2.Key([]byte("some password"), salt, 3, 32*1024, 4, 32)
}

func Encrypt(password string, text string) string {
	key := GenerateKey(password)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesgcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	return fmt.Sprintf("%x%x", nonce, ciphertext)
}

func Decrypt(password string, cipherText string) string {
	key := argon2.Key([]byte(password), []byte(cipherText[:17]), 3, 32*1024, 4, 32)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce, _ := hex.DecodeString(cipherText[17:29])
	text, _ := hex.DecodeString(cipherText[29:])

	plaintext, err := aesgcm.Open(nil, nonce, text, nil)

	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
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

func StringifyData(input Input) ([]byte, error) {
	if data, err := json.Marshal(input); err != nil {
		return []byte(""), err
	} else {
		return data, nil
	}
}
