package pswd

import (
    "time"
    "math/rand"
)

// @see http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const onlyNumber = "0123456789"
const (
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randBytes(n int, base string) []byte {
    src := rand.NewSource(time.Now().UnixNano())
    b := make([]byte, n)
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(base) {
            b[i] = base[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }
    return b
}

/*
    生成随机字符串
 */
func RandBytes(n int) []byte {
    return randBytes(n, letterBytes)
}

func RandNumber(n int) []byte {
    return randBytes(n, onlyNumber)
}
