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
	var kv []kvs
	file, _ := os.Open("mrtmp.job-2-2")
	dec := json.NewDecoder(file)
	err := dec.Decode(&kv)
	if err == nil {
		fmt.Println(kv)
	}

	m := make(map[string][]string)
	s := []string{"a","b","c"}
	m["1"] = append(s)
	fmt.Println(m["1"])
	/*for err == nil {
		fmt.Println(kv)
		err = dec.Decode(&kv)
	}*/
	
}
