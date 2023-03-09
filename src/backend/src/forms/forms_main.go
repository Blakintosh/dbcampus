package forms

import(
    "fmt"
	"os"
	_"strings"
	_ "log"
	_ "path"
    "os/exec"
	"time"
	"io/ioutil"
	"connector"
	"errors"
	"encoding/json"
	
)
// func main(){
// 	forms()
// }

func forms(data string)(error){
	// db, errdb := connector.ConnectDB()
	// if (errdb != nil){
	// 	return errors.New("failed to connect to database")
	// }

	// defer connector.CloseDB(db)

	
	// data := "{\"projectCode\": \"something\",\"function\":\"create\",\"title\":\"survey\",\"questions\":[{\"title\":\"you are happy\",\"type\":\"scale\",\"questionID\":\"1\"},{\"title\":\"how are you doing\",\"type\":\"scale\",\"questionID\":\"2\"}]}"
	// data := "{\"projectCode\": \"something\",\"function\": \"retrieve\"}"
	
	var dat map[string]interface{}
	errDat := json.Unmarshal([]byte(data), &dat)
	if errDat != nil {
		return errDat
	}
	
	if(dat["function"].(string) == "retrieve"){
		/*get forms id using project code if retrieve else just go ahead*/
		form_id, err := getFormId(dat["projectCode"].(string))
		if (err != nil){
			return err
		}
		// form_id :="1eBp4r7shhohOAAi36qpf2Wkbj1QGJsxjmRvbXQXmJs4"
		data = "{\"projectCode\": \"something\",\"function\": \"retrieve\",\"form_id\": \"" + fmt.Sprintf("%s",form_id) + "\"}"
	}


	result, err := getResults(data)
	if err != nil{
		return err
	}
		
	if (dat["function"].(string) == "retrieve"){
		insertResults(result, dat)	
		fmt.Println("ret")
		fmt.Println(result, dat)
		
		return nil

	} else if (dat["function"].(string) == "create"){
		insertForm(result, dat)
		fmt.Println("cre")
		fmt.Println(result)
		return nil
	}

	return errors.New("bad json")
	
}

func insertResults(result map[string]interface{}, data map[string]interface{})(error){
	db, errdb := connector.ConnectDB()
	if (errdb != nil){
			return errors.New("failed to connect to database")
		}
		
	// _, err = db.Exec(`UPDATE TeamSurveys SET sessionid=$1,  WHERE username=$2`, session.ID, inputtedUser.Username)
	var err error = nil
	for k, v := range result { 

			switch k{
		case "00000001":
			_, err = db.Exec(`UPDATE TeamSurveys SET supportFromTopManagement=$1  WHERE projectCode=$2`, v, data["projectCode"].(string))
			fmt.Printf("set supportFromTopManagement to %f\n", v)
		case "00000002":
			_, err = db.Exec(`UPDATE TeamSurveys SET testingQuality=$1  WHERE projectCode=$2`, v, data["projectCode"].(string))
			fmt.Printf("set testingQuality to %f\n", v)
		case "00000003":
			_, err = db.Exec(`UPDATE TeamSurveys SET documentationQuality=$1  WHERE projectCode=$2`, v, data["projectCode"].(string))
			fmt.Printf("set documentationQuality to %f\n", v)
		case "00000004":
			_, err = db.Exec(`UPDATE TeamSurveys SET clarityOfRequirements=$1  WHERE projectCode=$2`, v, data["projectCode"].(string))
			fmt.Printf("set clarityOfRequirements to %f\n", v)
		case "00000005":
			_, err = db.Exec(`UPDATE TeamSurveys SET taskTooMuch=$1  WHERE projectCode=$2`, v, data["projectCode"].(string))
			fmt.Printf("set taskTooMuch to %f\n", v)
		case "00000006":
			_, err = db.Exec(`UPDATE TeamSurveys SET teamSatisfaction=$1  WHERE projectCode=$2`, v, data["projectCode"].(string))
			fmt.Printf("set teamSatisfaction to %f\n", v)
		}
		if (err != nil){
			return err
		}
	}


	connector.CloseDB(db)

	return nil
}

func insertForm(result map[string]interface{}, data map[string]interface{})(error){
	db, errdb := connector.ConnectDB()
	if (errdb != nil){
		return errors.New("failed to connect to database")
	}

	_, err := db.Exec(`UPDATE TeamSurveys SET formID=$1, surveyLink=$2  WHERE projectCode=$3`, result["form_id"], result["url"], data["projectCode"].(string))
	
	if (err != nil){
		return err
	}

	connector.CloseDB(db)

	return nil
}

func getFormId(projectCode string)(string, error){
	db, errdb := connector.ConnectDB()
	if (errdb != nil){
		return "", errors.New("failed to connect to database")
	}
	var form_id string
	err := db.QueryRow(`SELECT "formID" from "TeamSurveys" where projectCode = $1`, projectCode).Scan(&form_id)

	if (err != nil){
		return "", err
	}

	connector.CloseDB(db)
	return form_id,nil
	// return "",nil
}

func getResults(data string)(map[string]interface{}, error){
	cmd := exec.Command("./dist/forms_handler.exe", data)

	doneChan := make(chan bool)
	
	go func(doneChan chan bool) {
		defer func() {
			doneChan <- true
		}()

		err := watchFile("./res.json")
		if err != nil {
			
			fmt.Println(err)
		}

	}(doneChan)

	
	errRun := cmd.Run()
	<-doneChan
	
	if errRun != nil {
		println("yo")
		println(errRun.Error())
		return nil, errors.New("failed to runform handler")
	}
	
	jsonFile, err := os.Open("res.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    // fmt.Println("Successfully Opened res.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    var result map[string]interface{}

    json.Unmarshal([]byte(byteValue), &result)

	return result, nil
}

func watchFile(filePath string) error {
    initialStat, err := os.Stat(filePath)
	count := 0
    if err != nil {
        return err
    }

    for {
        stat, err := os.Stat(filePath)
        if err != nil {
            return err
        }

        if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
            break
        }
		count+=1
        time.Sleep(1 * time.Second)
		if (count == 120){
			return errors.New("timeout when performing servey retrieval/creation")
		}
    }

    return nil
}

 