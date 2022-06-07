package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	bookdatastore "janes_book_store/book_data_store"
	"log"
	"net/http"

	"janes_book_store/jwt_secure_access"

	"goji.io"
	"goji.io/pat"
)

var (
	VERSION      = "1.0"
	SECURED_URIS = []string{
		"/books/publish", "/books/unpublish", "/signin",
	}
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Janes Book Store api version: %s", VERSION)
}

func signinHandler(w http.ResponseWriter, r *http.Request) {

}

func getBooksFromAuthor(w http.ResponseWriter, r *http.Request) {

}
func getDetailsOfaTitle(w http.ResponseWriter, r *http.Request) {

}
func getBooksGeneralQuery(w http.ResponseWriter, r *http.Request) {

}
func getDetailsGeneralQuery(w http.ResponseWriter, r *http.Request) {

}
func publishBook(w http.ResponseWriter, r *http.Request) {
	var booksParam bookdatastore.Book
	booksParam = booksParam.fillRequest(r)
	log.Println("Handler: ", booksParam)
	//fmt.Fprintf(w, "%d %s %s", booksParam.BookID, booksParam.Author, booksParam.Title)
	bookdatastore.PublishBook(booksParam)
}
func unpublishBook(w http.ResponseWriter, r *http.Request) {

}
func authorNew(w http.ResponseWriter, r *http.Request) {

}

func initHandlers(mux *goji.Mux) {
	mux.HandleFunc(pat.Get("/status"), statusHandler)
	mux.HandleFunc(pat.Get("/signin"), signinHandler)

	//Query on books
	mux.HandleFunc(pat.Get("/books/author/:author"), getBooksFromAuthor)
	mux.HandleFunc(pat.Get("/books/title/:title"), getDetailsOfaTitle)
	mux.HandleFunc(pat.Get("/books"), getBooksGeneralQuery)
	mux.HandleFunc(pat.Get("/books/details"), getDetailsGeneralQuery)

	//publication routes
	mux.HandleFunc(pat.Post("/books/publish"), publishBook)
	mux.HandleFunc(pat.Post("/books/unpublish"), unpublishBook)

	//Nice to have end point to create an author
	mux.HandleFunc(pat.Post("/author/new"), authorNew)

	//Middleware to authenticate using jwt, log the API start and end
	//This framework can be enhanced with features.
	mux.Use(logAndAuthenticate)
	http.ListenAndServe("localhost:8000", mux)
}

//middleware take care of authentication of secured routes.
func logAndAuthenticate(inner http.Handler) http.Handler {
	mw := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Started: ", r.RequestURI, r.Method, r.Host)
		if contains(SECURED_URIS, r.RequestURI) {
			if jwt_secure_access.AuthorizeRequest(r) != true {
				w.Write([]byte("Failed: Authorization"))
				w.WriteHeader(http.StatusUnauthorized)
			}
		}
		inner.ServeHTTP(w, r)
	}
	return http.HandlerFunc(mw)
}

//move the below to a useful package fore refactoring
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}


func (ba bookdatastore.Book) fillRequest(r *http.Request) bookdatastore.Book {
	body, err := getRawBody(r)
	log.Println("Debug: ", string(body))
	if err == nil {
		err = json.Unmarshal(body, &ba)
	}
	if err != nil {
		log.Println("Error: ", err.Error())
		ba.Title, ba.Author, ba.BookID = "", "", 0
	}
	log.Println("Debug: ", ba)
	return ba
}

func getRawBody(r *http.Request) ([]byte, error) {
	if r.ContentLength == 0 {
		return []byte{}, fmt.Errorf("error cannot be empty")
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("cannot read request payload")
	}
	return body, nil
}
