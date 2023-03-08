package main

import(
    "fmt"
	"os"
	_"strings"
	_ "log"
	_ "path"
    "os/exec"
	"time"
	"io/ioutil"
	// "github.com/fsnotify/fsnotify"
	
)

func main(){
	
	test:= "{\"function\": \"retrieve\",\"form_id\": \"1eBp4r7shhohOAAi36qpf2Wkbj1QGJsxjmRvbXQXmJs4\"}"
	
	cmd := exec.Command("./dist/forms_handler.exe", test)
	doneChan := make(chan bool)

	go func(doneChan chan bool) {
		defer func() {
			doneChan <- true
		}()

		err := watchFile("./res.txt")
		if err != nil {
			fmt.Println(err)
		}

		// fmt.Println("File has been changed")
	}(doneChan)

	
	errRun := cmd.Run()
	<-doneChan
	
	if errRun != nil {
		println("yo")
		println(errRun.Error())
		return
	}
	

	// time.Sleep(20*time.Second)
	// println("hmm")


	// file, err := os.Open("res.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()
	content, err := ioutil.ReadFile("res.txt")

     if err != nil {
          fmt.Println(err)
     }

    fmt.Println(string(content))

	// fmt.Println(file)
	// println(string(out))
	
}

func watchFile(filePath string) error {
    initialStat, err := os.Stat(filePath)
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

        time.Sleep(1 * time.Second)
    }

    return nil
}

 