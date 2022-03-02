package repository

import (
	"database/sql"
	"fmt"

	"github.com/gh0st3e/BillyBetProject/internal/config"
	"github.com/gh0st3e/BillyBetProject/internal/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func Connect(cfg config.Database) (*sql.DB, error) {
	db, err := sql.Open(util.DriverName, fmt.Sprintf("%s%s", util.ConnectionPrefix, cfg.Address))
	if err != nil {
		return nil, errors.Wrap(err, "repository.Connect.Open couldn't connect to sql")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Connect.Ping couldn't ping database")
	}

	return db, nil
}
