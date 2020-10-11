package main

import (
	"bookstore/models"
	"log"
	"net/http"
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

	http.HandleFunc("/books", env.booksIndex)
	http.ListenAndServe(":3000", nil)
}
