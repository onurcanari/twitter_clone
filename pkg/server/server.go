package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	DB "github.com/onurcanari/kartaca_spa/pkg/db"
	WS "github.com/onurcanari/kartaca_spa/pkg/websocket"
)

// TODO: get this key with .env
// 256-bit key
var jwtKey = []byte("576FB6F5488F1C75CF19A477BAB3B")

var upgrader = websocket.Upgrader{}

func signin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signin")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	}
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	print(r.Body)
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

}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
		}
		fmt.Println("auth middleware")
		exceptions := [4]string{"/signin", "/logout", "/getPosts", "/livePosts"}
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
func isSigned(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	}
	return
}
func addPost(hub *WS.Hub, w http.ResponseWriter, r *http.Request) {
	fmt.Println("add post")

	username := GetNameFromToken(r)
	var post DB.Posts
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Status Bad Request", http.StatusBadRequest)
		return
	}
	post.Content = strings.Trim(post.Content, "\t \n")
	if len(post.Content) > 0 {
		post.Username = username
		err := DB.AddPost(&post)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Status Bad Request", http.StatusBadRequest)
			return
		}
		postJSON, _ := json.Marshal(post)
		hub.Broadcast <- postJSON
		return
	}
	setupResponse(&w, r)

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
		user.Username = GetNameFromToken(r)
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
	setupResponse(&w, r)

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
	setupResponse(&w, r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(postsJSON)
}

// CreateServer creates server.
func CreateServer(hub *WS.Hub) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/signin", signin).Methods("POST", "OPTIONS")
	r.HandleFunc("/logout", logout).Methods("GET", "OPTIONS")
	r.HandleFunc("/isSigned", isSigned).Methods("GET", "OPTIONS")
	r.HandleFunc("/addPost", func(w http.ResponseWriter, r *http.Request) {
		addPost(hub, w, r)
	}).Methods("POST", "OPTIONS")
	r.HandleFunc("/getPosts", getPosts).Methods("POST", "OPTIONS")
	r.HandleFunc("/getUserDetails", getUserDetails).Methods("POST", "OPTIONS")
	r.HandleFunc("/livePosts", func(w http.ResponseWriter, r *http.Request) {
		WS.Serve(hub, w, r)
	}).Methods("OPTIONS", "GET")
	r.Use(authMiddleware)
	return r
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")
}
