package main

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:hjwan817@/account?charset=utf8")
	if err == nil {
		DB = db
		db.LogMode(true)
		db.AutoMigrate(&Page{}, &Post{}, &Tag{}, &PostTag{}, &User{}, &Comment{}, &Subscriber{}, &Link{})
		db.Model(&PostTag{}).AddUniqueIndex("uk_post_tag", "post_id", "tag_id")
		return db, err
	}
	//ddd
	return nil, err
}
