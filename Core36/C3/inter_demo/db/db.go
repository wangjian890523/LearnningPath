package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//global DB

var database *sql.DB

// init DB connet
func InitDB(dsn string) {
	var err error
	database, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to DB:%v", err)
	}

	if err = database.Ping(); err != nil {
		log.Fatalf("failed to ping DB:%v", err)
	}
	fmt.Println("Db connection establelished!")
}

//fetch

func fetchUsers() ([]User, error) {
	rows, err := database.Query("SELECT id, name, email  FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)

	}
	return users, nil
}

func FetchUsers() ([]User, error) {
	return fetchUsers()
}
