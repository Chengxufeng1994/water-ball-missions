package controller

import (
	"context"
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/webframework/internal/application/service"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/internal/constant"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/water"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/water/ioc"
)

type RenameRequest struct {
	Name string `json:"name"`
}

func (RenameRequest) Key() string {
	return "RenameRequest"
}

type RenameResponse struct {
	Name string `json:"name"`
}

func (RenameResponse) Key() string {
	return "RenameResponse"
}

func (u *UserController) Rename(c *water.Context) {
	ctx := u.container.Scoped(context.Background())

	authHeader := c.HttpRequest.Req.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.Status(int(water.StatusCodeUnauthorized))
		c.Render(int(water.StatusCodeUnauthorized), "plain/text", "Can't authenticate who you are.")
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	userSvc := ioc.Get(ctx, constant.UserApplicationServiceKey).(service.UserApplicationService)
	authenticatedUserID, err := userSvc.ValidateToken(token)
	if err != nil {
		c.Status(int(water.StatusCodeUnauthorized))
		c.Render(int(water.StatusCodeUnauthorized), "plain/text", "Can't authenticate who you are.")
		return
	}

	id := c.HttpRequest.Params["id"] // Get ID from path parameter
	if id == "" {
		c.Status(int(water.StatusCodeBadRequest))
		return
	}

	if authenticatedUserID != id {
		c.Status(int(water.StatusCodeForbidden))
		c.Render(int(water.StatusCodeForbidden), "plain/text", "Forbidden")
		return
	}

	reqI, err := c.Bind("RenameRequest")
	if err != nil {
		c.Status(int(water.StatusCodeBadRequest))
		return
	}

	req := reqI.(*RenameRequest)

	newName, err := userSvc.Rename(id, req.Name)
	if err != nil {
		if err.Error() == "Name's format invalid." {
			c.Status(int(water.StatusCodeBadRequest))
			c.Render(int(water.StatusCodeBadRequest), "plain/text", "Name's format invalid.")
		} else {
			c.Status(int(water.StatusCodeInternalServerError))
		}
		return
	}

	resp := &RenameResponse{Name: newName}
	c.Render(int(water.StatusCodeOK), "application/json", resp)
}
