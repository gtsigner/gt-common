package str

import (
    "regexp"
    "strings"
)

//MASK 字符
func MaskStr(str string, c int, v string) string {
    var l = len(str)
    if c > l {
        c = l
    }
    var ed = l - c
    str = Substr(str, 0, c)
    //追加末尾的长度
    for i := 0; i < ed; i++ {
        str += v
    }
    return str
}

func Substr(s string, start, length int) string {
    bt := []rune(s)
    if start < 0 {
        start = 0
    }
    if start > len(bt) {
        start = start % len(bt)
    }
    var end int
    if (start + length) > (len(bt) - 1) {
        end = len(bt)
    } else {
        end = start + length
    }
    return string(bt[start:end])
}

func DefaultStr(s string, def string) string {
    if s == "" {
        return def
    }
    return s
}

var regTrim, _ = regexp.Compile("[\r|\n|\t| |\t\n|\\t|\\n]")

func TrimAll(str string) string {
    str = strings.TrimSpace(str)
    str = regTrim.ReplaceAllString(str, "")
    return str
}
