package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"math/rand"
// 	"strconv"
// 	"github.com/gorilla/mux"
// )

//Book Stuct (Model)
type Book struct {
	ID 		 string `json:"id"`
	Isbn     string `json:"isbn"`
	Title	 string `json:"title"`
	Author	 *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname		string		`json:"firstname"`
	Lastname		string		`json:"lasttname"`
}

// Get All Books
func getBooks(w  http.ResponseWriter, r *http.Request) {

}
// Get All Book
func getBook(w  http.ResponseWriter, r *http.Request) {

}
// Create a New Books
func createBook(w  http.ResponseWriter, r *http.Request) {

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

	//Route Handlers /Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books", updateBook).Methods("PUT")
	r.HandleFunc("/api/books", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}