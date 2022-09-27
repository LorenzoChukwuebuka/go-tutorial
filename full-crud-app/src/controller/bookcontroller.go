package controller

import (
	"encoding/json"
	"fmt"
	"full-crud-tutorial/src/models"
	"full-crud-tutorial/src/util"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createbook := &models.Book{}
	util.ParseBody(r, createbook)
	b := createbook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println(("error while parsing"))
	}

	bookDetails, _ := models.GetBookById(Id)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBooks(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}

	util.ParseBody(r, updateBook)

	vars := mux.Vars(r)

	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	//find book

	bookDetails, db := models.GetBookById(Id)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBooks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteBook(Id)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
