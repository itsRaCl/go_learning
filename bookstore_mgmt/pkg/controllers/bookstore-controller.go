package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/itsRaCl/11_go_projects/bookstore_mgmt/pkg/models"
	"github.com/itsRaCl/11_go_projects/bookstore_mgmt/pkg/utils"
)

func GetBooks(res http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()

	js, _ := json.Marshal(newBooks)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(js)
}

func GetBookById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while Parsing")
	}

	bookDetails, _ := models.GetBookById(ID)

	js, _ := json.Marshal(bookDetails)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(js)
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	createdBook := &models.Book{}
	utils.ParseBody(req, createdBook)

	b := createdBook.CreateBook()

	js, _ := json.Marshal(b)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(js)
}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteBook(ID)
	js, _ := json.Marshal(book)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(js)
}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
	var updatedBook = &models.Book{}

	utils.ParseBody(req, updatedBook)

	vars := mux.Vars(req)

	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error in parsing")
	}

	oldBook, db := models.GetBookById(ID)

	if updatedBook.Name != "" {
		oldBook.Name = updatedBook.Name
	}

	if updatedBook.Author != "" {
		oldBook.Author = updatedBook.Author
	}

	if updatedBook.Publication != "" {
		oldBook.Publication = updatedBook.Publication
	}

	db.Save(&oldBook)

	js, _ := json.Marshal(oldBook)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(js)
}
