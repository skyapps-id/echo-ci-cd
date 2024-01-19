package usecase

import (
	"context"
	"echo-ci-cd/repository"
)

type BookUsecase interface {
	GetBookByID(ctx context.Context, code string) (resp GetBookResponse, err error)
	GetBooks(ctx context.Context) (resp []GetBookResponse, err error)
}

type usecase struct {
	bookRepository repository.Book
}

func NewUsecase(bookRepository repository.Book) *usecase {
	return &usecase{
		bookRepository: bookRepository,
	}
}

func (uc *usecase) GetBookByID(ctx context.Context, code string) (resp GetBookResponse, err error) {
	book, err := uc.bookRepository.GetBookByID(ctx, code)
	if err != nil {
		return
	}

	resp = GetBookResponse{
		UUID:        book.UUID,
		Name:        book.Name,
		Description: book.Description,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}

	return
}

func (uc *usecase) GetBooks(ctx context.Context) (resp []GetBookResponse, err error) {
	books, err := uc.bookRepository.GetBooks(ctx)
	if err != nil {
		return
	}

	for _, row := range books {
		resp = append(resp, GetBookResponse{
			UUID:        row.UUID,
			Name:        row.Name,
			Description: row.Description,
			CreatedAt:   row.CreatedAt,
			UpdatedAt:   row.UpdatedAt,
		})
	}

	return
}
