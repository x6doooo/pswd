package pswd

import (
    "testing"
    "fmt"
    "encoding/base64"
)

func TestNewToken(t *testing.T) {
    key := []byte("3zTvzr3p67VC61jmV54rIYu1545x4TlY")
    iv := []byte("abc1efghigklmnop")
    salt := []byte("Hi,Money!")
    token := NewToken(key, iv, salt)

    s := []byte("17|ios|127.0.0.1|1475976745")
    es := token.Encrypt(s)
    bs := token.Decrypt(es)

    fmt.Println(string(s), base64.StdEncoding.EncodeToString(es), string(bs))
    if string(s) != string(bs) {
        t.Error("nnnnnnn")
    }
}
