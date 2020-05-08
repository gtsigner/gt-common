// +build !windows

package common

import (
    "os"
)

func terminateProc(pid int, sig os.Signal) error {
    p, err := os.FindProcess(pid)
    if err != nil {
        return err
    }
    return p.Signal(sig)
}

