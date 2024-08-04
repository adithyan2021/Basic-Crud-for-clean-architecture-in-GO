package routes

import (
	"github.com/shakezidin/internal/server"
	"github.com/shakezidin/internal/user/delivery"
)

type UserRoutesstruct struct {
	Server *server.ServerStruct
	User   delivery.UserUserCase
}

func (u *UserRoutesstruct) UserRoutes() {
	u.Server.engine.POST("/signup", u.User.RegisterUserHandler)
	u.Server.engine.POST("/login", u.User.LoginUserHandler)
}

func NewUserInit(server *server.ServerStruct, user delivery.UserUserCase) *UserRoutesstruct {
	return &UserRoutesstruct{
		Server: server,
		User:   user,
	}
}
