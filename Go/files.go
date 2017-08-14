package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func reduceName(jobName string, mapTask int, reduceTask int) string {
	return "mrtmp." + jobName + "-" + strconv.Itoa(mapTask) + "-" + strconv.Itoa(reduceTask)
}


type kvs struct{
	Key string `json:"key"`
	Value string `json:"value"`
}

func main(){
	nReduce := 3
	jobName := "job"
	mapTaskNumber := 2
	kvs1 := kvs{"a","aa"}
	kvs2 := kvs{"v","vv"}
	list := []kvs{kvs1,kvs2}
	fmt.Println(list)	
	files := make([]*os.File,0)
	var file *os.File
	for i := 0; i < nReduce; i++ {
		outputFile := reduceName(jobName, mapTaskNumber, i)
		file, _ = os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		files = append(files, file)
		defer file.Close()
	}
		r := 2
		js, _ := json.Marshal(list)
		fmt.Println(js)
		var m []kvs
		err_ := json.Unmarshal(js, &m)
		if err_ == nil {
			fmt.Println(m)
		}
		enc := json.NewEncoder(files[r])
		err := enc.Encode(kvs1)
		if (err!=nil) {
			fmt.Println(err)
		}
}
