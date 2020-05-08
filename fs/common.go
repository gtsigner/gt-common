package fs

import (
    "encoding/json"
    "io/ioutil"
    "os"
)

func WriteJson(f string, data interface{}, perm os.FileMode) error {
    bfr, err := json.Marshal(data)
    if err != nil {
        return err
    }
    return ioutil.WriteFile(f, bfr, perm)
}

func ReadJson(f string, data interface{}) error {
    bytes, err := ioutil.ReadFile(f)
    if err != nil {
        return err
    }
    err = json.Unmarshal(bytes, data)
    if err != nil {
        return err
    }
    return nil
}
