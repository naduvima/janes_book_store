package main

import (
	"bytes"
	"encoding/json"
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
		publishBook(w,r)
	}
	type PostData struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		BookID int    `json:"book_id"`
	}
	
	data := PostData{"Fifty shades of gray","jane",1234}

	body, _ := json.Marshal(data)

	req := httptest.NewRequest("POST","http://jb.com/books/publish", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	parsedResp, status := ParseResponse(resp)
	if status != http.StatusOK {
		t.Error("invalid status code")
	}
	if parsedResp != `1234 jane Fifty shades of gray` {
		t.Error("invalid response",parsedResp )
	}
}
