package fns

import (
    "errors"
    "os"
    "path/filepath"
    "runtime"
    "strings"
    "testing"
)

func TestEncryptDecryptRoundTrip(t *testing.T) {
    const password = "correct horse battery staple"
    const secret = "GITHUB_TOKEN=ghp_example\nDB_PASSWORD=hunter2"

    blob, err := Encrypt(password, secret)
    if err != nil {
        t.Fatalf("Encrypt: %v", err)
    }

    got, err := Decrypt(password, blob)
    if err != nil {
        t.Fatalf("Decrypt: %v", err)
    }
    if got != secret {
        t.Errorf("round trip changed the plaintext:\n got  %q\n want %q", got, secret)
    }
}

func TestEncryptUniqueBlobs(t *testing.T) {
    // Same plaintext twice must produce different blobs: fresh salt
    // and nonce every time.
    first, err := Encrypt("pw", "same secret")
    if err != nil {
        t.Fatalf("Encrypt: %v", err)
    }
    second, err := Encrypt("pw", "same secret")
    if err != nil {
        t.Fatalf("Encrypt: %v", err)
    }
    if first == second {
        t.Error("Encrypt produced identical blobs; salt/nonce are not random")
    }
}

func TestDecryptWrongPassword(t *testing.T) {
    blob, err := Encrypt("right password", "secret")
    if err != nil {
        t.Fatalf("Encrypt: %v", err)
    }

    if _, err := Decrypt("wrong password", blob); !errors.Is(err, ErrWrongPassword) {
        t.Errorf("Decrypt with wrong password: got %v, want ErrWrongPassword", err)
    }
}

func TestDecryptTruncatedBlob(t *testing.T) {
    blob, err := Encrypt("pw", "some secret text")
    if err != nil {
        t.Fatalf("Encrypt: %v", err)
    }

    // A cut blob has valid salt and nonce, but the GCM tag cannot verify.
    if _, err := Decrypt("pw", blob[:len(blob)-4]); !errors.Is(err, ErrWrongPassword) {
        t.Errorf("Decrypt of truncated blob: got %v, want ErrWrongPassword", err)
    }
}

func TestDecryptRejectsBadInput(t *testing.T) {
    tests := []struct {
        name string
        blob string
    }{
        {"empty", ""},
        {"too short", "abcd"},
        {"not hex", strings.Repeat("z", 60)},
        {"salt not hex", "zz" + strings.Repeat("a", 30)},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if _, err := Decrypt("pw", tt.blob); !errors.Is(err, ErrNotEncrypted) {
                t.Errorf("Decrypt(%q): got %v, want ErrNotEncrypted", tt.name, err)
            }
        })
    }
}

func TestEnvelopeRoundTrip(t *testing.T) {
    blob, err := Encrypt("pw", "secret")
    if err != nil {
        t.Fatalf("Encrypt: %v", err)
    }

    data, err := StringifyData(Input{Data: blob})
    if err != nil {
        t.Fatalf("StringifyData: %v", err)
    }

    if !IsJson(string(data)) {
        t.Fatal("IsJson does not recognize a written envelope")
    }

    parsed, err := ParseJson(data)
    if err != nil {
        t.Fatalf("ParseJson: %v", err)
    }

    if _, err := Decrypt("pw", parsed.Data); err != nil {
        t.Errorf("Decrypt through the envelope: %v", err)
    }
}

func TestWriteFileIsOwnerOnly(t *testing.T) {
    path := filepath.Join(t.TempDir(), "secret.env")

    if err := WriteFile(path, "data"); err != nil {
        t.Fatalf("WriteFile: %v", err)
    }

    loaded, err := LoadFile(path)
    if err != nil {
        t.Fatalf("LoadFile: %v", err)
    }
    if string(loaded) != "data" {
        t.Errorf("loaded %q, want %q", loaded, "data")
    }

    if runtime.GOOS != "windows" {
        info, err := os.Stat(path)
        if err != nil {
            t.Fatalf("Stat: %v", err)
        }
        if got := info.Mode().Perm(); got != 0600 {
            t.Errorf("file mode is %o, want 600", got)
        }
    }
}