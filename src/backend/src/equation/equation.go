package equation

import (
	"connector"
	"jira"
	"log"
	"math"
	"sort"
	"time"
)

// Equation is a struct that holds the name of the equation, the max weight of the equation and the weight of the equation
type Equation struct {
	Name       string
	MaxWeight  float64
	Weight     float64
	Suggestion string
}

// ManagerWeights is a slice of Equation that holds the weights of the manager equations
var ManagerWeights = []Equation{
	{Name: "Manager expereince", MaxWeight: 0.7, Suggestion: "Contact top managment for support on managment choices to make sure you are making the right decisions for the team"},
	{Name: "Weekly team meetings", MaxWeight: 0.42, Suggestion: "Having 3 to 4 meetings per week is the ideal number of meetings for a team to have."},
	{Name: "Mean Team Experience", MaxWeight: 0.14, Suggestion: "Having a team with a mean experience of 5 years is the ideal number of years for a team to have. Consider getting more expereinced developers"},
	{Name: "Client meetings per month", MaxWeight: 0.2, Suggestion: "Having 4 meetings per month is the ideal number for client meetings."},
	{Name: "Budget", MaxWeight: 0.7, Suggestion: "Having a budget surplus is the ideal situation for a project but not having a budget deficit is very acceptable and is what we expect to be the case. Consider increasing the budget."},
	{Name: "Overdue tasks", MaxWeight: 0.6, Suggestion: "Having 0 overdue tasks is the ideal situation but the more you have the less the score will be, and if you have too much the score will be negative. Consider increasing the number of developers, reallocating tasks or rethinking the scale of the project."},
}

// SurveyWeights is a slice of Equation that holds the weights of the survey equations
var SurveyWeights = []Equation{
	{Name: "Support from top management", MaxWeight: 0.15, Suggestion: "Team members don't feel supported by top management. Consider making top management more included, or having more meetings with top management with team members."},
	{Name: "testing quality", MaxWeight: 0.21, Suggestion: "Testing quality is not considered good by the team. Consider improving the tests or automating it."},
	{Name: "Documentation quality", MaxWeight: 0.21, Suggestion: "Documentation quality is not considered good by the team. Consider talking to the team about their documentation practices or hiring a documentation team."},
	{Name: "Clarity of the requirements/misunderstandings", MaxWeight: 0.31, Suggestion: "The team is not clear about the requirements. Consider discussing this with your team members or having the client write the requirements in more detail."},
	{Name: "Task too much for the team", MaxWeight: 0.7, Suggestion: "The team feels like the tasks are too much for them. Consider increasing the number of developers, reallocating tasks or rethinking the scale of the project."},
	{Name: "Team satisfaction/motivation", MaxWeight: 0.14, Suggestion: "The team is not satisfied with the project. Consider talking to the team about their concerns."},
}

// GetMaxPossibleWeightForInput returns the maximum possible weight for the input we have
// It is the function responsible for the personalization of the equation
func GetMaxPossibleWeightForInput(e []Equation) float64 {
	var total float64
	for _, eq := range e {
		if eq.Weight > 0 {
			total += eq.MaxWeight
		}
	}
	return total
}

// Gets the manager score given the data needed from the database and jira if available
func MangerScore(budget float64, deadline time.Time, monthlyExpenses float64, customSpendings float64, name string, jiraUrl string, email string, token string) float64 {
	managerExperience := ManagerWeights[0].MaxWeight * (math.Log(math.Min(7.9, math.Max(1, ManagerWeights[0].Weight))) / math.Log(7.9))
	weeklyTeamMeetings := ManagerWeights[1].MaxWeight * (math.Log(math.Min(4, math.Max(1, ManagerWeights[1].Weight))) / math.Log(4))
	meanTeamExperience := ManagerWeights[2].MaxWeight * (math.Log(math.Min(4.8, math.Max(1, ManagerWeights[2].Weight))) / math.Log(4.8))
	clientMeetingsPerMonth := ManagerWeights[3].MaxWeight * (math.Log(math.Min(4, math.Max(1, ManagerWeights[3].Weight))) / math.Log(4))
	budgetScore := ManagerWeights[4].MaxWeight * ((budgetScore(budget, deadline, monthlyExpenses, customSpendings) + 1) / 2)
	overdueTasks := ManagerWeights[5].MaxWeight * overdueTasksScore(name, jiraUrl, email, token)
	return (managerExperience + weeklyTeamMeetings + meanTeamExperience + clientMeetingsPerMonth + budgetScore + overdueTasks)
}

