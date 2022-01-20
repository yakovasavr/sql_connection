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
	GetBook(ctx context.Context) (models.Book, error)
}

func (repo Repo) GetBook(ctx context.Context) (models.Book, error) {
	var book models.Book

	rows, err := repo.Conn.Query("select title, author from myschema.books")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	err = rows.Scan(&book.Title, &book.Author)
	if err != nil {
		log.Fatal(err)
	}

	return book, nil
}