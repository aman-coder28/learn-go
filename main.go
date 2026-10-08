package main

import (
    "fmt"
    "os"

    fns "gops/functions"
)

func main() {
    if err := run(); err != nil {
        fmt.Fprintln(os.Stderr, "gops:", err)
        os.Exit(1)
    }
}

func run() error {
    fileName := fns.GetArgs()

    content, err := fns.LoadFile(fileName)
    if err != nil {
        return err
    }

    password, err := fns.ReadPassword()
    if err != nil {
        return err
    }

    fns.ClearScreen()

    if json, ok := fns.IsJson(string(content)); ok {
        decrypted, err := fns.Decrypt(password, json.Data)
        if err != nil {
            return err
        }

        _, err = os.Stdout.WriteString(decrypted)
        return err
    }

    encrypted, err := fns.Encrypt(password, string(content))
    if err != nil {
        return err
    }

    data, err := fns.StringifyData(fns.Input{Data: encrypted})
    if err != nil {
        return err
    }

    return fns.WriteFile(fileName, string(data))
}