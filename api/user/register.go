package user

import (
	// "fmt"
	"encoding/json"
	"net/http"
	"restful/helper"
	"restful/models"

	"time"
	"github.com/globalsign/mgo/bson"
)

// 注册
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest,	helper.Response{Code: http.StatusBadRequest, Msg: "theow error"})
		return
	}
	if user.UserName == "" || user.Password == "" {
		helper.ResponseWithJson(w, http.StatusAccepted,	helper.Response{Code: http.StatusAccepted, Msg: "请输入账号或者密码"})
		return
	}
	// 根据username去查看是否存在数据
	exist := models.IsExist(collection, bson.M{"username": user.UserName})
	if exist {
		helper.ResponseWithJson(w, http.StatusAccepted,	helper.Response{Code: http.StatusAccepted, Msg: "该账号已被注册过"})
		return
	}
	user.Id = bson.NewObjectId()
	user.Created_at = time.Now().UTC().Format(time.RFC3339)
	err = models.Insert(collection, user)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusAccepted,	helper.Response{Code: http.StatusAccepted, Msg: "注册失败"})
	} else {
		helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Data: user })
	}
}
