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
)

const (
	userKey = "user_%s"
)

type bitcaskRepository struct {
	db *bitcask.Bitcask
}

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

func (r bitcaskRepository) RegisterUser(username, password string) error {
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

func (r bitcaskRepository) SaveMaze(username string, m *maze.Maze) error {

}

func (r bitcaskRepository) GetMazesForUser(username string) ([]maze.Maze, error) {

}

func (r bitcaskRepository) GetMaze(id int) (*maze.Maze, error) {

}
