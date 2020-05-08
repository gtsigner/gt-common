package common

import (
    "crypto/md5"
    "crypto/sha1"
    "encoding/hex"
    "fmt"
)

func Md5Encrypt(str string) string {
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}

func Sha1Encode(str string) string {
    h := sha1.New()
    h.Write([]byte(str))
    l := fmt.Sprintf("%x", h.Sum(nil))
    return l
}
