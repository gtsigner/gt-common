package common

import (
    "gopkg.in/yaml.v2"
)

//解析
func YamlParse(data []byte, t interface{}) error {
    err := yaml.Unmarshal(data, t)
    if err != nil {
        return err
    }
    return nil
}

//序列化
func YamlStringify(data interface{}) (string, error) {
    d, err := yaml.Marshal(data)
    if err != nil {
        return "", err
    }
    return string(d), err
}
