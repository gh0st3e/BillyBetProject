package repository

import (
	"database/sql"
	"fmt"

	"github.com/gh0st3e/BillyBetProject/internal/entity"
	"github.com/gh0st3e/BillyBetProject/internal/util"
	"github.com/pkg/errors"
)

var _ User = (*user)(nil) //Magic

var queryUser = map[string]string{
	util.AddUser: "INSERT INTO user (`name`, `surname`, `cashid`, `ban`) VALUES ('%v','%v','%v','%v')",
	util.GetUser: "SELECT * FROM user WHERE `id` = '%v'",
	util.DelUser: "DELETE FROM user WHERE `id`='%v'",
}

type user struct {
	db *sql.DB
}

func NewUser(db *sql.DB) User {
	return user{db: db}
}

func (u user) Add(user entity.User) error {
	rows, err := u.db.Query(fmt.Sprintf(queryUser[util.AddUser], user.Name, user.Surname, user.CashID, util.FromBoolToInt(user.Ban)))
	if err != nil {
		return errors.Wrap(err, "repository.Add.Query couldn't add user")
	}

	return rows.Err()

}

func (u user) Get(id int) (entity.User, error) {
	var man entity.User

	rows, err := u.db.Query(fmt.Sprintf(queryUser[util.GetUser], id))
	if err != nil {
		return entity.User{}, errors.Wrap(err, "repository.Get.Query couldn't find user")
	}

	for rows.Next() {
		err := rows.Scan(&man.ID, &man.Name, &man.Surname, &man.CashID, &man.Ban)
		if err != nil {
			return entity.User{}, errors.Wrap(err, "repository.Get.Scan couldn't scan user")
		}
	}

	return man, nil
}

func (u user) Remove(id int) error {
	_, err := u.db.Query(fmt.Sprintf(queryUser[util.DelUser], id))
	if err != nil {
		return errors.Wrap(err, "repository.Remove.Query couldn't delete user")
	}

	return nil
}
