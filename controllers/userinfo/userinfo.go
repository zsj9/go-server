package userinfo

import (
	"encoding/json"
	"net/http"
	// "restful/auth"
	"restful/helper"
	"restful/models"
	
	// "github.com/globalsign/mgo/bson"
)

const (
	db         = "blog"
	collection = "userinfos"
)

// 注册
func add_userinfo(w http.ResponseWriter, r *http.Request) {
	var userinfo models.Userinfo
	err := json.NewDecoder(r.Body).Decode(&userinfo)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest,	helper.Response{Code: http.StatusBadRequest, Msg: "theow error"})
		return
	}
	if userinfo.Name == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest,	helper.Response{Code: http.StatusBadRequest, Msg: "请输入用户名"})
		return
	}
	if userinfo.User_id == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest,	helper.Response{Code: http.StatusBadRequest, Msg: "注册失败"})
		return
	}
	err = models.Insert(db, collection, userinfo)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusInternalServerError,	helper.Response{Code: http.StatusInternalServerError, Msg: "注册失败"})
	} else {
		helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Data: userinfo })
	}
}
