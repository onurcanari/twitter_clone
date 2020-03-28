package db

import "time"

// Posts represent post table on db
type Posts struct {
	ID        int       `xorm:"int autoincr notnull pk unique 'Id'" json:"ID"`
	Username  string    `xorm:"varchar(20) notnull" json:"Username"`
	Content   string    `xorm:"text" json:"Content"`
	LikeCount int       `xorm:"int 'LikeCount'" json:"LikeCount"`
	CreatedAt time.Time `xorm:"created 'CreatedAt'" json:"CreatedAt"`
}

// Users represent user table on db
type Users struct {
	Email     string `xorm:"varchar(50) 'Email'" json:"Email"`
	Fullname  string `xorm:"varchar(50) 'Fullname'" json:"Fullname"`
	Password  string `xorm:"text" json:"Password"`
	Username  string `xorm:"varchar(20) notnull pk unique" json:"Username"`
	About     string `xorm:"text" json:"About"`
	Followers int    `xorm:"int 'Followers'" json:"Followers"`
	Follows   int    `xorm:"int 'Follows'" json:"Follows"`
	Posts     string
}
