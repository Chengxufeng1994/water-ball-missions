package controller

import (
	"context"
	"log"

	"github.com/Chengxufeng1994/water-ball-missions/webframework/internal/application/service"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/internal/constant"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/water"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/water/ioc"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequest) Key() string {
	return "LoginRequest"
}

type LoginResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func (r *LoginResponse) Key() string {
	return "LoginResponse"
}

func (u *UserController) Login(c *water.Context) {
	ctx := u.container.Scoped(context.Background())

	reqI, err := c.Bind("LoginRequest")
	if err != nil {
		log.Println("Bind LoginRequest error:", err)
		c.Status(int(water.StatusCodeBadRequest))
		return
	}

	req := reqI.(*LoginRequest)

	userSvc := ioc.Get(ctx, constant.UserApplicationServiceKey).(service.UserApplicationService)
	user, token, err := userSvc.Login(req.Email, req.Password)
	if err != nil {
		log.Println("Login failed:", err)
		if err.Error() == "Credentials Invalid" {
			c.Status(int(water.StatusCodeBadRequest))
			c.Render(int(water.StatusCodeBadRequest), "plain/text", "Credentials Invalid")
		} else if err.Error() == "Login's format incorrect." {
			c.Status(int(water.StatusCodeBadRequest))
			c.Render(int(water.StatusCodeBadRequest), "plain/text", "Login's format incorrect.")
		} else {
			c.Status(int(water.StatusCodeUnauthorized))
		}
		return
	}

	resp := &LoginResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Username,
		Token: token.Token,
	}

	c.Render(int(water.StatusCodeOK), "application/json", resp)
}
