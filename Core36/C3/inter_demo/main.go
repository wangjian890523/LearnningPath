package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wangjian890523/LearnningPath/Core36/C3/inter_demo/api"
	"github.com/wangjian890523/LearnningPath/Core36/C3/inter_demo/db"
)

func main() {
	// init
	db.InitDB("username:password@tcp(localhost:3306)/dbname")

	// route
	http.HandleFunc("/users", api.GetUsersHandler)
	fmt.Println("server is running on:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
