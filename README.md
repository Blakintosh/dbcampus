# Pitstop (aka. dbcampus)
Repository for the CS261 Group 7 software engineering project

### Resources
* Notion page: https://www.notion.so/Software-Engineering-Team-7-caa7e9e594db4bb9bf90858e09e32d39
* Slack page: https://app.slack.com/client/T04LVN7U96E/C04MKD63LPJ

### How to use
Frontend and Backend work on different ports (which is different than the database one). To run the frontend go to src/frontend and run `npm run dev`. Open another terminal or shell and go to src/backend and run `go run main.go`. You should then go to `localhost:5137/auth/login` to start the website. Before that make sure to have postgresql and run pitstop.sql from src/backend/src/databases
