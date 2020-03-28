package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// TODO: get this key with .env
// 256-bit key
var jwtKey = []byte("576FB6F5488F1C75CF19A477BAB3B")

func signin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	w.Header().Add("Auth", "false")
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}
	isValidated := ValidateUserPassword(&creds)

	if !isValidated {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &jwtClaims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	w.Header().Set("Auth", "true")
}

func authMiddleware(next http.Handler) http.Handler {
	fmt.Println("validating...")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		exceptions := [2]string{"/signin", "/logout"}
		for _, val := range exceptions {
			if r.URL.String() == val {
				next.ServeHTTP(w, r)
				return
			}
		}

		c, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "No Cookie", http.StatusForbidden)
			return
		}
		tokenString := c.Value
		claims := &jwtClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return

		}
		if !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
		fmt.Println("validated.")

	})
}
func welcome(w http.ResponseWriter, r *http.Request) {

	c, _ := r.Cookie("token")

	tokenString := c.Value
	claims := &jwtClaims{}
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	fmt.Fprint(w, "Hello ", claims.Username)
}

func logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: "LOGOUT",
	})

}

func refresh(w http.ResponseWriter, r *http.Request) {
	cokie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := cokie.Value
	claims := &jwtClaims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expirationTime := time.Now().Add(1 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}

func CreateServer() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/signin", signin).Methods("POST")
	r.HandleFunc("/logout", logout).Methods("GET")
	r.HandleFunc("/home", welcome).Methods("GET")
	r.Use(authMiddleware)
	return r
}
