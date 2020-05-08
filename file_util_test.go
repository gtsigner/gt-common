package common

import (
    "fmt"
    "testing"
)

func TestHashFileMd5(t *testing.T) {
    var hash, er = HashFileMd5("./array_util.go")
    fmt.Println(hash, er)
}
