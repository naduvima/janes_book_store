package jwt_secure_access

import (
	"fmt"
	bookdatastore "janes_book_store/book_data_store"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func decodetRequest(tokenString string, user string) bool {
	type MyCustomClaims struct {
		Payload  string `json:"data"`
		User     string `json:"user"`
		Password string `json:"password"`
		jwt.StandardClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return getPasswordForUsers(user), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
		return false
	}
	return true
}

func getPasswordForUsers(user string) []byte {
	//users := map[string]string{"jane": "eyJkYXRhIjoidGVzdCJ9", "margo": "eyJhbGciOiJIUzI1NiJ9", "nair": "8h8W4cmvJwIX3UUp9J5yf41ax3"}
	author, _ := bookdatastore.FindAuthor(bookdatastore.Author{AuthorName: user})
	log.Println("Password found: ", author.Password)
	return []byte(author.Password)
}

//read header for password and constructed token string
//decode with predefined users, password
func AuthorizeRequest(r *http.Request) bool {
	user := r.Header.Get("author")
	tokenString := r.Header.Get("token")
	if user == "" || tokenString == "" {
		return false
	}
	return decodetRequest(tokenString, user)
}
