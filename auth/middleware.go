package auth

import (
	"fmt"
	"net/http"
	"time"
	
	"restful/helper"
	"restful/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// 获取token
func GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.UserName,
		// 7天后过期
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	return token.SignedString([]byte("secret"))
}

// 验证token
func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("authorization")
		if tokenStr == "" {
			helper.ResponseWithJson(w, http.StatusAccepted, helper.Response{Code: http.StatusUnauthorized, Msg: "请输入令牌"})
		} else {
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					helper.ResponseWithJson(w, http.StatusAccepted, helper.Response{Code: http.StatusUnauthorized, Msg: "not authorized"})
					return nil, fmt.Errorf("not authorization")
				}
				return []byte("secret"), nil
			})
			if !token.Valid {
				helper.ResponseWithJson(w, http.StatusAccepted, helper.Response{Code: http.StatusUnauthorized, Msg: "令牌过期"})
			} else {
				next.ServeHTTP(w, r)
			}
		}
	})
}
