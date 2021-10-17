package main

import (
	//For executing commands
	"os/exec"
	//Printing Data
	"fmt"
	//JSON library
	"encoding/json"
	//XML library
	"encoding/xml"
	//API system
	"net/http"
	"log"
)

// XML/JSON structs:
type IndexType struct{
	IndexStruct []TableStruct	`xml:"table"	json:"table"`
}

type TableStruct struct{
	Desc		string		`xml:"desc"		json:"desc"`
	Tags		[]TagStruct	`xml:"tag"		json:"tag"`
}

type TagStruct struct{
	Descs	[]string	`xml:"desc"		json:"desc"`
}

func exiftoolExecuter(w http.ResponseWriter, r *http.Request){

	//Execute the Command and get the output
	cmd := exec.Command("./runexiftool.sh")
	stdout, err := cmd.CombinedOutput()
	if err != nil{
		fmt.Printf("Error Executing the command!")
	}

	//Parse XML to a struct
	data := string(stdout)
	var s IndexType
	xml.Unmarshal([]byte(data), &s)
	// fmt.Println(s.IndexStruct)
	// fmt.Println(s)

	//Create JSON Data
	result, err := json.Marshal(s)
	if nil != err {
        fmt.Println("Error marshalling to JSON", err)
        return
    }
	// fmt.Println(string(result))
	// respond := Payload{result}
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(result)
	w.Write(result)
}

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "API system works")
}

func handleRequests(){
	// http.HandleFunc("/", homepage)
	http.HandleFunc("/", exiftoolExecuter)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main(){
	handleRequests()
}