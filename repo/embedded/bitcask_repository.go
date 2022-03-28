package embedded

import (
	"encoding/json"
	"fmt"
	"git.mills.io/prologic/bitcask"
	"github.com/dgrijalva/jwt-go"
	"github.com/fitims/exercise/cmd/restapi/middleware/authentication"
	"github.com/fitims/exercise/log"
	"github.com/fitims/exercise/maze"
	"github.com/fitims/exercise/repo"
	"math/big"
	"sync"
)

const (
	userPrefix = "user_"
	userKey    = userPrefix + "%s"

	mazePrefix = "maze_%s_"
	mazeKey    = mazePrefix + "%d"
	idKey      = "id"
)

type bitcaskRepository struct {
	db *bitcask.Bitcask
}

var (
	repoGate = sync.Mutex{}
)

// NewBitCaskRepository creates a new embedded database based on Bitcask (https://pkg.go.dev/github.com/prologic/bitcask)
// Bitcask is an embedded key/value store
func NewBitCaskRepository(path string) (*bitcaskRepository, error) {
	dBase, err := bitcask.Open(path)
	if err != nil {
		log.Errorln("Could not open database", err)
		return nil, err
	}

	return &bitcaskRepository{
		db: dBase,
	}, nil
}

// GetUser return details user for the provided username
func (r bitcaskRepository) GetUser(username string) (*repo.User, error) {
	key := []byte(fmt.Sprintf(userKey, username))
	if r.db.Has(key) {
		data, err := r.db.Get(key)
		if err != nil {
			return nil, err
		}

		var user repo.User
		err = json.Unmarshal(data, &user)
		if err != nil {
			log.Errorln("Cannot unmarshall user. Error: ", err)
			return nil, err
		}
		return &user, nil
	}
	return nil, repo.UserDoesNotExistErr
}

// SaveUser saves the user to bitcask
func (r bitcaskRepository) SaveUser(username, password string) error {
	key := []byte(fmt.Sprintf(userKey, username))
	if r.db.Has(key) {
		return repo.UserExistsErr
	}

	user := repo.User{
		Username: username,
		Password: password,
	}
	usrJson, err := json.Marshal(user)
	if err != nil {
		log.Errorln("Cannot marshall user. Error: ", err)
		return err
	}

	err = r.db.Put(key, usrJson)
	if err != nil {
		return err
	}
	return nil
}

// GetUserValidator returns an authentication.ClaimsValidator that is used by middleware
// to validate the user that is provided in auth token as claims
func (r bitcaskRepository) GetUserValidator() authentication.ClaimsValidator {
	return func(claims jwt.MapClaims) (authentication.Identity, error) {
		usr, err := repo.NewUser(claims)
		if err != nil {
			return nil, err
		}

		key := []byte(fmt.Sprintf(userKey, usr.Username))
		if r.db.Has(key) {
			return usr, nil
		}
		return nil, repo.UserDoesNotExistErr
	}
}

// Maze repository implementation

// GetMazeId returns next free maze ID
func (r bitcaskRepository) GetMazeId() uint64 {
	repoGate.Lock()
	defer repoGate.Unlock()

	no := big.Int{}
	key := []byte(idKey)
	if !r.db.Has(key) {
		no.SetUint64(1)
		noJson, _ := no.MarshalJSON()
		r.db.Put(key, noJson)
		return 1
	}

	buf, _ := r.db.Get(key)
	_ = no.UnmarshalJSON(buf)
	return no.Uint64()
}

// SaveMaze saves the maze to bitcask database. If the data cannot be read then an error
// is returned
func (r bitcaskRepository) SaveMaze(username string, m *maze.Maze) error {
	key := []byte(fmt.Sprintf(mazeKey, username, m.Id))

	mazeJson, err := json.Marshal(m)
	if err != nil {
		log.Errorln("Cannot marshall maze. Error: ", err)
		return err
	}

	err = r.db.Put(key, mazeJson)
	if err != nil {
		return err
	}
	return nil
}

// GetMazesForUser returns all the mazes for the user provided
func (r bitcaskRepository) GetMazesForUser(username string) ([]maze.Maze, error) {
	mazes := make([]maze.Maze, 0)
	prefix := []byte(fmt.Sprintf(mazePrefix, username))
	_ = r.db.Scan(prefix, func(key []byte) error {
		val, err := r.db.Get(key)
		if err == nil {
			var m maze.Maze
			err = json.Unmarshal(val, &m)
			if err == nil {
				mazes = append(mazes, m)
			}
		}
		return nil
	})
	return mazes, nil
}

// GetMaze return a specific maze. If the maze does not exist  repo.MazeDoesNotExistErr
// is returned
func (r bitcaskRepository) GetMaze(username string, id uint64) (*maze.Maze, error) {
	key := []byte(fmt.Sprintf(mazeKey, username, id))
	if r.db.Has(key) {
		data, err := r.db.Get(key)
		if err != nil {
			return nil, err
		}

		var maze maze.Maze
		err = json.Unmarshal(data, &maze)
		if err != nil {
			log.Errorln("Cannot unmarshall user. Error: ", err)
			return nil, err
		}
		return &maze, nil
	}
	return nil, repo.MazeDoesNotExistErr
}
