package pswd

import (
    "crypto/cipher"
    "crypto/aes"
    "strings"
)

type Token struct {
    key      []byte
    iv       []byte
    salt     []byte
    keyBlock cipher.Block
}

func NewToken(key, iv, salt []byte) *Token {
    keyBlock, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }
    iv = iv[:aes.BlockSize]
    return &Token{
        key: key,
        iv: iv,
        salt: salt,
        keyBlock: keyBlock,
    }
}

func(me *Token) Encrypt(content []byte) []byte {
    encrypter := cipher.NewCFBEncrypter(me.keyBlock, me.iv)
    encrypted := make([]byte, len(content))
    encrypter.XORKeyStream(encrypted, content)
    return encrypted
}

func(me *Token) Decrypt(content []byte) []byte {
    decrypter := cipher.NewCFBDecrypter(me.keyBlock, me.iv)
    decrypted := make([]byte, len(content))
    decrypter.XORKeyStream(decrypted, content)
    if strings.Contains(string(decrypted), string(me.salt)) {
        str := strings.Replace(string(decrypted), string(me.salt), "", 1)
        return []byte(str)
    }
    return []byte("")
}