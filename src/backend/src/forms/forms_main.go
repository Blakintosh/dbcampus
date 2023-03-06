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
	cmd := exec.Command("./dist/test.exe", "yop")
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println(string(out))
}

 