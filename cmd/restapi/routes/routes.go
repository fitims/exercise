package routes

import (
	"github.com/fitims/exercise/cmd/restapi/controllers"
	"github.com/fitims/exercise/cmd/restapi/middleware/authentication"
	"github.com/fitims/exercise/cmd/restapi/params"
)

const (
	Register = "/user"
	Login    = "/login"
	Maze     = "/maze"
	Solution = "/maze/:" + params.MazeId + "/solution"
)

func SetUserRoutes(grp GroupWrapper, c controllers.UserController) {
	grp.POST(Register, c.Register)
	grp.POST(Login, c.Login)
}

func SetMazeRoutes(grp GroupWrapper, c controllers.MazeController, a authentication.Authenticator) {
	grp.POST(Maze, c.CreateMaze, a.AuthenticateUser)
	grp.GET(Maze, c.GetMazes, a.AuthenticateUser)
	grp.GET(Solution, c.GetSolution, a.AuthenticateUser)
}
