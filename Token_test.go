package pswd

import (
    "testing"
    "fmt"
)

func TestNewToken(t *testing.T) {
    //key :=
    key := []byte("01h345a789t12u45")
    iv := []byte("abc1efghigklmnop")
    salt := []byte("Hi,Money!")
    token := NewToken(key, iv, salt)

    s := []byte("17|ios|127.0.0.1|1475976745")
    es := token.Encrypt(s)
    bs := token.Decrypt(es)

    fmt.Println(string(s), string(bs))
    if string(s) != string(bs) {
        t.Error("nnnnnnn")
    }
}
