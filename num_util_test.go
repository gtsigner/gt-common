package common

import (
    "fmt"
    "testing"
)

func TestDecHex(t *testing.T) {
    //0.cf3862b52251 0.8094541256162451
    var val = 0.8094541256162451
    fmt.Println(Float64ToString16(val))
}
