// Created by Hisen at 2019-06-19.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/sony/sonyflake"
	"net/http"
)

var (
	snowFlake *sonyflake.Sonyflake
)

const (
	IntervalError = 1001
	Success       = 0
)

type ResponseData struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       map[string]interface{}
}

func initSnowFlake() (err error) {
	setting := sonyflake.Settings{}
	setting.MachineID = func() (u uint16, e error) {
		return 1, nil
	}
	snowFlake = sonyflake.NewSonyflake(setting)
	return
}
func ErrResponse(w http.ResponseWriter) {
	var rd ResponseData
	rd.StatusCode = IntervalError
	rd.Message = "内部错误"
	r, _ := json.Marshal(rd)
	w.Write(r)
}

func SuccResponse(id uint64, w http.ResponseWriter) {
	var rd ResponseData
	rd.StatusCode = Success
	rd.Message = "success"
	rd.Data = make(map[string]interface{})
	rd.Data["id"] = id
	r, _ := json.Marshal(rd)
	w.Write(r)
}
func idGen() (id uint64, err error) {
	id, err = snowFlake.NextID()
	return
}
func idGenHandle(w http.ResponseWriter, r *http.Request) {
	id, err := idGen()
	fmt.Println(r.RemoteAddr, id)
	if err != nil {
		ErrResponse(w)
	} else {
		SuccResponse(id, w)
	}
}
func main() {
	initSnowFlake()
	http.HandleFunc("/id", idGenHandle)
	http.ListenAndServe(":8888", nil)
}
