package aes

import (
    cryptoAes "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "io"
    "crypto/rand"
)


/*
    aes加密
 */
func Encrypt(data []byte, key []byte) (cipherData []byte, err error) {
    encrypted, err := encryptAES(key, data)
    if err != nil {
        return
    }
    encLength := base64.URLEncoding.EncodedLen(len(encrypted))
    cipherData = make([]byte, encLength)
    base64.URLEncoding.Encode(cipherData, encrypted)
    return
}

func encryptAES(key, data []byte) ([]byte, error) {

    block, err := cryptoAes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    output := make([]byte, cryptoAes.BlockSize + len(data))
    iv := output[:cryptoAes.BlockSize]
    encrypted := output[cryptoAes.BlockSize:]

    if _, err = io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(encrypted, data)
    return output, nil
}

