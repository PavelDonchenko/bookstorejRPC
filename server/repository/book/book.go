package repository

import (
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	model "github.com/PavelDonchenko/bookstorejRPC/server/models"
	"github.com/jinzhu/gorm"
)

type BookRepo struct {
	db   *gorm.DB
	book *model.Book
}

func NewBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (b *BookRepo) Validate() error {

	if b.book.Name == "" {
		return errors.New("Required Name")
	}
	if b.book.BookAuthor == "" {
		return errors.New("Required Author")
	}
	return nil
}

func (b *BookRepo) Prepare() {
	b.book.ID = 0
	b.book.Name = html.EscapeString(strings.TrimSpace(b.book.Name))
	b.book.BookAuthor = html.EscapeString(strings.TrimSpace(b.book.BookAuthor))
	b.book.CreatedAt = time.Now()
	b.book.UpdatedAt = time.Now()
}

func (u *BookRepo) GetAllBooks(offset int, limit int) ([]model.Book, error) {
	books := []model.Book{}
	err := u.db.Debug().Model(&model.Book{}).Limit(limit).Offset(offset).Find(&books).Error
	if err != nil {
		return []model.Book{}, err
	}

	return books, nil
}

func (b *BookRepo) GetBook(id uint32) (*model.Book, error) {
	result := &model.Book{}
	err := b.db.Debug().Model(&model.Book{}).Where("id = ?", id).Take(result).Error
	if err != nil {
		return &model.Book{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &model.Book{}, errors.New("Book Not Found")
	}

	return result, nil
}

func (b *BookRepo) CreateBook(book model.Book) (model.Book, error) {
	err := b.db.Debug().Create(&book).Error
	if err != nil {
		return model.Book{}, err
	}

	return book, err
}

func (b *BookRepo) UpdateBook(book model.Book) (model.Book, error) {
	if err := b.db.First(&book).Error; err != nil {
		return book, err
	}

	err := b.db.Debug().Model(&book).Updates(&book).Error
	fmt.Println(err)
	return book, err
}

func (b *BookRepo) DeleteBook(id uint32) (bool, error) {
	db := b.db.Debug().Model(&model.Book{}).Where("id = ?", id).Take(&model.Book{}).Delete(&model.Book{})
	result := true
	if db.Error != nil {
		result = false
		return result, db.Error
	}
	return result, nil
}
