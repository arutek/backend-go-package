package db

import (
	"fmt"

	"github.com/arutek/backend-go-package/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormInit(
	host string,
	user string,
	pass string,
	name string,
	port string,
	sslMode string,
	tz string,
	disableTrx bool,
) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, pass, name, port, sslMode, tz)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: disableTrx,
	})
	if err != nil {
		helper.LoggerErr(err.Error())
		panic(err.Error())
	}
	helper.Logger(fmt.Sprintf("DB %s connection open", name))
	return db
}

func PaginateDb(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if size < 0 {
			size = 10
		}
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
