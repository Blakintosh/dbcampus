package authentication

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

/**************************** Entities/Main Objects ***************************/

// Manager data
type Manager struct {
	TeammanagerID            int
	Username                 string
	Password                 string
	SessionID                string
	ManagerSuccessPercentage string
}

// Project data
type Project struct {
	ProjectID                int
	TeamMangerID             string
	ProjectSuccess           float32
	Budget                   int64
	MonthlyOutgoings         int
	CurrentSpending          int64
	Deadline                 time.Time
	NotMetDeadlinePercentage float32
	TeamCapability           float32
	DocumentationLevel       float32
}

/**************************** Code Metrics ************************************/

// Quality of code
type CodeQuality struct {
	ProjectID             int
	NumberOfLanguagesUsed int
	Reusability           float32
	Interfacing           float32
	TestQuality           float32
	CodeErrorDensity      float32
}

// Code enviroment
type CodeEnviroment struct {
	ProjectID         int
	Stability         float32
	Complexity        float32
	Clarity           float32
	Dependence        float32
	Schedule          float32
	ObjectivesClarity float32
	DevEnviroment     float32
}

/**************************** Survey Data *************************************/

// Team metrics from team surveys
type TeamMetrics struct {
	ProjectID           int
	TaskDifficulty      float32
	ProjectSatisfaction float32
	TeamMotivation      float32
	TeamHappiness       float32
	Turnover            int32
}

// Client metrics from client surveys
type ClientMetrics struct {
	ClientID            int
	ProjectID           int
	ProductSatisfaction float32
	ScopeSatisfaction   float32
	YearlyMeetingNumber float32
}

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

var store = sessions.NewCookieStore([]byte("cookiesmakerfactory"))
var Psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func SignupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Println("Not a POST request")
		return
	}

	var inputtedUser Manager

	if req.Method == "POST" {
		// Access the request body
		reqBody, err := ioutil.ReadAll(req.Body)
		fmt.Println("reqBody register: ", string(reqBody))
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

	db, err := sql.Open("postgres", Psqlconn)
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error, unable to create your account.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	err = db.QueryRow(`SELECT "username" FROM "teammanager" WHERE username=$1`, inputtedUser.Username).Scan(&inputtedUser.Username)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inputtedUser.Password), bcrypt.DefaultCost)
		fmt.Println("hashedPassword: ", hashedPassword)
		if err != nil {
			log.Println(err)
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		_, err = db.Exec(`INSERT INTO teammanager (username, password) VALUES ($1, $2)`, inputtedUser.Username, hashedPassword)
		if err != nil {
			log.Println(err)
			http.Error(res, "Server error, unable to create your account.", http.StatusUnauthorized)
			return
		}

		log.Println("User created!")
		res.Write([]byte("User created!"))
		res.WriteHeader(200)
		return
	case err != nil:
		log.Println(err)
		http.Error(res, "Server error, unable to create your account.", http.StatusUnauthorized)
		return
	default:
		fmt.Println("User already exists")

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
	// db, err := sql.Open("postgres", Psqlconn)
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

func LoginPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		return
	}

	var inputtedUser Manager

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

	// CheckCookies(res, req)
	// if err != nil {
	// 	log.Println("couldn't find cookie. proceeding with logging in normally")
	// }

	// if loggedUsername != "" {
	// 	getProjectDataForUser(inputtedUser.Username)
	// }

	var databaseUsername string
	var databasePassword string

	db, err := sql.Open("postgres", Psqlconn)
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error, unable to create your account.", 500)
		return
	}
	defer db.Close()

	err = db.QueryRow(`SELECT "username", "password" FROM "teammanager" WHERE username=$1`, inputtedUser.Username).Scan(&databaseUsername, &databasePassword)

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

	log.Printf("User %s logged in\n", inputtedUser.Username)
	// create a new session if successful
	session, err := store.Get(req, inputtedUser.Username)
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error, unable to create your account.", 500)
		return
	}

	log.Println("session: ", session)
	// set session values
	session.Values["username"] = inputtedUser.Username
	session.Values["authenticated"] = true

	// save session
	err = session.Save(req, res)
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error, unable to store session", 500)
		return
	}

	log.Println("session saved")

	// set cookie
	cookie := http.Cookie{
		Name:     "SessionID",
		Value:    session.ID,
		Secure:   true,
		HttpOnly: true,
		MaxAge:   2628000,
	}
	http.SetCookie(res, &cookie)
	log.Println("cookie set")
	log.Println("Session ID: ", session.ID)

	// update database with session
	_, err = db.Exec(`UPDATE teammanager SET sessionid=$1 WHERE username=$2`, session.ID, inputtedUser.Username)
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error, unable to update session ID.", 500)
		return
	}

	// Select the logged in user and return their session ID
	var sessionID string
	err = db.QueryRow(`SELECT "sessionID" FROM "teammanager" WHERE username=$1`, inputtedUser.Username).Scan(&sessionID)
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error, unable to update session ID.", 500)
		return
	}

	log.Println("session ID updated in database")

	res.WriteHeader(200)
	// json.NewEncoder(res).Encode(inputtedUser)
	// json.NewEncoder(res).Encode(fmt.Sprintf("Wrote cookie %v with session id %v for user %v", cookie, session.ID, inputtedUser.Username))

}

func CheckCookies(res http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", Psqlconn)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	var sessionID string

	for _, cookie := range req.Cookies() {
		if cookie.Name == "SessionID" {
			sessionID = cookie.Value
			log.Println("Cookie session id: ", cookie.Value)
		}
	}

	// Check if there is a session cookie, if so login the user
	// If the cookie is set, verify the session
	// session := store.Get()
	var user string

	err = db.QueryRow(`SELECT "username" FROM "teammanager" WHERE session=$1`, sessionID).Scan(&user)
	if err != nil {
		log.Println(errors.New("couldn't find a user with the session id specified. Redirect to login"))
	}
	res.WriteHeader(200)
	res.Write([]byte(user))
}

func DashboardPage(res http.ResponseWriter, req *http.Request) {
	// Run the dashboard on svelte.js
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")

}

// func getProjectDataForUser(username string)

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
