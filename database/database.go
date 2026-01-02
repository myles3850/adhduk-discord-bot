package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Db struct {
	session *sql.DB
}

func Setup() *Db {
	db := Db{}
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	database := os.Getenv("DATABASE_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pass, database)

	d, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.session = d
	defer db.session.Close()

	err = db.session.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("database connected and online: \n %+v \n", db.session.Stats())
	return &db
}

func (d Db) getUser (userId int) {
	sqlQuery := "SELECT * FROM USERS WHERE 
