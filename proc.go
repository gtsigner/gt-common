package common

import (
    "github.com/mitchellh/go-ps"
    "os"
    "os/signal"
    "strings"
)

func TerminateProc(pid int, sig os.Signal) error {
    return terminateProc(pid, sig)
}
func KillProc(pid int) error {
    p, err := os.FindProcess(pid)
    if err != nil {
        return err
    }
    return p.Kill()
}
func GetProcesses(name string, callback func(p ps.Process)) error {
    pes, _ := ps.Processes()
    for _, p := range pes {
        if strings.Index(p.Executable(), name) != -1 {
            if callback != nil {
                callback(p)
            }
        }
    }
    return nil
}

//等待程序退出
func Waiting(sig ...os.Signal) {
    ch := make(chan os.Signal)
    signal.Notify(ch, sig...)
}
