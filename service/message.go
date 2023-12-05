package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MsgHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	if r.Method == http.MethodPost {
		res.Data = "hello,world"
		msg, err := json.Marshal(res)
		if err != nil {
			fmt.Fprint(w, "内部错误")
			return
		}
		w.Write(msg)
	}
}
