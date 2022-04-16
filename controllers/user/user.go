package user

import (
	// "fmt"
	"encoding/json"
	"net/http"
	"restful/auth"
	"restful/helper"
	"restful/models"

	"github.com/globalsign/mgo/bson"
)

const (
	db         = "blog"
	collection = "users"
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
		helper.ResponseWithJson(w, http.StatusBadRequest,	helper.Response{Code: http.StatusBadRequest, Msg: "请输入账号或者密码"})
		return
	}
	// 根据username去查看是否存在数据
	exist := models.IsExist(db, collection, bson.M{"username": user.UserName})
	if exist {
		helper.ResponseWithJson(w, http.StatusBadRequest,	helper.Response{Code: http.StatusBadRequest, Msg: "该账号已被注册过"})
		return
	}
	user.Id = bson.NewObjectId()
	err = models.Insert(db, collection, user)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusInternalServerError,	helper.Response{Code: http.StatusInternalServerError, Msg: "注册失败"})
	} else {
		helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Data: user })
	}
}

// 登录
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest,	helper.Response{Code: http.StatusBadRequest, Msg: "bad params"})
	}
	// 根据username去查看是否存在数据
	exist := models.IsExist(db, collection, bson.M{"username": user.UserName})
	if exist {
		token, _ := auth.GenerateToken(&user)
		helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Data: models.JwtToken{Token: token}})
	} else {
		helper.ResponseWithJson(w, http.StatusNotFound, helper.Response{Code: http.StatusNotFound, Msg: "the user not exist"})
	}
}
