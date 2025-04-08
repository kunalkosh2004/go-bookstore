package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kunalkosh2004/go-bookstore/pkg/models"
	"github.com/kunalkosh2004/go-bookstore/pkg/utils"
)


var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res ,_ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("Error while parsing")
	}
	bookDetails,_ := models.GetBookById(Id)
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := models.Book{}
	utils.ParseBody(r,CreateBook)
	b := CreateBook.CreateBook()
	res,_ :=json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r* http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("Error While parsing")
	}
	book := models.DeleteBook(ID)
	res,_ := json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updatedBook = &models.Book{}
	utils.ParseBody(r, updatedBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("Error While Parsing")
	}
	bookDetails, db := models.GetBookById(Id)
	if updatedBook.Name != ""{
		bookDetails.Name = updatedBook.Name
	}
	if updatedBook.Author !=""{
		bookDetails.Author = updatedBook.Author
	}
	if updatedBook.Publication != ""{
		bookDetails.Publication = updatedBook.Publication
	}
	db.Save(&bookDetails)
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}