package common

import (
	"fmt"
)

func Logger(arr ...interface{}) {
	var t = []interface{}{"[" + TimeNow() + "] INFO:"}
	var log = append(t, arr...)
	fmt.Println(log...)
}
