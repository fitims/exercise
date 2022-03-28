package main

import (
	"fmt"
	"github.com/fitims/exercise/cmd/restapi/middleware/authentication"
	"github.com/fitims/exercise/cmd/restapi/routes"
	"github.com/fitims/exercise/cmd/restapi/routes/controllers"
	"github.com/fitims/exercise/log"
	"github.com/fitims/exercise/repo/embedded"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	DbLocation = "./db"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	repo, err := embedded.NewBitCaskRepository(DbLocation)
	if err != nil {
		log.Errorln("Error creating the embedded database. Error: ", err)
		return
	}

	userController := controllers.NewUserController(repo)
	mazeController := controllers.NewMazeController(repo)
	authenticator := authentication.NewUserAuthenticator(repo.GetUserValidator())

	routes.SetUserRoutes(e, userController)
	routes.SetMazeRoutes(e, mazeController, authenticator)

	log.Infoln("Starting the web service ...")
	log.Errorln(e.Start(fmt.Sprintf(":%d", 9090)))
}
