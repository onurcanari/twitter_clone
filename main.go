package main

import (
	"log"
	"net/http"

	Server "github.com/onurcanari/kartaca_spa/pkg/server"
)

func main() {
	/* 	db, _ := DB.Connect()
	   	post := DB.Posts{Username: "admin", LikeCount: 0, Content: "Olmak ya da daha fazla olmamak."}
	   	err := DB.AddPost(db, &post)
	   	if err != nil {
	   		fmt.Println(err)
	   	} */
	r := Server.CreateServer()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8888", r))
}
