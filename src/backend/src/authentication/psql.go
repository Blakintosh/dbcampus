package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
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

type User struct {
	Username string
	Password string
}

var psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func signupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		return
	}

	var inputtedUser User

	if req.Method == "POST" {
		// Access the request body
		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
			http.Error(res, "Server error, unable to create your account.", 500)
		}
		err = json.Unmarshal([]byte(reqBody), &inputtedUser)
		if err != nil {
			log.Println(err)
			http.Error(res, "Server error, unable to create your account.", 500)
		}

	}

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error, unable to create your account.", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inputtedUser.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		_, err = db.Exec(`INSERT INTO users (username, password) VALUES ($1, $2)`, inputtedUser.Username, hashedPassword)
		if err != nil {
			log.Println(err)
			http.Error(res, "Server error, unable to create your account.", http.StatusUnauthorized)
			return
		}

		res.Write([]byte("User created!"))
		res.WriteHeader(200)
		return
	case err != nil:
		log.Println(err)
		http.Error(res, "Server error, unable to create your account.", http.StatusUnauthorized)
		return
	}

	res.WriteHeader(200)
	json.NewEncoder(res).Encode(inputtedUser)

	// if req.Method != "POST" {
	// 	http.ServeFile(res, req, "signup.html")
	// 	return
	// }

	// username := req.FormValue("username")
	// password := req.FormValue("password")

	// var user string
	// db, err := sql.Open("postgres", psqlconn)
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(res, "Server error, unable to create your account.", http.StatusInternalServerError)
	// 	return
	// }
	// defer db.Close()

	// err = db.QueryRow(`SELECT "username" FROM "users" WHERE username=$1`, username).Scan(&user)

	// switch {
	// case err == sql.ErrNoRows:
	// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 	if err != nil {
	// 		log.Println(err)
	// 		http.Error(res, "Server error, unable to create your account.", 500)
	// 		return
	// 	}

	// 	_, err = db.Exec(`INSERT INTO users (username, password) VALUES ($1, $2)`, username, hashedPassword)
	// 	if err != nil {
	// 		log.Println(err)
	// 		http.Error(res, "Server error, unable to create your account.", 500)
	// 		return
	// 	}

	// 	res.Write([]byte("User created!"))
	// 	return
	// case err != nil:
	// 	log.Println(err)
	// 	http.Error(res, "Server error, unable to create your account.", 500)
	// 	return
	// default:
	// 	http.Redirect(res, req, "/", http.StatusMovedPermanently)
	// }
}

func loginPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		return
	}

	var inputtedUser User

	if req.Method == "POST" {
		// Access the request body
		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
		}
		err = json.Unmarshal([]byte(reqBody), &inputtedUser)
		if err != nil {
			log.Println(err)
		}

	}

	var databaseUsername string
	var databasePassword string

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error, unable to create your account.", 500)
		return
	}
	defer db.Close()

	err = db.QueryRow(`SELECT "username", "password" FROM "users" WHERE username=$1`, inputtedUser.Username).Scan(&databaseUsername, &databasePassword)

	if err != nil {
		log.Println(err)
		res.WriteHeader(401)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(inputtedUser.Password))
	if err != nil {
		log.Println(err)
		res.WriteHeader(401)
	}

	res.WriteHeader(200)
	json.NewEncoder(res).Encode(inputtedUser)

}

func dashboard(res http.ResponseWriter, req *http.Request) {
	// Run the dashboard on svelte.js
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")

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

	mime.AddExtensionType(".js", "application/javascript")

	http.HandleFunc("/register", signupPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/dashboard", dashboard)
	http.ListenAndServe(":8081", nil)
}

// type User struct {
// 	Username string
// 	Password string
// }
// ​
// func loginPage(res http.ResponseWriter, req *http.Request) {
// 	var tmpUser User
// 	tempUser := `{"Username":"test","Password":"test"}`
// 	var inputtedUser User
// 	err := json.Unmarshal([]byte(tempUser), &tmpUser)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println(tmpUser)
// ​
// 	if req.Method != "POST" {
// 		http.ServeFile(res, req, "login.html")
// 	}
// ​
// 	if req.Method == "POST" {
// 		// get username and password from form
// 		fmt.Println("AHHHH")
// 		// Access the request body
// 		reqBody, err := ioutil.ReadAll(req.Body)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		fmt.Println(string(reqBody))
// 		// Unmarshal the request body into a struct
// 		err = json.Unmarshal(reqBody, &inputtedUser)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		fmt.Println(inputtedUser.Username)
// 		fmt.Println(inputtedUser.Password)
// ​
// 		// receivedData := `{"Username":` + username + `,"Password":` + password + `}`
// 		// fmt.Println(receivedData)
// ​
// 		// check if username and password are correct
// 		// if username == tmpUser.Username && password == tmpUser.Password {
// 		// 	// if correct, redirect to welcome page
// 		// 	http.Redirect(res, req, "/welcome", 301)
// 		// } else {
// 		// 	// if not correct, redirect to login page
// 		// 	http.Redirect(res, req, "/login", 301)
// 		// }
// 		// err = json.Unmarshal([]byte(receivedData), &inputtedUser)
// 		// if err != nil {
// 		// 	log.Println(err)
// 		// }
// 		json.NewEncoder(res).Encode(inputtedUser)
// ​
// 	}
// ​
// 	// return tmpuser as json as response
// 	// res.Header().Set("Content-Type", "application/json")
// ​
// }
// ​
// func main() {
// 	logFile, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer logFile.Close()
// 	log.SetOutput(logFile)
// ​
// 	http.HandleFunc("/login", loginPage)
// 	http.ListenAndServe(":8081", nil)
// }
