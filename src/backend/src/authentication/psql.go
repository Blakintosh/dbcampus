package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

var psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func signupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "signup.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var user string
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error, unable to create your account.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	err = db.QueryRow(`SELECT "username" FROM "users" WHERE username=$1`, username).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		_, err = db.Exec(`INSERT INTO users (username, password) VALUES ($1, $2)`, username, hashedPassword)
		if err != nil {
			log.Println(err)
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		res.Write([]byte("User created!"))
		return
	case err != nil:
		log.Println(err)
		http.Error(res, "Server error, unable to create your account.", 500)
		return
	default:
		http.Redirect(res, req, "/", http.StatusMovedPermanently)
	}
}

func loginPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "login.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var databaseUsername string
	var databasePassword string

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error, unable to create your account.", 500)
		return
	}
	defer db.Close()

	err = db.QueryRow(`SELECT "username", "password" FROM "users" WHERE username=$1`, username).Scan(&databaseUsername, &databasePassword)

	if err != nil {
		log.Println(err)
		http.Redirect(res, req, "/login", http.StatusMovedPermanently)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		log.Println(err)
		http.Redirect(res, req, "/login", http.StatusMovedPermanently)
		return
	}

	res.Write([]byte("Hello " + databaseUsername))

}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

func main() {
	logFile, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/signup", signupPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8081", nil)
}
