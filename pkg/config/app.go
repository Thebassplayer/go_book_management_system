package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("host=aws-0-us-west-1.pooler.supabase.com user=postgres.pnpblrbqmwlqwjqqkofu password=C^wSCUL3e$KJcpsGk71fMxzb port=5432 dbname=postgres")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
