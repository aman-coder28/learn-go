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
	"golang.org/x/term"
)

func ReadPassword() string {
	fmt.Fprint(os.Stderr, "Enter Password: ")

	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	return string(password)
}

func LoadFile(fileName string) ([]byte, error) {
	if file, err := os.ReadFile(fileName); err != nil {
		return []byte(""), err
	} else {
		return file, nil
	}
}

func IsJson(text string) (*Input, bool) {
	if json, err := ParseJson([]byte(text)); err != nil {
		return nil, false
	} else {
		return json, len(json.Data) != 0
	}
}

func GetArgs() string {
	fileName := ".env"

	if args := os.Args; len(args) > 1 {
		return args[2]
	} else {
		return fileName
	}
}

func WriteFile(data string) {
	os.WriteFile(".env", []byte(data), 0777)
}

func GenerateKey(password string, salt []byte) []byte {
	return argon2.Key([]byte(password), salt, 3, 32*1024, 4, 32)
}

func Encrypt(password string, text string) string {
	salt := make([]byte, 16)

	if _, err := rand.Read(salt); err != nil {
		panic(err.Error())
	}

	key := GenerateKey(password, salt)
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

	return fmt.Sprintf("%x%x%x", salt, nonce, ciphertext)
}

func Decrypt(password string, cipherText string) string {
	salt, _ := hex.DecodeString(cipherText[:32])
	nonce, _ := hex.DecodeString(cipherText[32:56])
	text, _ := hex.DecodeString(cipherText[56:])

	key := GenerateKey(password, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

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
