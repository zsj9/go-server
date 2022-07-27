package userinfo

import (
	"encoding/json"
	"net/http"
	"restful/helper"
	"restful/models"

	"github.com/globalsign/mgo/bson"
)


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
	err = models.FindOne(collection, bson.M{"user_id": userinfo.User_id}, nil, &userinfo)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusInternalServerError,	helper.Response{Code: http.StatusInternalServerError, Msg: "失败"})
	} else {
		helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Data: userinfo })
	}
}