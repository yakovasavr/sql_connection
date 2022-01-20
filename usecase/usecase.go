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
	GetLastBook(ctx context.Context) (models.Book, error)
}

func (s BookUsecase) GetLastBook(ctx context.Context) (models.Book, error) {

	book, err := s.BookRepo.GetLastBook(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return book, nil
}
