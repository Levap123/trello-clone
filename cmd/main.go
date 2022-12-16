package main

import (
	"log"

	"github.com/Levap123/trello-clone/internal/repository"
)

func main() {
	db, err := repository.InitDb()
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatalln(err.Error())
	}
}
