package fs

import "os"

func PathCheckAndCreate(s string, perm os.FileMode) error {
    ext, err := PathExists(s)
    if !ext {
        return os.MkdirAll(s, perm)
    }
    return err
}

func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}
