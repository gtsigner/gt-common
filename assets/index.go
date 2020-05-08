package assets

import (
    "github.com/rakyll/statik/fs"
    "io/ioutil"
    "net/http"
    "os"
    "path"
)

var (
    // store assets files in memory by statik
    FileSystem http.FileSystem
    // if prefix is not empty, we get file content from disk
    prefixPath string
)

// if path is empty, load assets in memory
// or set FileSystem using disk files
func Load(path string) (err error) {
    prefixPath = path
    if prefixPath != "" {
        FileSystem = http.Dir(prefixPath)
        return nil
    } else {
        FileSystem, err = fs.New()
    }
    return err
}

func ReadFile(file string) (content string, err error) {
    if prefixPath == "" {
        file, err := FileSystem.Open(path.Join("/", file))
        if err != nil {
            return content, err
        }
        defer file.Close()
        buf, err := ioutil.ReadAll(file)
        if err != nil {
            return content, err
        }
        content = string(buf)
    } else {
        file, err := os.Open(path.Join(prefixPath, file))
        if err != nil {
            return content, err
        }
        defer file.Close()
        buf, err := ioutil.ReadAll(file)
        if err != nil {
            return content, err
        }
        content = string(buf)
    }
    return content, err
}

func CopyFile(source string, target string) (err error) {
    var file http.File
    if prefixPath == "" {
        file, err = FileSystem.Open(path.Join("/", source))
        if err != nil {
            return err
        }

    } else {
        file, err = os.Open(path.Join(prefixPath, source))
        if err != nil {
            return err
        }
    }
    defer file.Close()
    buf, err := ioutil.ReadAll(file)
    if err != nil {
        return err
    }
    //write
    err = ioutil.WriteFile(target, buf, os.ModePerm)
    return err
}
