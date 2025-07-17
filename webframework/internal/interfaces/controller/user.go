package controller

import (
	"github.com/Chengxufeng1994/water-ball-missions/webframework/water/ioc"
)

type UserController struct {
	container ioc.Container
}

func NewUserController(
	c ioc.Container,
) *UserController {
	return &UserController{
		container: c,
	}
}
