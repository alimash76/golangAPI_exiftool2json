package main

import (
	"os/exec"
	"fmt"
	// "encoding/json"
	"encoding/xml"
	// "strconv"
)

// XML structs:
type IndexType struct{
	IndexStruct []TableStruct	`xml:"table"`
}

type TableStruct struct{
	Desc		string		`xml:"desc"`
	Tags		[]TagStruct	`xml:"tag"`
}

type TagStruct struct{
	Descs	[]string	`xml:"desc"`
}


func main(){

	//Execute the Command and get the output
	cmd := exec.Command("./runexiftool.sh")
	stdout, err := cmd.CombinedOutput()
	if err != nil{
		fmt.Printf("Error")
	}

	//Create xml struct
	data := string(stdout)
	var s IndexType
	xml.Unmarshal([]byte(data), &s)

	// fmt.Println(s.IndexStruct)
	// fmt.Println(s)
	
	//Json
}

