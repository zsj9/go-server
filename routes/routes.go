package routes

import (
	"fmt"
	"net/http"
	"restful/api/user"
	"restful/api/userinfo"
	"restful/auth"
	"github.com/gorilla/mux"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	register("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome!n")
	}, nil)
	register("POST", "/api/user/register", user.Register, nil)
	register("POST", "/api/user/login", user.Login, nil)

	register("POST", "/api/userinfo/add", userinfo.Add_userinfo, nil)
	register("POST", "/api/userinfo/find", userinfo.Find_userinfo, auth.TokenMiddleware)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		r := router.Methods(route.Method).
			Path(route.Pattern)
		if route.Middleware != nil {
			r.Handler(route.Middleware(route.Handler))
		} else {
			r.Handler(route.Handler)
		}
	}
	return router
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
