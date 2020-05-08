//author: https://github.com/zhaojunlike
//date: 2019/12/15
package common

import (
    "net/http"
    "strconv"
    "time"
)

var (
    CnTimeFormat = "2006-01-02 15:04:05"
    MTTimeFormat = "2006-01-02T15:04:05.999999999Z"
)

//毫秒时间
func CreateTimestamp() int64 {
    return time.Now().UnixNano() / 1e6
}

func CreateUnix() int64 {
    return time.Now().Unix()
}
func TimeNow() string {
    timeStr := time.Now().Format(CnTimeFormat)
    return timeStr
}
func TimeString() string {
    timeStr := time.Now().Format(CnTimeFormat)
    return timeStr
}
func FormatTimestamp(t int64) string {
    return time.Unix(t/1e3, 0).Format(CnTimeFormat)
}
func FormatTimestampString(str string) string {
    var t, err = strconv.ParseInt(str, 10, 64)
    if err != nil {
        return ""
    }
    return time.Unix(t, 0).Format(CnTimeFormat)
}

func ParseTimeToTimestamp(str string) int64 {
    t, _ := time.ParseInLocation(CnTimeFormat, str, time.Local)
    return t.UnixNano() / 1e6
}
func DateString() string {
    timeStr := time.Now().Format("2006-01-02")
    return timeStr
}

func GetShanghaiGmtTimeString() string {
    return time.Now().UTC().Add(8 * time.Hour).Format(http.TimeFormat)
}

func GetTodayTimestamp() int64 {
    timeStr := time.Now().Format("2006-01-02")
    t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
    return t.UnixNano() / 1e6
}
