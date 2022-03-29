package controllers

import (
	"errors"
	"github.com/fitims/exercise/cmd/restapi/controllers/requests"
	"github.com/fitims/exercise/cmd/restapi/controllers/responses"
	"github.com/fitims/exercise/cmd/restapi/middleware/authentication"
	"github.com/fitims/exercise/log"
	"github.com/fitims/exercise/maze"
	"github.com/fitims/exercise/repo"
	"github.com/labstack/echo/v4"
	"net/http"
)

// MazeController is the interface that defines the behaviour of the
// maze controller. MazeController handles all the requests regarding the maze
// ie. maze creation, retrieving mazes for user, getting maze solution
type MazeController interface {
	CreateMaze(c echo.Context) error
	GetMazes(c echo.Context) error
	GetSolution(c echo.Context) error
}

// defaultMazeController is the default implementation of maze controller
type defaultMazeController struct {
	repository repo.MazeRepository
}

// NewMazeController creates new maze controller that handles maze routes
func NewMazeController(r repo.MazeRepository) MazeController {
	return &defaultMazeController{
		repository: r,
	}
}

// CreateMaze handles maze creation
func (ctrl defaultMazeController) CreateMaze(c echo.Context) error {
	var req requests.CreateMazeRequest
	err := c.Bind(&req)
	if err != nil {
		log.Errorln("Could not bind CreateMazeRequest. Error: ", err)
		return c.JSON(http.StatusBadRequest, responses.MazeCreationFailed("The request is not valid"))
	}

	mazeId := ctrl.repository.GetMazeId()
	m, err := maze.NewMaze(mazeId, req.Entrance, req.GridSize, req.Walls)
	if err != nil {
		if errors.Is(err, maze.InvalidGridSizeErr) {
			return c.JSON(http.StatusBadRequest, responses.MazeCreationFailed(err.Error()))
		}

		if errors.Is(err, maze.InvalidCellErr) {
			return c.JSON(http.StatusBadRequest, responses.MazeCreationFailed(err.Error()))
		}

		if errors.Is(err, maze.InvalidRowErr) {
			return c.JSON(http.StatusBadRequest, responses.MazeCreationFailed(err.Error()))
		}

		if errors.Is(err, maze.InvalidColumnNameErr) {
			return c.JSON(http.StatusBadRequest, responses.MazeCreationFailed(err.Error()))
		}

		log.Errorln("Could not build the maze. Error: ", err)
		return c.JSON(http.StatusBadRequest, responses.MazeCreationFailed("Invalid request"))
	}

	authUsr := c.Get(authentication.AuthUser).(authentication.Identity)

	err = ctrl.repository.SaveMaze(authUsr.GetUsername(), m)
	if err != nil {
		log.Errorln("Could not get the details for the user. Error: ", err)
		return c.JSON(http.StatusInternalServerError, responses.MazeCreationFailed("Could not get user details. please try later on."))
	}
	return c.JSON(http.StatusOK, responses.MazeCreationSuccessful())
}

func (ctrl defaultMazeController) GetMazes(c echo.Context) error {
	authUsr := c.Get(authentication.AuthUser).(authentication.Identity)
	mazes, err := ctrl.repository.GetMazesForUser(authUsr.GetUsername())
	if err != nil {
		log.Errorln("Could not get mazes for the user. Error: ", err)
		return c.JSON(http.StatusInternalServerError, responses.MazeRetrievalFailed("Could not get mazes for user. please try later on."))
	}

	return c.JSON(http.StatusOK, responses.MazeRetrievalSuccessful(mazes))
}

func (ctrl defaultMazeController) GetSolution(c echo.Context) error {
	return nil
}
