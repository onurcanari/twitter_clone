package db

import (
	"fmt"
	"path/filepath"

	//xorm needs qlite
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

// Connect connects to the database
func Connect() (db *xorm.Engine, err error) {
	DbPath := "./pkg/db/DB.db"
	path, _ := filepath.Abs(DbPath)
	db, err = xorm.NewEngine("sqlite3", path)
	if err != nil {
		return db, err
	}
	user := Users{Username: "admin"}

	fmt.Println(user)
	return db, err
}

// AddPost adds new post to db
func AddPost(db *xorm.Engine, post *Posts) error {
	db, err := Connect()
	if err != nil {
	}
	_, err = db.InsertOne(post)
	if err != nil {
	}
	return err
}

// GetUser get user from db
func GetUser(user *Users) {
	return
}
