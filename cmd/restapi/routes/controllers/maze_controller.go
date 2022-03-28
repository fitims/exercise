package controllers

import (
	"github.com/fitims/exercise/repo"
	"github.com/labstack/echo/v4"
)

type MazeController interface {
	CreateMaze(c echo.Context) error
	GetMazes(c echo.Context) error
	GetSolution(c echo.Context) error
}

type defaultMazeController struct {
	repository repo.MazeRepository
}

func NewMazeController(r repo.MazeRepository) MazeController {
	return &defaultMazeController{
		repository: r,
	}
}

func (ctrl defaultMazeController) CreateMaze(c echo.Context) error {

	return nil
}

func (ctrl defaultMazeController) GetMazes(c echo.Context) error {
	return nil
}

func (ctrl defaultMazeController) GetSolution(c echo.Context) error {
	return nil
}
