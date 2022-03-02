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
	Add(cash entity.Cash) error
	Get(id int) (entity.Cash, error)
	Sub(id, rub int) (entity.Cash, error)
}

type Repositories struct {
	User User
	Cash Cash
}

func NewRepositories(db *sql.DB) Repositories {
	return Repositories{
		User: NewUser(db),
		Cash: NewCash(db),
	}
}
