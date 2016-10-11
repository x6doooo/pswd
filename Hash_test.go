package pswd

import (
    "testing"
)

var (
    password = []byte("test_password")
    userKey = RandBytes(5)
    salt = RandBytes(10)
    cost = 10
)

func TestHash(t *testing.T) {
    //_, e := Hash(password, userKey, salt, cost)
    //if e != nil {
    //    t.Error(e)
    //}
    r0, e := Hash(password, userKey, salt, cost)
    if e != nil {
        t.Error(e)
    }
    e = Verfiy(password, userKey, salt, r0)
    if e != nil {
        t.Error(e)
    }

    r1, e := Hash(password, userKey, salt, cost)
    if e != nil {
        t.Error(e)
    }
    e = Verfiy(password, userKey, salt, r1)
    if e != nil {
        t.Error(e)
    }
}

func BenchmarkHash(b *testing.B) {
    r0, _ := Hash(password, userKey, salt, cost)
    for i := 0; i < b.N; i++ {
        Verfiy(password, userKey, salt, r0)
    }
}
