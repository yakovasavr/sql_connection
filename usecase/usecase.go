package usecase

import (
	"context"
	"log"

	"github.com/yakovasavr/sql_connection/models"
	"github.com/yakovasavr/sql_connection/store"
)

type BookUsecase struct {
	BookRepo store.BookRepository
}

type BookService interface {
	GetBook(ctx context.Context) (models.Book, error)
}

func (s BookUsecase) GetBook(ctx context.Context) (models.Book, error) {

	book, err := s.BookRepo.GetBook(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return book, nil
}
