package controllers

import (
	"errors"
	"github.com/fitims/exercise/cmd/restapi/middleware/authentication"
	"github.com/fitims/exercise/cmd/restapi/params"
	"github.com/fitims/exercise/maze"
	"github.com/fitims/exercise/repo"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GetMazeFunc func() (*maze.Maze, error)

type MockRepository struct {
	fn GetMazeFunc
}

func (r MockRepository) SaveMaze(username string, m *maze.Maze) error {
	return nil
}

func (r MockRepository) GetMazesForUser(username string) ([]maze.Maze, error) {
	return nil, nil
}

func (r MockRepository) GetMaze(username string, id uint64) (*maze.Maze, error) {
	return r.fn()
}

func (MockRepository) GetMazeId() uint64 {
	return 1
}

func TestDefaultMazeController_GetSolution_InvalidSteps(t *testing.T) {
	// setup
	mockDb := MockRepository{
		fn: func() (*maze.Maze, error) {
			return &maze.Maze{}, nil
		},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/maze/1/solution?steps=mid", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controller := NewMazeController(mockDb)

	// act
	controller.GetSolution(c)

	// assert
	if rec.Code != http.StatusBadRequest {
		t.Error("Invalid response code. Expected: BadRequest (400), Actual: ", rec.Code)
	}

	resp := rec.Body.String()
	if resp != `{"is_success":false,"message":"Invalid request. Steps can be 'min' or 'max"}`+"\n" {
		t.Error("Invalid response body")
	}
}

func TestDefaultMazeController_GetSolution_InvalidMazeId(t *testing.T) {
	// setup
	mockDb := MockRepository{
		fn: func() (*maze.Maze, error) {
			return &maze.Maze{}, nil
		},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/maze/invalid/solution?steps=min", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames(params.MazeId)
	c.SetParamValues("invalid")
	controller := NewMazeController(mockDb)

	// act
	controller.GetSolution(c)

	// assert
	if rec.Code != http.StatusBadRequest {
		t.Error("Invalid response code. Expected: BadRequest (400), Actual: ", rec.Code)
	}

	resp := rec.Body.String()
	if resp != `{"is_success":false,"message":"Invalid request. Maze Id has to be valid number"}`+"\n" {
		t.Error("Invalid response body")
	}
}

func TestDefaultMazeController_GetSolution_MazeDoesNotExist(t *testing.T) {
	// setup
	mockDb := MockRepository{
		fn: func() (*maze.Maze, error) {
			return nil, repo.MazeDoesNotExistErr
		},
	}

	authUser := repo.User{
		Username: "user",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/maze/1/solution?steps=min", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames(params.MazeId)
	c.SetParamValues("1")
	c.Set(authentication.AuthUser, &authUser)
	controller := NewMazeController(mockDb)

	// act
	controller.GetSolution(c)

	// assert
	if rec.Code != http.StatusBadRequest {
		t.Error("Invalid response code. Expected: BadRequest (400), Actual: ", rec.Code)
	}

	resp := rec.Body.String()
	if resp != `{"is_success":false,"message":"Maze does not exist"}`+"\n" {
		t.Error("Invalid response body")
	}
}

func TestDefaultMazeController_GetSolution_InternalRepoError(t *testing.T) {
	// setup
	mockDb := MockRepository{
		fn: func() (*maze.Maze, error) {
			return nil, errors.New("Internal Repo Error")
		},
	}

	authUser := repo.User{
		Username: "user",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/maze/1/solution?steps=min", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames(params.MazeId)
	c.SetParamValues("1")
	c.Set(authentication.AuthUser, &authUser)
	controller := NewMazeController(mockDb)

	// act
	controller.GetSolution(c)

	// assert
	if rec.Code != http.StatusInternalServerError {
		t.Error("Invalid response code. Expected: InternalServerError (500), Actual: ", rec.Code)
	}

	resp := rec.Body.String()
	if resp != `{"is_success":false,"message":"Could not get maze details. please try later on."}`+"\n" {
		t.Error("Invalid response body")
	}
}

func TestDefaultMazeController_GetSolution_NoSolutions(t *testing.T) {
	// setup
	mockDb := MockRepository{
		fn: func() (*maze.Maze, error) {
			return &maze.Maze{
				Id:        1,
				State:     maze.NoSolutions,
				Solutions: nil,
			}, nil
		},
	}

	authUser := repo.User{
		Username: "user",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/maze/1/solution?steps=min", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames(params.MazeId)
	c.SetParamValues("1")
	c.Set(authentication.AuthUser, &authUser)
	controller := NewMazeController(mockDb)

	// act
	controller.GetSolution(c)

	// assert
	if rec.Code != http.StatusOK {
		t.Error("Invalid response code. Expected: InternalServerError (500), Actual: ", rec.Code)
	}

	resp := rec.Body.String()
	if resp != `{"is_success":false,"message":"maze does not have a solution"}`+"\n" {
		t.Error("Invalid response body")
	}
}

func TestDefaultMazeController_GetSolution_TooManySolutions(t *testing.T) {
	// setup
	mockDb := MockRepository{
		fn: func() (*maze.Maze, error) {
			return &maze.Maze{
				Id:        1,
				State:     maze.TooManySolutions,
				Solutions: nil,
			}, nil
		},
	}

	authUser := repo.User{
		Username: "user",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/maze/1/solution?steps=min", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames(params.MazeId)
	c.SetParamValues("1")
	c.Set(authentication.AuthUser, &authUser)
	controller := NewMazeController(mockDb)

	// act
	controller.GetSolution(c)

	// assert
	if rec.Code != http.StatusOK {
		t.Error("Invalid response code. Expected: InternalServerError (500), Actual: ", rec.Code)
	}

	resp := rec.Body.String()
	if resp != `{"is_success":false,"message":"maze has more than one solution"}`+"\n" {
		t.Error("Invalid response body")
	}
}

func TestDefaultMazeController_GetSolution_GetShortestPath(t *testing.T) {
	// setup
	mockDb := MockRepository{
		fn: func() (*maze.Maze, error) {
			return &maze.Maze{
				Id:    1,
				State: maze.Solved,
				Solutions: []maze.Path{
					{{Row: 0, Col: 1}, {Row: 1, Col: 1}, {Row: 2, Col: 1}, {Row: 3, Col: 1}},
					{{Row: 0, Col: 1}, {Row: 1, Col: 1}, {Row: 2, Col: 1}, {Row: 3, Col: 1}, {Row: 4, Col: 1}, {Row: 5, Col: 1}},
				},
			}, nil
		},
	}

	authUser := repo.User{
		Username: "user",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/maze/1/solution?steps=min", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames(params.MazeId)
	c.SetParamValues("1")
	c.Set(authentication.AuthUser, &authUser)
	controller := NewMazeController(mockDb)

	// act
	controller.GetSolution(c)

	// assert
	if rec.Code != http.StatusOK {
		t.Error("Invalid response code. Expected: Ok (200), Actual: ", rec.Code)
	}

	resp := rec.Body.String()
	if resp != `{"path":["B1","B2","B3","B4"]}`+"\n" {
		t.Error("Invalid response body")
	}
}

func TestDefaultMazeController_GetSolution_GetLongestPath(t *testing.T) {
	// setup
	mockDb := MockRepository{
		fn: func() (*maze.Maze, error) {
			return &maze.Maze{
				Id:    1,
				State: maze.Solved,
				Solutions: []maze.Path{
					{{Row: 0, Col: 1}, {Row: 1, Col: 1}, {Row: 2, Col: 1}, {Row: 3, Col: 1}},
					{{Row: 0, Col: 1}, {Row: 1, Col: 1}, {Row: 2, Col: 1}, {Row: 3, Col: 1}, {Row: 4, Col: 1}, {Row: 5, Col: 1}},
				},
			}, nil
		},
	}

	authUser := repo.User{
		Username: "user",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/maze/1/solution?steps=max", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames(params.MazeId)
	c.SetParamValues("1")
	c.Set(authentication.AuthUser, &authUser)
	controller := NewMazeController(mockDb)

	// act
	controller.GetSolution(c)

	// assert
	if rec.Code != http.StatusOK {
		t.Error("Invalid response code. Expected: Ok (200), Actual: ", rec.Code)
	}

	resp := rec.Body.String()
	if resp != `{"path":["B1","B2","B3","B4","B5","B6"]}`+"\n" {
		t.Error("Invalid response body")
	}
}
