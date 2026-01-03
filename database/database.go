package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Db struct {
	Session *sql.DB
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
	db.Session = d

	err = db.Session.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("database connected and online: \n %+v \n", db.Session.Stats())
	return &db
}

func (d *Db) GetUser(userId int) User {
	var user User
	sqlQuery := "SELECT * FROM users WHERE discord_id = $1"
    err := d.Session.QueryRow(sqlQuery, userId).Scan(
        &user.id,
        &user.Discord_id,
        &user.Discord_username,
        // Add other fields here if needed
    )
	switch err {
	case sql.ErrNoRows:
		println(err.Error())
		return User{}
	default:
		return user
	}
}

func (d *Db) SaveUser(u User) {
	// todo - an upsert would be better here to keep idempotency

	sqlQuery := "INSERT INTO users (discord_id, discord_user)" +
		"VALUES ($1, $2)"

	_, err := d.Session.Exec(sqlQuery, u.Discord_id, u.Discord_username)
	if err != nil {
		fmt.Printf("unable to save user %s: %+v \n", u.Discord_username, err.Error())
	}
}
