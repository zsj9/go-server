package user

import (
	// "fmt"
	"encoding/json"
	"net/http"
	"restful/auth"
	"restful/helper"
	"restful/models"

	"time"
	"github.com/globalsign/mgo/bson"
)

// 登录
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest,	helper.Response{Code: http.StatusBadRequest, Msg: "bad params"})
		return
	}
	// 查询账号密码是否正确
	exist := models.IsExist("users", bson.M{"username": user.UserName, "password": user.Password})
	if exist {
		// 登入往user表插入token
		token, _ := auth.GenerateToken(&user)
		err := models.FindOne("users", bson.M{"username": user.UserName}, nil, &user)
		// fmt.Println(user)
		user.Token = token
		err = models.Update("users", bson.M{"username": user.UserName}, user)
		if err != nil {
			helper.ResponseWithJson(w, http.StatusAccepted,	helper.Response{Code: http.StatusAccepted, Msg: "令牌失效，登入失败"})
			return
		}
		helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Data: user })
	} else {
		helper.ResponseWithJson(w, http.StatusAccepted, helper.Response{Code: http.StatusAccepted, Msg: "请输入正确的账号密码"})
	}
}

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
	exist := models.IsExist("users", bson.M{"username": user.UserName})
	if exist {
		helper.ResponseWithJson(w, http.StatusAccepted,	helper.Response{Code: http.StatusAccepted, Msg: "该账号已被注册过"})
		return
	}
	user.Id = bson.NewObjectId()
	user.Created_at = time.Now().UTC().Format(time.RFC3339)
	err = models.Insert("users", user)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusAccepted,	helper.Response{Code: http.StatusAccepted, Msg: "注册失败"})
	} else {
		helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Data: user })
	}
}
