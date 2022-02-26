package repository

import (
	"database/sql"
	"github.com/gh0st3e/BillyBetProject/internal/entity"
)

type User interface {
	Add(user entity.User) error
	Get(id int) (entity.User, error)
	Remove(id int) error
}
type Cash interface {
	Add(rub int) (entity.Cash, error)
	Get(id int) (entity.Cash, error)
	Sub(rub int) (entity.Cash, error)
}

type Repositories struct {
	User User
}

func NewRepositories(db *sql.DB) Repositories {
	return Repositories{User: NewUser(db)}
}
