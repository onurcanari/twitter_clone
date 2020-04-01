package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	DB "github.com/onurcanari/kartaca_spa/pkg/db"
)

// TODO: get this key with .env
// 256-bit key
var jwtKey = []byte("576FB6F5488F1C75CF19A477BAB3B")

func signin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signin")
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

	expirationTime := time.Now().Add(50 * time.Minute)
	claims := &JwtClaims{
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
		Name:    "x-csrf-token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	w.Header().Set("Auth", "true")
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("auth middleware")
		exceptions := [2]string{"/signin", "/logout"}
		for _, val := range exceptions {
			if r.URL.String() == val {
				next.ServeHTTP(w, r)
				return
			}
		}
		_, token, err := GetClaimsToken(r)

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
		fmt.Println("validated.")
		next.ServeHTTP(w, r)

	})
}

func logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "x-csrf-token",
		Value: "LOGOUT",
	})

}
func addPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add post")

	username := GetNameFromToken(r)
	var post DB.Posts
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Status Bad Request", http.StatusBadRequest)
		return
	}

	if len(post.Content) > 0 {
		post.Username = username
		DB.AddPost(&post)
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
}
func getUserDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get user details")

	var user DB.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Print("user: ")
		fmt.Println(user)
		fmt.Println(err)
		http.Error(w, "Status Bad Request", http.StatusBadRequest)
		return
	}
	if user.Username == "" {
		http.Error(w, "Username must be entered.", http.StatusBadRequest)
		return
	}
	user, err = DB.GetUser(user)
	if err != nil {
		http.Error(w, "Cant get user details.", http.StatusNotFound)
		return
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
	}
	initHeader(w.Header())
	w.Write(userJSON)
	return
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	var decodedJSON map[string]int
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &decodedJSON)
	offset := decodedJSON["offset"]
	print(decodedJSON)
	posts := DB.GetPosts(offset)
	postsJSON, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, "Parse Error", http.StatusServiceUnavailable)
		return
	}
	initHeader(w.Header())
	w.Write(postsJSON)
}

// CreateServer creates server.
func CreateServer() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/signin", signin).Methods("POST")
	r.HandleFunc("/logout", logout).Methods("GET")
	r.HandleFunc("/addPost", addPost).Methods("POST")
	r.HandleFunc("/getPosts", getPosts).Methods("POST")
	r.HandleFunc("/getUserDetails", getUserDetails).Methods("POST")
	r.Use(authMiddleware)
	return r
}
func initHeader(h http.Header) {
	h.Set("Content-Type", "application/json")
	h.Set("Access-Control-Allow-Origin", "*")
	h.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	h.Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

/*
func refresh(w http.ResponseWriter, r *http.Request) {
	cokie, err := r.Cookie("x-csrf-token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := cokie.Value
	claims := &JwtClaims{}
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
		Name:    "x-csrf-token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}
*/
