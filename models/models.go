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

// mỗi lần gorun lại reset database
// func initData() {
// 	db.Drop(&Blog{})
// 	blogs := []Blog{
// 		{
// 			ID:          1,
// 			Description: "Day la desc bai1",
// 			Title:       "Day la tittle bai1",
// 			Content:     "Day la content bai1",
// 		}, {
// 			ID:          2,
// 			Description: "Day la desc bai2",
// 			Title:       "Day la tittle bai2",
// 			Content:     "Day la content bai2",
// 		}, {
// 			ID:          3,
// 			Description: "Day la desc bai3",
// 			Title:       "Day la tittle bai3",
// 			Content:     "Day la content bai3",
// 		},
// 	}
// 	for _, blog := range blogs {
// 		err := db.Save(&blog)
// 		fmt.Println(blog, err)
// 	}
// }
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
