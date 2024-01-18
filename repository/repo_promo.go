package repository

import (
	"context"
	"echo-ci-cd/entity"

	"gorm.io/gorm"
)

type Book interface {
	GetBookByID(ctx context.Context, code string) (result entity.Book, err error)
	GetBooks(ctx context.Context) (results []entity.Book, err error)
}

type book struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) Book {
	if db == nil {
		panic("db is nil")
	}

	return &book{
		db: db,
	}
}

func (r *book) GetBookByID(ctx context.Context, code string) (result entity.Book, err error) {
	err = r.db.WithContext(ctx).Table("books").Where("books.uuid = ?", code).Take(&result).Error
	if err != nil {
		return
	}

	return
}

func (r *book) GetBooks(ctx context.Context) (results []entity.Book, err error) {
	err = r.db.WithContext(ctx).Table("books").Find(&results).Error
	if err != nil {
		return
	}

	return
}