// Gets the survey score given the data needed from the database
func SurveyScore() float64 {
	supportFromTopManagement := SurveyWeights[0].MaxWeight * (math.Log(math.Max(1, SurveyWeights[0].Weight)) / math.Log(5))
	testingQuality := SurveyWeights[1].MaxWeight * (math.Log(math.Max(1, SurveyWeights[1].Weight)) / math.Log(5))
	documentationQuality := SurveyWeights[2].MaxWeight * (math.Log(math.Max(1, SurveyWeights[2].Weight)) / math.Log(5))
	clarityOfTheRequirements := SurveyWeights[3].MaxWeight * (math.Log(math.Max(1, SurveyWeights[3].Weight)) / math.Log(5))
	taskTooMuchForTheTeam := SurveyWeights[4].MaxWeight * (math.Log(math.Max(1, SurveyWeights[4].Weight)) / math.Log(5))
	teamSatisfaction := SurveyWeights[5].MaxWeight * (math.Log(math.Max(1, SurveyWeights[5].Weight)) / math.Log(5))
	return (supportFromTopManagement + testingQuality + documentationQuality + clarityOfTheRequirements + taskTooMuchForTheTeam + teamSatisfaction)
}

// Gets the budget a score given the monthly expenses, budget and deadline
func budgetScore(budget float64, deadline time.Time, monthlyExpenses float64, customSpendings float64) float64 {
	score := budget - ((((deadline.Sub(time.Now())).Hours())/730.5)*monthlyExpenses + customSpendings)
	if score < 0 {
		return -1
	} else if score > 0 {
		return 1
	} else {
		return 0
	}
}

// Gets the overdue tasks score given the project name, jira url, email and token
func overdueTasksScore(projectCode string, jiraUrl string, email string, token string) float64 {
	overdue, total, err := jira.GetNumOverDueIssuesFromProject(projectCode, jiraUrl, email, token)
	if err != nil {
		return 0
	}
	return 3*(float64(overdue)/float64(total)) - 2
}

