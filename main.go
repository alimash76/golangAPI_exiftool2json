package main

import (
	"os/exec"
	"fmt"
	// "strconv"
)

func main(){
	// out, cmd := exec.Command("./command.sh").Output()
	// err := cmd.Start()
	// if err != nil{
	// 	fmt.Printf("Error!")
	// }
	// fmt.Printf("wating for the command to finish...")
	// err = cmd.Wait()
	// fmt.Printf("Command finished:\n%s", out)


	// cmd := exec.Command("sleep", "5")
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Printf("Error")
	// }
	// fmt.Printf("Waiting for command to finish...")
	// err = cmd.Wait()
	// fmt.Printf("Command finished with error: %v", err)

	cmd := exec.Command("./command.sh")
	stdout, err := cmd.CombinedOutput()
	if err != nil{
		fmt.Printf("Error")
	}
	convStd := string(stdout[:])
	fmt.Printf(convStd)
	
}

