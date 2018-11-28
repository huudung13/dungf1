package models

import (
	"github.com/asdine/storm"
)

type (
	Blog struct {
		ID          int `storm:"id,increment"`
		Title       string
		Description string
		Content     string
	}
)

var (
	db *storm.DB
)

func InitModels() (err error) {
	db, err = storm.Open("blog.db")
	return
}

func (b *Blog) Create() error {
	return db.Save(b)
}

func GetBlogs() (blogs []Blog, err error) {
	err = db.Select().Find(&blogs)
	return
}

func GetOnePost(getid int, errr error) (bblog Blog, err error) {
	err = db.One("ID", getid, &bblog)
	return
}

func (b *Blog) Update() error {
	return db.Save(b)
}

func DeleteOnePost(getid int, errr error) {
	var blog Blog
	err := db.One("ID", getid, &blog)
	if err == nil {
		db.DeleteStruct(&blog)
	}

}
