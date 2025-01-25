package models

import (
	"github.com/itsRaCl/11_go_projects/bookstore_mgmt/pkg/config"
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

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.First(&getBook, Id)
	return &getBook, db
}

func DeleteBook(Id int64) *Book {
	var delBook Book
	db.Delete(&delBook, Id)

	return &delBook
}
