package main

import (
	"bookstore/models"
	"log"
)

type Env struct {
	db models.Datastore
}

func main() {
	db, err := models.NewDB("postgres://postgress:Volkan21!@localhost/bookstore")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}
}
