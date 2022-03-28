package repo

import (
	"errors"
	"github.com/fitims/exercise/maze"
)

type MazeRepository interface {
	SaveMaze(username string, m *maze.Maze) error
	GetMazesForUser(username string) ([]maze.Maze, error)
	GetMaze(username string, id uint64) (*maze.Maze, error)
	GetMazeId() uint64
}

var (
	MazeDoesNotExistErr = errors.New("maze does not exist")
)