// Gets the success percentage given the data provided
func GetPercentage(username string, projectCode string) (float64, error) {
	db, err := connector.ConnectDB()
	if err != nil {
		return -1, err
	}
	defer connector.CloseDB(db)

	// Manager data not in ManagerWeights
	var jiraEmail string
	var jiraToken string

	// Project data not in ManagerWeights
	var budget float64
	var deadline time.Time
	var monthlyExpenses float64
	var jiraUrl string
	var customSpendings float64

	// Get the data from the database
	err = db.QueryRow(`SELECT managerExperience FROM TeamManager WHERE username=$1`, username).
		Scan(&ManagerWeights[0].Weight)
	if err != nil {
		log.Println("error getting manager experience: ", err)
	}
	err = db.QueryRow(`SELECT jiraEmail FROM TeamManager WHERE username=$1`, username).
		Scan(&jiraEmail)
	if err != nil {
		log.Println("error getting jira email: ", err)
	}

	err = db.QueryRow(`SELECT jiraApiToken FROM TeamManager WHERE username=$1`, username).
		Scan(&jiraToken)
	if err != nil {
		log.Println("error getting jira token: ", err)
	}

	err = db.QueryRow(`SELECT budget FROM project WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&ManagerWeights[4].Weight)
	if err != nil {
		log.Println("error getting budget: ", err)
	}

	err = db.QueryRow(`SELECT deadline FROM project WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&deadline)
	if err != nil {
		log.Println("error getting deadline: ", err)
	}

	err = db.QueryRow(`SELECT monthlyExpenses FROM project WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&monthlyExpenses)
	if err != nil {
		log.Println("error getting monthly expenses: ", err)
	}

	err = db.QueryRow(`SELECT jiraURL FROM project WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&jiraUrl)
	if err != nil {
		log.Println("error getting jira url: ", err)
	}

	err = db.QueryRow(`SELECT teamMeanExperience FROM project WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&ManagerWeights[2].Weight)
	if err != nil {
		log.Println("error getting team mean experience: ", err)
	}

	err = db.QueryRow(`SELECT customSpendings FROM project WHERE projectCode=$1 AND username =$2`, projectCode, username).
		Scan(&customSpendings)
	if err != nil {
		log.Println("error getting custom spendings: ", err)
	}

	err = db.QueryRow(`SELECT weeklyTeamMeetings FROM project WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&ManagerWeights[1].Weight)
	if err != nil {
		log.Println("error getting weekly team meetings: ", err)
	}

	err = db.QueryRow(`SELECT clientMeetingsPerMonth FROM project WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&ManagerWeights[3].Weight)
	if err != nil {
		log.Println("error getting client meetings per month: ", err)
	}

	err = db.QueryRow(`SELECT supportFromTopManagement FROM survey WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&SurveyWeights[0].Weight)
	if err != nil {
		log.Println("error getting support from top management: ", err)
	}

	err = db.QueryRow(`SELECT testingQuality FROM survey WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&SurveyWeights[1].Weight)
	if err != nil {
		log.Println("error getting testing quality: ", err)
	}

	err = db.QueryRow(`SELECT documentationQuality FROM survey WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&SurveyWeights[2].Weight)
	if err != nil {
		log.Println("error getting documentation quality: ", err)
	}

	err = db.QueryRow(`SELECT clarityOfRequirements FROM survey WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&SurveyWeights[3].Weight)
	if err != nil {
		log.Println("error getting clarity of requirements: ", err)
	}
	err = db.QueryRow(`SELECT taskTooMuch FROM survey WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&SurveyWeights[4].Weight)
	if err != nil {
		log.Println("error getting task too much: ", err)
	}

	err = db.QueryRow(`SELECT teamSatisfaction FROM survey WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&SurveyWeights[5].Weight)
	if err != nil {
		log.Println("error getting team satisfaction: ", err)
	}

	// Change budget weight and overdue tasks weight depending on the data
	if ManagerWeights[4].Weight != 0 {
		ManagerWeights[4].Weight = 1
	}
	overdue, _, _ := jira.GetNumOverDueIssuesFromProject(projectCode, jiraUrl, jiraEmail, jiraToken)
	if overdue != -1 {
		ManagerWeights[5].Weight = 1
	}

	// Get data needed from jira and calculate the scores
	totalScore := (MangerScore(budget, deadline, monthlyExpenses, customSpendings, projectCode, jiraUrl, jiraEmail, jiraToken) +
		SurveyScore()) / (GetMaxPossibleWeightForInput(ManagerWeights) + GetMaxPossibleWeightForInput(SurveyWeights))

	return totalScore, nil
}

// Gets the top 4 suggestions for the project by going through the data provided and finding what is weighing the project down the most and give its suggestion back
func GetSuggestions(username string, projectCode string) ([]string, error) {
	db, err := connector.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer connector.CloseDB(db)

	// Manager data not in ManagerWeights
	var jiraEmail string
	var jiraToken string

	// Project data not in ManagerWeights
	var deadline time.Time
	var monthlyExpenses float64
	var jiraUrl string
	var customSpendings float64

	// Get the data from the database
	err = db.QueryRow(`SELECT managerExperience FROM TeamManager WHERE username=$1`, username).
		Scan(&ManagerWeights[0].Weight)
	if err != nil {
		log.Println("error getting manager experience: ", err)
	}

	err = db.QueryRow(`SELECT jiraEmail FROM TeamManager WHERE username=$1`, username).
		Scan(&jiraEmail)
	if err != nil {
		log.Println("error getting jira email: ", err)
	}

	err = db.QueryRow(`SELECT jiraApiToken FROM TeamManager WHERE username=$1`, username).
		Scan(&jiraToken)
	if err != nil {
		log.Println("error getting jira token: ", err)
	}

	err = db.QueryRow(`SELECT budget FROM project WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&ManagerWeights[4].Weight)
	if err != nil {
		log.Println("error getting budget: ", err)
	}

	err = db.QueryRow(`SELECT deadline FROM project WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&deadline)
	if err != nil {
		log.Println("error getting deadline: ", err)
	}

	err = db.QueryRow(`SELECT monthlyExpenses FROM project WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&monthlyExpenses)
	if err != nil {
		log.Println("error getting monthly expenses: ", err)
	}

	err = db.QueryRow(`SELECT jiraURL FROM project WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&jiraUrl)
	if err != nil {
		log.Println("error getting jira url: ", err)
	}

	err = db.QueryRow(`SELECT teamMeanExperience FROM project WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&ManagerWeights[2].Weight)
	if err != nil {
		log.Println("error getting team mean experience: ", err)
	}

	err = db.QueryRow(`SELECT customSpendings FROM project WHERE projectCode=$1 AND username =$2`, projectCode, username).
		Scan(&customSpendings)
	if err != nil {
		log.Println("error getting custom spendings: ", err)
	}

	err = db.QueryRow(`SELECT weeklyTeamMeetings FROM project WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&ManagerWeights[1].Weight)
	if err != nil {
		log.Println("error getting weekly team meetings: ", err)
	}

	err = db.QueryRow(`SELECT clientMeetingsPerMonth FROM project WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&ManagerWeights[3].Weight)
	if err != nil {
		log.Println("error getting client meetings per month: ", err)
	}

	err = db.QueryRow(`SELECT supportFromTopManagement FROM survey WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&SurveyWeights[0].Weight)
	if err != nil {
		log.Println("error getting support from top management: ", err)
	}

	err = db.QueryRow(`SELECT testingQuality FROM survey WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&SurveyWeights[1].Weight)
	if err != nil {
		log.Println("error getting testing quality: ", err)
	}

	err = db.QueryRow(`SELECT documentationQuality FROM survey WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&SurveyWeights[2].Weight)
	if err != nil {
		log.Println("error getting documentation quality: ", err)
	}

	err = db.QueryRow(`SELECT clarityOfRequirements FROM survey WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&SurveyWeights[3].Weight)
	if err != nil {
		log.Println("error getting clarity of requirements: ", err)
	}
	err = db.QueryRow(`SELECT taskTooMuch FROM survey WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&SurveyWeights[4].Weight)
	if err != nil {
		log.Println("error getting task too much: ", err)
	}

	err = db.QueryRow(`SELECT teamSatisfaction FROM survey WHERE projectCode=$1 AND username=$2`, projectCode, username).
		Scan(&SurveyWeights[5].Weight)
	if err != nil {
		log.Println("error getting team satisfaction: ", err)
	}

	// Change budget weight and overdue tasks weight depending on the data
	if ManagerWeights[4].Weight != 0 {
		ManagerWeights[4].Weight = 1
	}
	overdue, _, _ := jira.GetNumOverDueIssuesFromProject(projectCode, jiraUrl, jiraEmail, jiraToken)
	if overdue != -1 {
		ManagerWeights[5].Weight = 1
	}

	// Sort weights in increasing order
	sort.Slice(ManagerWeights, func(i, j int) bool {
		return ManagerWeights[i].Weight < ManagerWeights[j].Weight
	})
	sort.Slice(SurveyWeights, func(i, j int) bool {
		return SurveyWeights[i].Weight < SurveyWeights[j].Weight
	})

	// Get top 4 lowest weights
	var suggestions []string
	for _, weight := range ManagerWeights {
		if weight.Weight > 0 && len(suggestions) < 4 {
			suggestions = append(suggestions, weight.Suggestion)
		}
	}
	for _, weight := range SurveyWeights {
		if weight.Weight > 0 && len(suggestions) < 4 {
			suggestions = append(suggestions, weight.Suggestion)
		}
	}
	return suggestions, nil
}
