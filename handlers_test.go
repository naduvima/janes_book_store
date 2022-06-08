package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ParseResponse(res *http.Response) (string, int) {
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(contents), res.StatusCode
}

//copy this to test any handler and change URI and handler function to test
func Test_publishBook(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		publishBook(w, r)
	}

	body := []byte(`{"author":{"author_name": "jane"},"book":{"title": "Fifty shades of gray","price": 20.99}}`)

	req := httptest.NewRequest("POST", "http://jb.com/books/publish", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	parsedResp, status := ParseResponse(resp)
	if status != http.StatusOK {
		t.Error("invalid status code")
	}
	if parsedResp != `1234 jane Fifty shades of gray` {
		t.Error("invalid response", parsedResp)
	}
}

//To do:-
//Create test for unpublish

//Create test for list books

//Create test for details of a published book

//Create an end point for user which gives a password back

