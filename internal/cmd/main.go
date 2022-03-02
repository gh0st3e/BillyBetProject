package main

import (
	"fmt"
	"log"

	"github.com/gh0st3e/BillyBetProject/internal/config"
	"github.com/gh0st3e/BillyBetProject/internal/entity"
	"github.com/gh0st3e/BillyBetProject/internal/repository"
)

func AddUser(repos repository.Repositories) {

	err := repos.User.Add(entity.User{
		Name:    "Denis",
		Surname: "Leonov",
		CashID:  2,
		Ban:     false,
	})
	if err != nil {
		log.Fatalf("cmd.Main.AddUser couldn't add new user, %s", err)
	}
}

func AddCash(repos repository.Repositories) {

	err := repos.Cash.Add(entity.Cash{
		USD: 3,
		RUB: 300,
	})
	if err != nil {
		log.Fatalf("cmd.Main.AddCash couldn't add new cash, %s", err)
	}
}

func GetUser(repos repository.Repositories) {
	get, err := repos.User.Get(1)
	if err != nil {
		log.Fatalf("cmd.Main.GetUser couldn't get user, %s", err)
	}
	fmt.Println(get)
}
func GetCash(repos repository.Repositories) {
	get, err := repos.Cash.Get(2)
	if err != nil {
		log.Fatalf("cmd.Main.GetCash couldn't get cash, %s", err)
	}
	fmt.Println(get)
}
func SubCash(repos repository.Repositories) {
	get, err := repos.Cash.Sub(2, 300)
	if err != nil {
		log.Fatalf("cmd.Main.GetCash couldn't sub cash, %s", err)
	}
	fmt.Println(get)
}
func RemoveUser(repos repository.Repositories) {
	err := repos.User.Remove(2)
	if err != nil {
		log.Fatalf("cmd.Main.RemoveUser couldn't delete user, %s", err)
	}
}

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
	//TODO Change arguments!!!
	//AddCash(repos)
	//GetCash(repos)
	//SubCash(repos)
	//AddUser(repos)
	//GetUser(repos)
	//RemoveUser(repos)
	log.Println(cfg)
	log.Println(db)
	log.Println(repos)

}
