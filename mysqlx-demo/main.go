package main

import (
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var schema_person = `
CREATE TABLE IF NOT EXISTS person (
    first_name varchar(64) NOT NULL,
    last_name varchar(64) NOT NULL,
    email varchar(64)
);`
var schema_place = `
CREATE TABLE IF NOT EXISTS place (
    country varchar(64) NOT NULL,
    city varchar(64),
    telcode bigint 
);
`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

func main() {
	dsn := "user_sqlx:password_sqlx@tcp(database:3306)/db_sqlx"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema_person)
	db.MustExec(schema_place)
}
