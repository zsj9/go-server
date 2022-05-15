package userinfo

import (
	"encoding/json"
	"net/http"
	"restful/helper"
	"restful/models"
	
	"time"

	"github.com/globalsign/mgo/bson"
)

const (
	db         = "blog"
	collection = "userinfos"
)

// 注册
func Add_userinfo(w http.ResponseWriter, r *http.Request) {
	var userinfo models.Userinfo
	err := json.NewDecoder(r.Body).Decode(&userinfo)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest,	helper.Response{Code: http.StatusBadRequest, Msg: "theow error"})
		return
	}
	if userinfo.Name == "" {
		helper.ResponseWithJson(w, http.StatusAccepted,	helper.Response{Code: http.StatusAccepted, Msg: "请输入用户名"})
		return
	}
	if userinfo.User_id == "" {
		helper.ResponseWithJson(w, http.StatusAccepted,	helper.Response{Code: http.StatusAccepted, Msg: "注册失败"})
		return
	}
	userinfo.Id = bson.NewObjectId()
	userinfo.Created_at = time.Now().UTC().Format(time.RFC3339)
	err = models.Insert(db, collection, userinfo)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusInternalServerError,	helper.Response{Code: http.StatusInternalServerError, Msg: "注册失败"})
	} else {
		helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Data: userinfo })
	}
}

// 获取用户信息
func Find_userinfo (w http.ResponseWriter, r *http.Request) {
	var userinfo models.Userinfo
	err := json.NewDecoder(r.Body).Decode(&userinfo)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest,	helper.Response{Code: http.StatusBadRequest, Msg: "theow error"})
	}
	if userinfo.User_id == "" {
		helper.ResponseWithJson(w, http.StatusAccepted,	helper.Response{Code: http.StatusAccepted, Msg: "请输入用户id"})
		return
	}
	err = models.FindOne(db, collection, bson.M{"user_id": userinfo.User_id}, nil, &userinfo)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusInternalServerError,	helper.Response{Code: http.StatusInternalServerError, Msg: "失败"})
	} else {
		helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Data: userinfo })
	}
}