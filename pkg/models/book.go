package models

import (
	"fmt"

	"github.com/Thebassplayer/go_book_management_system/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("id=?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("id=?", id).Delete(book)
	return book
}

func UpdateBook(id int64, updatedBook *Book) (*Book, *gorm.DB) {
	var book Book
	db = db.Where("id = ?", id).Find(&book)

	if db.Error != nil {
		// Handle error, book not found
		return nil, db
	}

	db = db.Model(&book).Updates(Book{
		Name:        updatedBook.Name,
		Author:      updatedBook.Author,
		Publication: updatedBook.Publication,
	})

	fmt.Println("book: ", book)

	return &book, db
}
