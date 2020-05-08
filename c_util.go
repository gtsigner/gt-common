package common

import (
    "fmt"
    "math/rand"
    "sort"
    "strings"
    "time"
)

func SortMap(mapper map[string]interface{}) map[string]interface{} {
    var keys = make([]string, 0)
    for k, _ := range mapper {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    var sorted = make(map[string]interface{}, 0)
    for _, v := range keys {
        sorted[v] = mapper[v]
    }
    return sorted
}
func MapToQueryString(mapper map[string]interface{}) string {
    var str = ""
    //根据key排序
    for k, v := range mapper {
        //判断
        str += fmt.Sprintf("&%s=%v", k, v)
    }
    return strings.Replace(str, "&", "", 1)
}

// 排序map
func MapToQueryStringSort(mapper map[string]interface{}) string {
    var str = ""
    var keys = make([]string, 0)
    for k, _ := range mapper {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    //根据key排序
    for _, k := range keys {
        //判断
        str += fmt.Sprintf("&%s=%v", k, mapper[k])
    }
    return strings.Replace(str, "&", "", 1)
}

func MapToSplit(mapper map[string]string) string {
    var str = ""
    var char = ";"
    for k, v := range mapper {
        //判断
        str += fmt.Sprintf("%s=%v%s", k, v, char)
    }
    return str
}

func UniqString(a []string) []string {
    length := len(a)

    seen := make(map[string]struct{}, length)
    j := 0

    for i := 0; i < length; i++ {
        v := a[i]

        if _, ok := seen[v]; ok {
            continue
        }

        seen[v] = struct{}{}
        a[j] = v
        j++
    }

    return a[0:j]
}

func CharCodeAt(a string, i int) int {
    return int(a[i])
}
func GetRandomStringZ2z(c int) string {
    str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    bytes := []byte(str)
    var result []byte
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i := 0; i < c; i++ {
        var lv = len(bytes)
        result = append(result, bytes[r.Intn(lv)])
    }
    return string(result)
}

//替换字符串指定字符
func ReplaceStringIndexChar(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}

