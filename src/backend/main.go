package main

import (
	auth "authentication"
	connector "connector"
	dash "dashboard"
	"log"
	"net/http"
	"os"

	mux "github.com/gorilla/mux"
)

func main() {
	logFile, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	log.Println("Server started")

	db, err := connector.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	defer connector.CloseDB(db)

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	auth.Init()

	router := mux.NewRouter()

	router.HandleFunc("/register", auth.SignupPage)
	router.HandleFunc("/login", auth.LoginPage)
	router.HandleFunc("/dashboard/project", dash.ProjectPage)
	router.HandleFunc("/dashboard", dash.DashboardPage)
	router.HandleFunc("/newProject", dash.CreateProject)
	err = http.ListenAndServe(":8081", router)
	log.Println("Server stopper with error: ", err)
}
