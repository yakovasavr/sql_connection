package store

import (
	"context"
	"database/sql"
	"log"

	"github.com/yakovasavr/sql_connection/models"
)

type Repo struct {
	Conn *sql.DB
}

type BookRepository interface {
	GetLastBook(ctx context.Context) (models.Book, error)
}

func (repo Repo) GetLastBook(ctx context.Context) (models.Book, error) {
	var book models.Book

	rows, err := repo.Conn.Query("select title, author from myschema.books")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.Title, &book.Author)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return book, nil
}