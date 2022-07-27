package user

import (
	// "fmt"
	"encoding/json"
	"net/http"
	"restful/auth"
	"restful/helper"
	"restful/models"

	// "time"
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
	exist := models.IsExist(collection, bson.M{"username": user.UserName, "password": user.Password})
	if exist {
		// 登入往user表插入token
		token, _ := auth.GenerateToken(&user)
		err := models.FindOne(collection, bson.M{"username": user.UserName}, nil, &user)
		// fmt.Println(user)
		user.Token = token
		err = models.Update(collection, bson.M{"username": user.UserName}, user)
		if err != nil {
			helper.ResponseWithJson(w, http.StatusAccepted,	helper.Response{Code: http.StatusAccepted, Msg: "令牌失效，登入失败"})
			return
		}
		helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Data: user })
	} else {
		helper.ResponseWithJson(w, http.StatusAccepted, helper.Response{Code: http.StatusAccepted, Msg: "请输入正确的账号密码"})
	}
}
