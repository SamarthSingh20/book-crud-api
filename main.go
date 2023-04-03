package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Price  string  `json:"price"`
	Author *author `json:"author"`
}
type author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func createbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func getbooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)
	for _, item := range books {
		if item.ID == parms["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func updatebook(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    parms := mux.Vars(r)
	for index, item := range books {
		if item.ID == parms["id"] {
			books = append(books[:index], books[index+1:]...)
            var book Book
            json.NewDecoder(r.Body).Decode(&book)
            books = append(books[:index], book)
            json.NewEncoder(w).Encode(book)
		}
	}
}

func deletebook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)
	for index, item := range books {
		if item.ID == parms["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
}

func main() {
	fmt.Println("Books crud API")
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Title: "ABC", Price: "100", Author: &author{Firstname: "Ram", Lastname: "Singh"}})
	books = append(books, Book{ID: "2", Title: "Ramsetu", Price: "500", Author: &author{Firstname: "Rama", Lastname: "Lingam"}})

	r.HandleFunc("/books", createbook).Methods("POST")
	r.HandleFunc("/books", getbooks).Methods("GET")
	r.HandleFunc("/books/{id}", getbook).Methods("GET")
	r.HandleFunc("/books/{id}", deletebook).Methods("DELETE")
	r.HandleFunc("/update/{id}", updatebook).Methods("PUT")

	http.ListenAndServe(":8000", r)

}
