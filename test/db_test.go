package test

import (
	"fmt"
	"testing"

	DB "github.com/onurcanari/kartaca_spa/pkg/db"
)

func TestAddPost(t *testing.T) {
	db, _ := DB.Connect()
	post := DB.Posts{Username: "admin", LikeCount: 0, Content: "Olmak ya da daha fazla olmamak."}
	err := DB.AddPost(db, &post)
	if err != nil {
		fmt.Println(err)
		t.Error("Cant add new post")
	}
}
