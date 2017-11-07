package database

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func DatabaseConnect() (*sql.DB,error) {

	DB_HOST  		:= "database.yourhost.com"
	DB_PASSWORD  	:= "mypassword"
	DB_USER  		:= "username"
	DB_DATABASE 	:= "database_name"
	DB_PORT			:= "3306"

	db, err := sql.Open("mysql", DB_USER+":"+DB_PASSWORD+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_DATABASE+"?charset=utf8")

	if err!=nil{
		return nil,err
	}

	return db,nil
}
