package common

import (
    "fmt"
    "testing"
)

func TestGetTodayTimestamp(t *testing.T) {
    var tt = GetTodayTimestamp()
    fmt.Println(tt)
}
