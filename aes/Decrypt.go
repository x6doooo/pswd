package aes

import (
    "encoding/base64"
    cryptoAes "crypto/aes"
    "crypto/cipher"
    "fmt"
)

/*
    aes解密
 */
func Decrypt(cipherData, key []byte) (data []byte, err error) {

    dctLength := base64.URLEncoding.DecodedLen(len(cipherData))
    dst := make([]byte, dctLength)
    _, err = base64.URLEncoding.Decode(dst, cipherData)
    if err != nil {
        return
    }
    if len(dst) < cryptoAes.BlockSize {
        err = fmt.Errorf("cipherText too short. It decodes to %v bytes but the minimum length is 16", len(dst))
        return
    }
    data, err = decryptAES(key, dst)
    if err != nil {
        return
    }

    return
}

func decryptAES(key, data []byte) ([]byte, error) {
    iv := data[:cryptoAes.BlockSize]
    data = data[cryptoAes.BlockSize:]

    block, err := cryptoAes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    stream := cipher.NewCFBDecrypter(block, iv)

    stream.XORKeyStream(data, data)
    return data, nil
}
