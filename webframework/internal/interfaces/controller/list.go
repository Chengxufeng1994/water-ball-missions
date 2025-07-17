package controller

import (
	"context"
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/webframework/internal/application/service"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/internal/constant"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/internal/interfaces/dto"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/water"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/water/ioc"
)

func (u *UserController) List(c *water.Context) {
	ctx := u.container.Scoped(context.Background())

	authHeader := c.HttpRequest.Req.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.Status(int(water.StatusCodeUnauthorized))
		c.Render(int(water.StatusCodeUnauthorized), "plain/text", "Can't authenticate who you are.")
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	userSvc := ioc.Get(ctx, constant.UserApplicationServiceKey).(service.UserApplicationService)
	_, err := userSvc.ValidateToken(token)
	if err != nil {
		c.Status(int(water.StatusCodeUnauthorized))
		c.Render(int(water.StatusCodeUnauthorized), "plain/text", "Can't authenticate who you are.")
		return
	}

	keyword := c.Query("keyword")

	appUsers, err := userSvc.List(keyword)
	if err != nil {
		c.Status(int(water.StatusCodeInternalServerError))
		return
	}

	// Convert application.dto.User to interfaces.dto.User
	var usersPtr []*dto.User
	for _, appUser := range appUsers {
		usersPtr = append(usersPtr, &dto.User{
			ID:       appUser.ID,
			Email:    appUser.Email,
			Username: appUser.Username,
		})
	}

	c.Render(int(water.StatusCodeOK), "application/json", usersPtr)
}
