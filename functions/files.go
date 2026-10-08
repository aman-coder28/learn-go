package fns

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "os"
    "os/exec"
    "runtime"

    "golang.org/x/crypto/argon2"
    "golang.org/x/term"
)

// Errors returned by Encrypt and Decrypt.
var (
    // ErrWrongPassword means the AES-GCM authentication tag failed to
    // verify: the password is wrong, or the encrypted data is damaged.
    ErrWrongPassword = errors.New("wrong password, or the encrypted data is damaged")
    // ErrNotEncrypted means the input is not in the format gops writes.
    ErrNotEncrypted = errors.New("not a gops-encrypted file")
)

// ReadPassword reads a password from the terminal without echoing it.
func ReadPassword() (string, error) {
    fmt.Fprint(os.Stderr, "Enter Password: ")

    password, err := term.ReadPassword(int(os.Stdin.Fd()))
    if err != nil {
        return "", fmt.Errorf("read password: %w", err)
    }

    return string(password), nil
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
        return args[1]
    } else {
        return fileName
    }
}

// WriteFile writes data to fileName, readable by the owner only.
func WriteFile(fileName string, data string) error {
    return os.WriteFile(fileName, []byte(data), 0600)
}

func GenerateKey(password string, salt []byte) []byte {
    return argon2.Key([]byte(password), salt, 3, 32*1024, 4, 32)
}

// Encrypt seals text with AES-256-GCM under an Argon2id key derived from
// the password, and returns hex(salt) + hex(nonce) + hex(ciphertext).
func Encrypt(password string, text string) (string, error) {
    salt := make([]byte, 16)
    if _, err := rand.Read(salt); err != nil {
        return "", fmt.Errorf("generate salt: %w", err)
    }

    key := GenerateKey(password, salt)

    block, err := aes.NewCipher(key)
    if err != nil {
        return "", fmt.Errorf("create cipher: %w", err)
    }

    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", fmt.Errorf("create GCM: %w", err)
    }

    nonce := make([]byte, aesgcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", fmt.Errorf("generate nonce: %w", err)
    }

    ciphertext := aesgcm.Seal(nil, nonce, []byte(text), nil)

    return fmt.Sprintf("%x%x%x", salt, nonce, ciphertext), nil
}

// Decrypt opens a blob produced by Encrypt.
func Decrypt(password string, cipherText string) (string, error) {
    if len(cipherText) < 56 { // hex of the 16-byte salt and 12-byte nonce
        return "", ErrNotEncrypted
    }

    salt, err := hex.DecodeString(cipherText[:32])
    if err != nil {
        return "", ErrNotEncrypted
    }

    nonce, err := hex.DecodeString(cipherText[32:56])
    if err != nil {
        return "", ErrNotEncrypted
    }

    text, err := hex.DecodeString(cipherText[56:])
    if err != nil {
        return "", ErrNotEncrypted
    }

    key := GenerateKey(password, salt)

    block, err := aes.NewCipher(key)
    if err != nil {
        return "", fmt.Errorf("create cipher: %w", err)
    }

    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", fmt.Errorf("create GCM: %w", err)
    }

    plaintext, err := aesgcm.Open(nil, nonce, text, nil)
    if err != nil {
        // The authentication tag failed: the password is wrong, or the
        // data was damaged. There is no way to tell the two apart.
        return "", ErrWrongPassword
    }

    return string(plaintext), nil
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

func ClearScreen() {
    var cmd *exec.Cmd
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/c", "cls")
    } else {
        cmd = exec.Command("clear")
    }
    cmd.Stdout = os.Stderr
    cmd.Run()
}