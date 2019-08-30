package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

//Book Stuct (Model)
type Book struct {
	ID 		 string  `json:"id"`
	Isbn     string  `json:"isbn"`
	Title	 string  `json:"title"`
	Author	 *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname		string		`json:"firstname"`
	Lastname		string		`json:"lasttname"`
}

//Init books var as a sclice book struct
var books []Book

// Get All Books
func getBooks(w  http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
// Get Single Book
func getBook(w  http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r) //Get Params
	//Loop through books and find with id
	for _, item:= range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
// Create a New Books
func createBook(w  http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) // just Mock Id
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
// Update  Books
func updateBook(w  http.ResponseWriter, r *http.Request) {

}
// Get All Books
func deleteBook(w  http.ResponseWriter, r *http.Request) {

}



func main()  {
	//Init Router
	r:= mux.NewRouter()

	// Mock Data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "324245", Title: "Book One", Author: &Author{Firstname: "Elteyab", Lastname: "Ali"}})
books = append(books, Book{ID: "2", Isbn: "80823", Title: "Book Two", Author: &Author{Firstname: "Hassan", Lastname: "Omer "}})
books = append(books, Book{ID: "3", Isbn: "86788", Title: "Book Three", Author: &Author{Firstname: "Omer", Lastname: "Khalid "}})
books = append(books, Book{ID: "3", Isbn: "1234567", Title: "Book Fore", Author: &Author{Firstname: "Mudether", Lastname: "Ghade "}})

	//Route Handlers /Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books", updateBook).Methods("PUT")
	r.HandleFunc("/api/books", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}