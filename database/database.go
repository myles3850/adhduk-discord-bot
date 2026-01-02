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

type User struct {
	id               int
	Discord_id       int
	Discord_username string
}

func Setup() *Db {
	var db Db
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

func (d *Db) GetUser(userId int) *User {
	var user User
	sqlQuery := "SELECT * FROM users WHERE discord_id = $1"
	err := d.session.QueryRow(sqlQuery, userId).Scan(&user)
	switch err {
	case sql.ErrNoRows:
		return nil
	default:
		return &user
	}
}

func (d *Db) SaveUser(u *User) {
	// todo - an upsert would be better here to keep idempotency
	sqlQuery := "INSERT INTO users (discord_id, discord_username)" +
		"VALUES ($1, $2)"

	_, err := d.session.Query(sqlQuery, u.Discord_id, u.Discord_username)
	if err != nil {
		fmt.Printf("unable to save user %s: %+v", u.Discord_username, err.Error())
	}
}
