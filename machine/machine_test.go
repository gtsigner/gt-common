package machine

import (
    "fmt"
    "testing"
)

func TestUniqueId(t *testing.T) {

}

func TestGetCpuId(t *testing.T) {
    var id = GetCpuId()
    var id2, err = GetMachineGuid()
    fmt.Print(id)
    fmt.Println(id2, err)
}
