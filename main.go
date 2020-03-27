package main

import (
	"fmt"

	DB "github.com/onurcanari/kartaca_spa/pkg/db"
)

func main() {
	db, _ := DB.Connect()
	post := DB.Posts{Username: "admin", LikeCount: 0, Content: "Olmak ya da daha fazla olmamak."}
	/* err := DB.AddPost(db, &post)
	if err != nil {
		fmt.Println(err)
	} */

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8888", r))
}
