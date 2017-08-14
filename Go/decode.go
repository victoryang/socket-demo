package main

import (
	"encoding/json"
	"os"
	"fmt"
)

type kvs struct{
    Key string `json:"key"`
    Value string `json:"value"`
}


func main(){
	file, _ := os.Open("mrtmp.job-2-2")
	dec := json.NewDecoder(file)
	err := dec.Decode(&kv)
	if err == nil {
		fmt.Println(kv)
	}
	
}
