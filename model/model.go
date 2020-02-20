package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

func init()  {
	db, err := sqlx.Open(`mysql`, `root:112112@tcp(127.0.0.1:3306)/community?charset=utf8&parseTime=true`)
	if err != nil{
		log.Fatalln(err.Error())
	}
	DB = db
}