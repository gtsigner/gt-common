package chttp

import (
    "net/http"
    "zhaojunlike/common"
)

//响应
//TODO 修訂版本,保證Header頭排序一致
type Res struct {
    Code       string              `json:"code"`
    Message    string              `json:"message"`
    Data       interface{}         `json:"data"`
    Ok         bool                `json:"ok"`
    RespStr    string              `json:"resp_str"`
    Time       int                 `json:"time"`
    StatusCode int                 `json:"status_code"`
    Url        string              `json:"url"`
    Proto      string              `json:"proto"`
    Headers    map[string][]string `json:"headers"`
    Req        *Req                `json:"-"`
    Success    bool                `json:"success"`
    Buffer     []byte              `json:"-"`
}

func NewHttpRes() *Res {
    res := Res{Code: "0", Message: "", Data: nil, Ok: false, StatusCode: 599}
    return &res
}

func NewFailRes(msg string) *Res {
    res := NewHttpRes()
    res.Message = msg
    return res
}

func (res *Res) String() string {
    str, _ := common.JSONStringify(res)
    return str
}

func (res *Res) GetHeaders() http.Header {
    return res.Headers
}
