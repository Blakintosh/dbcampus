package main

import (
	auth "authentication"
	"database/sql"
	"log"
	"mime"
	"net/http"
	"os"
)

func main() {
	logFile, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	db, err := sql.Open("postgres", auth.Psqlconn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	mime.AddExtensionType(".js", "application/javascript")

	http.HandleFunc("/register", auth.SignupPage)
	http.HandleFunc("/login", auth.LoginPage)
	http.HandleFunc("/dashboard", auth.DashboardPage)
	http.ListenAndServe(":8081", nil)
}
