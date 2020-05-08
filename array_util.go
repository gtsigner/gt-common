//author: https://github.com/zhaojunlike
//date: 2019/12/11
package common

import (
    "container/list"
    "encoding/json"
    "errors"
    "math/rand"
    "strings"
    "time"
)

//转array
func ListToArray(li *list.List) []interface{} {
    var ls = make([]interface{}, li.Len())
    var ix = 0
    for e := li.Front(); e != nil; e = e.Next() {
        ls[ix] = e.Value
        ix++
    }
    return ls
}

func JSONParse(jsonStr string, t interface{}) error {
    var bytes = []byte(jsonStr)
    err := json.Unmarshal(bytes, t)
    if err != nil {
        return err
    }
    return nil
}

func JSONStringify(v interface{}) (string, error) {
    var bytes, err = json.Marshal(v)
    return string(bytes[:]), err
}

// 打印列表
func LoopList(l list.List, call func(index int, e *list.Element)) {
    var line = 0
    for e := l.Front(); e != nil; e = e.Next() {
        call(line, e)
        line++
    }
}

//随机取集合中的一个元素
func Sample(arr []interface{}, vptr interface{}) error {
    if len(arr) <= 0 {
        return errors.New("not found")
    }
    //随机
    var count = len(arr)
    var v = 0
    if count > 1 {
        rand.Seed(time.Now().UnixNano())
        v = rand.Intn(count - 1)
    }
    vptr = arr[v]
    return nil
}

func ArrayRemoveNullString(arr []string) []string {
    var newArr = make([]string, 0)
    for _, v := range arr {
        if v != "" {
            newArr = append(newArr, v)
        }
    }
    return newArr
}

type CompareFunc func(interface{}, interface{}) int

func ArrayIndexOf(a []interface{}, e interface{}, cmp CompareFunc) int {
    n := len(a)
    var i int = 0
    for ; i < n; i++ {
        if cmp(e, a[i]) == 0 {
            return i
        }
    }
    return -1
}

func StringArrayIndexOf(a []string, e string) int {
    n := len(a)
    var i = 0
    for ; i < n; i++ {
        if strings.Compare(e, a[i]) == 0 {
            return i
        }
    }
    return -1
}

func StringArrayShuffle(arr []string) []string {
    for i := len(arr) - 1; i > 0; i-- {
        num := rand.Intn(i + 1)
        arr[i], arr[num] = arr[num], arr[i]
    }
    return arr
}
