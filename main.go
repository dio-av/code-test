package main

import (
	"encoding/json"
	"fmt"

	"github.com/dio-av/code-test/parser"
)

func main() {
	var p parser.PlanData

	data, err := p.GetJson("./data.json")
	if err != nil {
		panic(err)
	}
	data.AdjustPlanDate()
	priorityList := data.PriorityFilter()
	//fmt.Printf("Fisrt Data %v\n", data)
	fmt.Println(priorityList)

	result, _ := json.MarshalIndent(priorityList, "", "  ")
	send, _ := json.Marshal(priorityList)
	fmt.Println(string(result))
	fmt.Println(string(send))
}
