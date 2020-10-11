package main

import (
	"bookstore/models"
)

type Env struct {
	db models.Datastore
}

func main() {
	db, err := models.NewDB("postgres://postgress:Volkan21!@localhost/bookstore")
}
