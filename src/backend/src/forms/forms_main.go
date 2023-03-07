package main

import(
    "fmt"
	_ "os"
	_ "strings"
	_ "log"
	_ "path"
    "os/exec"

)

func main(){

	test:= "{
		\"function\": \"retrieve\",
		\"form_id\": \"1eBp4r7shhohOAAi36qpf2Wkbj1QGJsxjmRvbXQXmJs4\"
	}"
	
	cmd := exec.Command("./dist/forms_handler.exe", test)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println(string(out))
}

 