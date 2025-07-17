package main

import (
	"log"

	"github.com/Chengxufeng1994/water-ball-missions/webframework/internal/application/service"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/internal/constant"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/internal/interfaces/controller"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/water"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/water/ioc"
	"github.com/Chengxufeng1994/water-ball-missions/webframework/water/serdes"
)

func main() {
	port := ":8000"
	log.Printf("server running at %s", port)
	app, err := water.New(port)
	if err != nil {
		log.Fatal(err)
	}
	app.AddSerDesPlugin(serdes.NewXmlSerDesHandler())

	reg := app.GetRegistry()
	reg.Register(&controller.RegisterRequest{})
	reg.Register(&controller.RegisterResponse{})
	reg.Register(&controller.LoginRequest{})
	reg.Register(&controller.LoginResponse{})
	reg.Register(&controller.RenameRequest{})
	reg.Register(&controller.RenameResponse{})

	container := ioc.New()

	container.AddScoped(constant.UserApplicationServiceKey, func(c ioc.Container) (any, error) {
		return service.NewUserApplicationService(), nil
	})

	container.AddSingleton(constant.UserControllerKey, func(c ioc.Container) (any, error) {
		return controller.NewUserController(c), nil
	})

	userCtrl := container.Get(constant.UserControllerKey).(*controller.UserController)

	app.AddRoute(water.MethodPOST, "/api/users", water.HttpHandlerFunc(userCtrl.Register))
	app.AddRoute(water.MethodGET, "/api/users", water.HttpHandlerFunc(userCtrl.Login))
	app.AddRoute(water.MethodPATCH, "/api/users/:id", water.HttpHandlerFunc(userCtrl.Rename))
	app.AddRoute(water.MethodGET, "/api/users", water.HttpHandlerFunc(userCtrl.List))

	app.Launch()
}
