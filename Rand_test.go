package pswd

import (
    "testing"
)

func TestRandBytes(t *testing.T) {
    if len(RandBytes(10)) != 10 {
        t.Error("RandBytes Error")
    }
}
