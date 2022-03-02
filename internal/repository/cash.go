package repository

import (
	"database/sql"
	"fmt"

	"github.com/gh0st3e/BillyBetProject/internal/entity"
	"github.com/gh0st3e/BillyBetProject/internal/util"
	"github.com/pkg/errors"
)

var queryCash = map[string]string{
	util.AddCash: "INSERT INTO cash (`usd`,`rub`) VALUES ('%d','%d')",
	util.GetCash: "SELECT * FROM cash WHERE `cashid` = '%v'",
	util.SubCash: "UPDATE cash SET usd=usd-'%v',rub=rub-'%v' WHERE `cashid`='%v'",
}

var _ Cash = (*cash)(nil) //Magic

type cash struct {
	db *sql.DB
}

func NewCash(db *sql.DB) Cash {
	return cash{db: db}
}

func (c cash) Add(cash entity.Cash) error {
	rows, err := c.db.Query(fmt.Sprintf(queryCash[util.AddCash], cash.USD, cash.RUB))
	if err != nil {
		return errors.Wrap(err, "repository.Add.Query couldn't add new cash")
	}

	return rows.Err()
}

func (c cash) Get(id int) (entity.Cash, error) {
	rows, err := c.db.Query(fmt.Sprintf(queryCash[util.GetCash], id))
	if err != nil {
		return entity.Cash{}, errors.Wrap(err, "repository.Get.Query couldn't find user")
	}

	var balance entity.Cash
	for rows.Next() {
		err := rows.Scan(&balance.ID, &balance.USD, &balance.RUB)
		if err != nil {
			return entity.Cash{}, errors.Wrap(err, "repository.Get.Scan couldn't scan user")
		}
	}

	return balance, nil
}

func (c cash) Sub(id, rub int) (entity.Cash, error) {
	rows, err := c.db.Query(fmt.Sprintf(queryCash[util.SubCash], rub/100, rub, id))
	if err != nil {
		return entity.Cash{}, errors.Wrap(err, "repository.Add.Query couldn't add new cash")
	}

	rows, err = c.db.Query(fmt.Sprintf(queryCash[util.GetCash], id))
	var balance entity.Cash
	for rows.Next() {
		err := rows.Scan(&balance.ID, &balance.USD, &balance.RUB)
		if err != nil {
			return entity.Cash{}, errors.Wrap(err, "repository.Get.Scan couldn't scan user")
		}
	}

	return balance, nil
}
