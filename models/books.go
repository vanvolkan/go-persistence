package models

import "fmt"

type Book struct {
	isbn   string
	ttitle string
	author string
}

// AllBooks implementation
func (db *DB) AllBooks() ([]*Book, error) {
	rows, err := db.Query("Select * From books")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.isbn, &bk.ttitle, &bk.author)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return bks, nil
}
