package common

import (
    "archive/zip"
    "io"
    "os"
    "path"
)

//zip 文件解压
//@param name 解压文件名称
//@param dest 存放到文件夹
func ZipDeCompress(name string, dest string) error {
    zipFile, err := zip.OpenReader(name)
    if err != nil {
        return err
    }
    defer zipFile.Close()
    for _, innerFile := range zipFile.File {
        info := innerFile.FileInfo()
        toName := path.Join(dest, innerFile.Name)
        if info.IsDir() {
            err = os.MkdirAll(toName, os.ModePerm)
            if err != nil {
                return err
            }
            continue
        }
        //如果是文件
        srcFile, err := innerFile.Open()
        if err != nil {
            continue
        }
        //拷贝文件信息
        func() {
            defer srcFile.Close()
            newFile, err := os.Create(toName)
            if err != nil {
                return
            }
            defer newFile.Close()
            _, _ = io.Copy(newFile, srcFile)
            newFile.Chmod(0777)
            //拷贝完成后，需要给一个权限
        }()

    }
    return nil
}
