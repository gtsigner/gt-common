package machine

import (
    "crypto/md5"
    "crypto/rand"
    "encoding/base64"
    "encoding/hex"
    "errors"
    "net"
    "os/exec"
    "regexp"
    "strings"
)

func getMACAddress() (string, error) {
    netInterfaces, err := net.Interfaces()
    if err != nil {
        panic(err.Error())
    }
    mac, macerr := "", errors.New("无法获取到正确的MAC地址")
    for i := 0; i < len(netInterfaces); i++ {
        if (netInterfaces[i].Flags&net.FlagUp) != 0 && (netInterfaces[i].Flags&net.FlagLoopback) == 0 {
            addrs, _ := netInterfaces[i].Addrs()
            for _, address := range addrs {
                ipnet, ok := address.(*net.IPNet)
                if ok && ipnet.IP.IsGlobalUnicast() {
                    // 如果IP是全局单拨地址，则返回MAC地址
                    mac = netInterfaces[i].HardwareAddr.String()
                    return mac, nil
                }
            }
        }
    }
    return mac, macerr
}

type cpuInfo struct {
    CPU        int32  `json:"cpu"`
    VendorID   string `json:"vendorId"`
    PhysicalID string `json:"physicalId"`
}

type win32Processor struct {
    Manufacturer string
    ProcessorID  *string
}

/**
 * 获取电脑CPUId
 */
func GetCpuId() string {
    cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return ""
    }
    str := string(out)
    //匹配一个或多个空白符的正则表达式
    reg := regexp.MustCompile("\\s+")
    str = reg.ReplaceAllString(str, "")
    return strings.TrimSpace(str[11:])
}

//生成32位md5字串
func GetMd5String(s string, upper bool, half bool) string {
    h := md5.New()
    h.Write([]byte(s))
    result := hex.EncodeToString(h.Sum(nil))
    if upper == true {
        result = strings.ToUpper(result)
    }
    if half == true {
        result = result[8:24]
    }
    return result
}

//利用随机数生成Guid字串
func UniqueId() string {
    b := make([]byte, 48)
    if _, err := rand.Read(b); err != nil {
        return ""
    }
    return GetMd5String(base64.URLEncoding.EncodeToString(b), true, false)
}
