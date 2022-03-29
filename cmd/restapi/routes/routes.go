package routes

import (
	controllers2 "github.com/fitims/exercise/cmd/restapi/controllers"
	"github.com/fitims/exercise/cmd/restapi/middleware/authentication"
)

const (
	MazeId             = ":mazeId"
	Register           = "/user"
	Login              = "/login"
	Maze               = "/maze"
	Solution           = "/maze/" + MazeId + "/solution"
	SolutionQueryParam = "steps"
)

func SetUserRoutes(grp GroupWrapper, c controllers2.UserController) {
	grp.POST(Register, c.Register)
	grp.POST(Login, c.Login)
}

func SetMazeRoutes(grp GroupWrapper, c controllers2.MazeController, a authentication.Authenticator) {
	grp.POST(Maze, c.CreateMaze, a.AuthenticateUser)
	grp.GET(Maze, c.GetMazes, a.AuthenticateUser)
	grp.GET(Solution, c.GetSolution, a.AuthenticateUser)
}
