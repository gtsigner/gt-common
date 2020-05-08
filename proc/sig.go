package proc

import (
    "os"
    "os/signal"
)

//等待程序退出
func Waiting(sig ...os.Signal) {
    ch := make(chan os.Signal)
    signal.Notify(ch, sig...)
}
