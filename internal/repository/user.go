package repository

import (
	"database/sql"
	"fmt"
	"github.com/gh0st3e/BillyBetProject/internal/entity"
	"github.com/gh0st3e/BillyBetProject/internal/util"
	"github.com/pkg/errors"
)

var _ User = (*user)(nil) //Magic

var query = map[string]string{
	util.AddUser: "INSERT INTO user (`name`, `surname`, `cashid`, `ban`) VALUES ('%v','%v','%v','%v')",
}

//insert, err := db.Query(fmt.Sprintf("INSERT INTO `todo` (`id`,`todo`) VALUES ('%d','%s')", size+1, todo))

type user struct {
	db *sql.DB
}

func NewUser(db *sql.DB) User {
	return user{db: db}
}

func (u user) Add(user entity.User) error {
	rows, err := u.db.Query(fmt.Sprintf(query[util.AddUser], user.Name, user.Surname, user.CashID, util.FromBoolToInt(user.Ban)))
	if err != nil {
		return errors.Wrap(err, "repository.Add.Query couldn't add user")
	}

	return rows.Err()

}
func (u user) Get(id int) (entity.User, error) {
	return entity.User{}, nil
}
func (u user) Remove(id int) error {
	return nil
}
