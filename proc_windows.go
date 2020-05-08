// +build windows

package common

import (
    "golang.org/x/sys/windows"
    "os"
    "syscall"
)

func terminateProc(pid int, sig os.Signal) error {
    dll, err := windows.LoadDLL("kernel32.dll")
    if err != nil {
        return err
    }
    defer dll.Release()
    f, err := dll.FindProc("AttachConsole")
    if err != nil {
        return err
    }
    r1, _, err := f.Call(uintptr(pid))
    if r1 == 0 && err != syscall.ERROR_ACCESS_DENIED {
        return err
    }

    f, err = dll.FindProc("SetConsoleCtrlHandler")
    if err != nil {
        return err
    }
    r1, _, err = f.Call(0, 1)
    if r1 == 0 {
        return err
    }
    f, err = dll.FindProc("GenerateConsoleCtrlEvent")
    if err != nil {
        return err
    }
    var v = windows.CTRL_C_EVENT
    if sig == syscall.SIGKILL {
        v = windows.CTRL_SHUTDOWN_EVENT
    }
    r1, _, err = f.Call(uintptr(v), uintptr(pid)) //关闭指定CTRL+C进程
    if r1 == 0 {
        return err
    }
    return nil
}
