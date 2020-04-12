package main

import (
	//"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	//"../go_iot_common"
)

type Http struct {
}

func (self *Http) HttpHandle(rw http.ResponseWriter, req *http.Request) {
	var res MojiResp
	res.ResponseHeader.RspTime = int64(time.Now().Unix())
	res.ResponseHeader.ProcessCode = "weatherAlert"
	if req.Method != "POST" {
		res.Msgbody.Code = -1
		res.Msgbody.Message = "not post"
		by_data, _ := json.Marshal(res)
		rw.Write(by_data)
		return
	}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil || len(data) <= 0 {
		if err != nil {
			log.Println("post err", err)
			res.Msgbody.Code = -1
			res.Msgbody.Message = "Faild"
			by_data, _ := json.Marshal(res)
			rw.Write(by_data)
			return
		}
		log.Println(string(data))
		res.Msgbody.Code = -1
		res.Msgbody.Message = "Faild"
		by_data, _ := json.Marshal(res)
		rw.Write(by_data)
		return
	}
	//log.Println(len(data))
	log.Println(string(data))
	var req_data MojiAlertData
	err = json.Unmarshal(data, &req_data)
	if err != nil {
		log.Println(err)
		//log.Println(string(data))
		log.Println("invalid body")
		res.Msgbody.Code = -1
		res.Msgbody.Message = "Faild"
		by_data, _ := json.Marshal(res)
		rw.Write(by_data)
		return
	}
	var cal_orig string = fmt.Sprintf("%s%s%s%s", req_data.RequestHeader.AppId, req_data.RequestHeader.ProcessCode, req_data.RequestHeader.Timestamp, "06bc436cf7db4c0d950e3342413d7a47")
	h := md5.New()
	h.Write([]byte(cal_orig))
	key := hex.EncodeToString(h.Sum(nil))
	fmt.Println(key)
	if req_data.RequestHeader.Sign != key {
		res.Msgbody.Code = 101
		res.Msgbody.Message = "鉴权失败"
		by_data, _ := json.Marshal(res)
		rw.Write(by_data)
		return
	}
	res.Msgbody.Code = 0
	res.Msgbody.Message = "操作成功"
	by_data, _ := json.Marshal(res)
	rw.Write(by_data)
	//log.Println(string(by_data))
}

func (self *Http) StartHttp() {
	fmt.Println("Running at port 9080 ...")
	http.HandleFunc("/v1/ymy", self.HttpHandle) //设置访问的路由
	err := http.ListenAndServe(":9080", nil)             //设置监听的端口
	if err != nil {
		return
	}
}

