package controller

import (
	"context"
	"log"

	"github.com/Chengxufeng1994/water-ball-missions/webframework/internal/application/service"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/internal/constant"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/water"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/water/ioc"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (r *RegisterRequest) Key() string {
	return "RegisterRequest"
}

type RegisterResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (r *RegisterResponse) Key() string {
	return "RegisterResponse"
}

func (u *UserController) Register(c *water.Context) {
	ctx := u.container.Scoped(context.Background())

	reqI, err := c.Bind("RegisterRequest")
	if err != nil {
		log.Println("Bind failed:", err)
		c.Status(int(water.StatusCodeBadRequest))
		return
	}

	req := reqI.(*RegisterRequest)

	userSvc := ioc.Get(ctx, constant.UserApplicationServiceKey).(service.UserApplicationService)
	newUser, err := userSvc.Register(req.Email, req.Name, req.Password)
	if err != nil {
		log.Println("Register failed:", err)
		if err.Error() == "Duplicate email" {
			c.Status(int(water.StatusCodeBadRequest))
			c.Render(int(water.StatusCodeBadRequest), "plain/text", "Duplicate email")
		} else if err.Error() == "Registration's format incorrect." {
			c.Status(int(water.StatusCodeBadRequest))
			c.Render(int(water.StatusCodeBadRequest), "plain/text", "Registration's format incorrect.")
		} else {
			c.Status(int(water.StatusCodeInternalServerError))
		}
		return
	}

	resp := &RegisterResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
		Name:  newUser.Username,
	}

	// Respond with JSON
	c.Render(int(water.StatusCodeOK), "application/json", resp)
}
