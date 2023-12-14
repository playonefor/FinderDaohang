package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	orm *gorm.DB
	err error
)

func Init(c db.Connection) {
	orm, err = gorm.Open("sqlite", c.GetDB("default"))

	if err != nil {
		panic("initialize orm failed")
	}
}

func GetDB() *gorm.DB {
	return orm
}
