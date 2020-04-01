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
	return db, err
}

// AddPost adds new post to db
func AddPost(post *Posts) error {
	db, err := Connect()
	if err != nil {

	}
	_, err = db.InsertOne(post)
	if err != nil {
	}
	return err
}

// GetPosts gets
func GetPosts(offset int) []Posts {
	db, _ := Connect()
	var posts []Posts
	err := db.Table("Posts").Desc("CreatedAt").Limit(10, offset).Find(&posts)
	if err != nil {
		fmt.Println(err)
	}
	return posts
}

// GetUserPassword get user pass from db
func GetUserPassword(username string) (string, error) {
	db, err := Connect()
	if err != nil {
		print("connect error")
	}
	user := Users{Username: username}
	isFound, err := db.Table("Users").Select("Password").Get(&user)
	if err != nil {
		print("connect error")
	}
	if isFound {
		fmt.Println(user)
	}
	return user.Password, nil
}

// GetUser get user from db
func GetUser(user Users) (Users, error) {
	db, err := Connect()
	if err != nil {
		print("connect error")
	}
	isFound, err := db.Select("fullname, username, about, followers, follows").Get(&user)
	if err != nil {
		print("connect error")
	}
	if isFound {
		fmt.Println(user)
		return user, nil
	}
	return user, err
}
