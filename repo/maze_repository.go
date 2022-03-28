package repo

import "github.com/fitims/exercise/maze"

type MazeRepository interface {
	SaveMaze(username string, m *maze.Maze) error
	GetMazesForUser(username string) ([]maze.Maze, error)
	GetMaze(id int) (*maze.Maze, error)
}
