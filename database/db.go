package database

import (
	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func InitDB() {
	var err error
	Engine, err = xorm.NewEngine("sqlite", "./test.db")

	if err != nil {
		panic(err)
	}

	SyncDB()
}

func SyncDB() {
	if err := Engine.Sync2(new(User)); err != nil {
		panic(err)
	}
}
