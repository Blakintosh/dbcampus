package authentication

import (
	"connector"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

/**************************** Entities/Main Objects ***************************/
type Manager struct {
	TeammanagerID            int    `json:"teammanagerid"`
	Username                 string `json:"username"`
	Password                 string `json:"password"`
	SessionID                string `json:"sessionid"`
	ManagerSuccessPercentage string `json:"managersuccesspercentage"`
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

var Store = sessions.NewCookieStore([]byte("cookiesmakerfactory"))

func Init() {
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}
}

func CreateSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func SignupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Println("Not a POST request")
		http.Error(res, "GET request. Should be POST", 500)
		return
	}

	var inputtedUser Manager
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

	if inputtedUser.Username == "" || inputtedUser.Password == "" {
		log.Println(err)
		http.Error(res, "Server error, No username or password.", 500)
		return
	}

	db, err := connector.ConnectDB()
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error, unable to create your account.", http.StatusInternalServerError)
		return
	}
	defer connector.CloseDB(db)

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
	// json.NewEncoder(res).Encode(inputtedUser)

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

	// Access the request body
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal([]byte(reqBody), &inputtedUser)
	if err != nil {
		log.Println(err)
	}
	if inputtedUser.Username == "" || inputtedUser.Password == "" {
		log.Println(err)
		http.Error(res, "Server error, No username or password.", http.StatusUnauthorized)
		return
	}

	err = HasSessionAlready(res, req, inputtedUser)
	if err == nil {
		log.Println("User already logged in")
		// change to json
		json.NewEncoder(res).Encode("{\"message\": \"User already logged in\"}")
		res.WriteHeader(200)
		return
	}

	log.Println("User not logged in. Proceeding with login")

	// CheckCookies(res, req)
	// if err != nil {
	// 	log.Println("couldn't find cookie. proceeding with logging in normally")
	// }

	// if loggedUsername != "" {
	// 	getProjectDataForUser(inputtedUser.Username)
	// }

	var databaseUsername string
	var databasePassword string

	db, err := connector.ConnectDB()
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error. Unable to access database.", 500)
		return
	}
	defer connector.CloseDB(db)

	err = db.QueryRow(`SELECT "username", "password" FROM "teammanager" WHERE username=$1`, inputtedUser.Username).Scan(&databaseUsername, &databasePassword)
	if err != nil {
		log.Println(err)
		http.Error(res, "Username doesn't exist. You will need to register first", 401)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(inputtedUser.Password))
	if err != nil {
		log.Println(err)
		http.Error(res, "Wrong password.", http.StatusUnauthorized)

	}

	log.Printf("User %s logged in\n", inputtedUser.Username)
	// create a new session if successful
	session, err := Store.Get(req, "session")
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error, unable to get session.", 500)
		return
	}

	log.Println("session: ", session)
	// set session values
	session.ID = CreateSessionID()
	session.Values["username"] = inputtedUser.Username
	session.Values["authenticated"] = true
	session.Values["sessionId"] = session.ID

	// save session
	err = session.Save(req, res)
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error, unable to store session", 500)
		return
	}
	// set cookie
	cookie := http.Cookie{
		Name:     "session_token",
		Value:    session.ID,
		Secure:   true,
		HttpOnly: true,
		MaxAge:   2628000,
		SameSite: http.SameSiteNoneMode,
	}
	log.Println("cookie: ", cookie)
	//http.SetCookie(res, &cookie)
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
	err = db.QueryRow(`SELECT "sessionid" FROM "teammanager" WHERE username=$1`, inputtedUser.Username).Scan(&sessionID)
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error, unable to update session ID.", 500)
		return
	}

	log.Println("session ID updated in database")

	// json.NewEncoder(res).Encode(inputtedUser)
	log.Println("Wrote cookie", cookie, "with session id", session.ID, "for user", inputtedUser.Username)
	res.WriteHeader(200)

}

// if err is nil then the user is logged in so give 301 and redirect to home page
// if err is not nil then the user is not logged in so give 401 and redirect to login page
func HasSessionAlready(res http.ResponseWriter, req *http.Request, inputtedUser Manager) error {
	db, err := connector.ConnectDB()
	if err != nil {
		log.Println(err)
	}
	defer connector.CloseDB(db)

	var sessionID string

	// Check if there is a session cookie, if so login the user
	// If the cookie is set, verify the session
	// session := store.Get()
	var user string

	session, err := Store.Get(req, "session")

	if err != nil || session.Values["authenticated"] == nil {
		log.Println("Unable to get session: ", err)
		return errors.New("couldn't find a user with the session id specified. Redirect to login")
	}

	sessionID = session.Values["sessionId"].(string)
	log.Println("Session ID: ", sessionID)

	err = db.QueryRow(`SELECT "username" FROM "teammanager" WHERE sessionid=$1`, sessionID).Scan(&user)
	if err != nil {
		log.Println("Error checking session: ", err)
	}
	if user != "" {
		log.Printf("User %s logged in with session ID %s. Redirect to dashboard", user, sessionID)
		return nil
	}
	return errors.New("couldn't find a user with the session id specified. Redirect to login")
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
