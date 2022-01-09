package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)
type Book struct {
	title	string
	author	string
}

type BookStore []Book

type Store struct {
	config 		*Config
	db     		*sql.DB
	BookList	BookStore
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) GetTable() error {
	newBookStore := make(BookStore, 0)

	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	rows, err := s.db.Query("select title, author from myschema.books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		newBook := new(Book)
		err := rows.Scan(&newBook.title, &newBook.author)
		if err != nil {
			log.Fatal(err)
		}
		newBookStore = append(newBookStore, *newBook)
		// log.Println(title, author)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	s.BookList = newBookStore

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}