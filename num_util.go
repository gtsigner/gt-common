package common

import (
    "math/big"
    "regexp"
)

var reg, _ = regexp.Compile(`x|(p.*\d+)`)

func Float64ToString16(n float64) string {
    var str = new(big.Float).SetFloat64(n).Text('p', 0)
    return reg.ReplaceAllString(str, "")
}
