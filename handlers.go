package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	var publishBooksParam bookdatastore.BooksWithAuthor
	publishBooksParam = publishBooksParam.FillRequest(r)
	log.Println("Handler: ", publishBooksParam)

	book, err := bookdatastore.PublishBook(publishBooksParam)
	if err == nil {
		book_as_json, _ := json.Marshal(book)
		fmt.Fprintf(w, string(book_as_json))
	}
}
func unpublishBook(w http.ResponseWriter, r *http.Request) {
	var publishBooksParam bookdatastore.BooksWithAuthor
	publishBooksParam = publishBooksParam.FillRequest(r)
	log.Println("Handler: ", publishBooksParam)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "current_user", r.Header.Get("author")) // This is authenticated already
	err := bookdatastore.UnPublishBook(publishBooksParam, ctx)
	if err == nil {
		book_as_json, _ := json.Marshal(publishBooksParam)
		fmt.Fprintf(w, "Unpublised: "+string(book_as_json))
	}
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
