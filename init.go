package main

import (
	"log"

	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
)

const (
	errPlatformUnstable = "Platform unstable. Do not proceed."
)

var (
	gormDatabase *gorm.DB

	sessionStore = sessions.NewCookieStore([]byte("the-road-ahead"))
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	var err error
	gormDatabase, err = gorm.Open("mysql", "root:  @/bernd?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "fx_" + defaultTableName
	}

	gormDatabase.AutoMigrate(&User{})

	gormDatabase.AutoMigrate(&Job{})
	gormDatabase.AutoMigrate(&Connection{})
}
