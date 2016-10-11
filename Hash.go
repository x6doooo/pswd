package pswd

import (
    "crypto/sha512"
    "golang.org/x/crypto/bcrypt"
    "encoding/base64"
)


// 将密码加盐然后sha1计算
func sha1password(password, userKey, salt []byte) []byte {
    // step1: step1result = sha512(password)
    password_sha512 := sha512.Sum512(password)
    //fmt.Println(password_sha512[:])

    // step2: step2result = step1result + userKey + salt0
    userKeyAndSalt0 := append(userKey, salt...)
    b := append(password_sha512[:], userKeyAndSalt0...)

    // 再做一次base64处理 以便能够和nodejs的编码统一 可以互相调用加密解密（便于以后替换）
    str := base64.StdEncoding.EncodeToString(b)
    return []byte(str)
}

/*
    密码加密

    @param password 用户密码
    @param userKey 用户标识
    @param salt 盐
    @param cost bcrypt的cost值

    @return
        passwordHash 加密结果
        err 错误
 */
func Hash(password, userKey, salt []byte, cost int) (passwordHash []byte, err error) {
    theContentNeedEncrypt := sha1password(password, userKey, salt)
    passwordHash, err = bcrypt.GenerateFromPassword(theContentNeedEncrypt, cost)
    return
}

/*
    密码校验

    @param password 用户密码
    @param userKey 用户标识
    @param salt 盐
    @param cost bcrypt的cost值

    @return
        error 错误
 */
func Verfiy(password, userKey, salt, passwordHash []byte) error {
    passwordSha1 := sha1password(password, userKey, salt)
    return bcrypt.CompareHashAndPassword(passwordHash, passwordSha1)
}

