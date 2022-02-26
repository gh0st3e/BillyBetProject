package main

import (
	"fmt"
	"github.com/gh0st3e/BillyBetProject/internal/config"
	"github.com/gh0st3e/BillyBetProject/internal/entity"
	"github.com/gh0st3e/BillyBetProject/internal/repository"
	"log"
)

func main() {

	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("cmd.Main.Init couldn't init config: %s", err)
	}
	fmt.Println(cfg)
	db, err := repository.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("cmd.Main.Connect couldn't connect database: %s", err)
	}

	repos := repository.NewRepositories(db)
	err = repos.User.Add(entity.User{
		Name:    "Denis",
		Surname: "Leonov",
		CashID:  2,
		Ban:     false,
	})
	if err != nil {
		log.Fatalf("cmd.Main.Add couldn't add new user, %s", err)
	}

	log.Println(cfg)
	log.Println(db)
	log.Println(repos)
	//db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:8597)/golang")
}
