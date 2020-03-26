package db

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

type post struct {
	id        int    `xorm:"int autoincr not null unique 'id'"`
	username  string `xorm:"varchar(20) not null"`
	date      string `xorm:"text"`
	content   string `xorm:"text"`
	likeCount int    `xorm:"int"`
}

func Connect() (db *xorm.Engine, err error) {
	db, err = xorm.NewEngine("sqlite3", "./DB.db")
	if err != nil {
		print("error")
		panic(err)
	}
	print("readed database")
	fmt.Println(db.DBMetas())
	return
}
