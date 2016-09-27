package aes

import (
    "testing"
    "pswd"
    "bytes"
)

var (
    content = []byte("wwerrwer434rr3434r4r34r34r4rrewwreaw厄尔娃儿哇哦额eaewr")
    key = pswd.RandBytes(16)
)

func TestEncrypt(t *testing.T) {
    encryptedData, err := Encrypt(content, key)
    if err != nil {
        t.Error(err)
    }
    decryptedData, err := Decrypt(encryptedData, key)
    if err != nil {
        t.Error(err)
    }

    if !bytes.Equal(content, decryptedData) {
        t.Error("aes error!!!!")
    }
}

func BenchmarkEncrypt(b *testing.B) {
    k := pswd.RandBytes(16)
    for i := 0; i < b.N; i++ {
        c := pswd.RandBytes(i)
        Encrypt(c, k)
    }
}

func BenchmarkDecrypt(b *testing.B) {
    c := pswd.RandBytes(100)
    k := pswd.RandBytes(16)
    encrypted, _ := Encrypt(c, k)
    for i := 0; i < b.N; i++ {
        Decrypt(encrypted, k)
    }
}
