package main

import (
	"fmt"

	"github.com/dio-av/code-test/parser"
)

func main() {
	var p parser.PlanData

	data, err := p.GetJson("./data.json")
	if err != nil {
		panic(err)
	}
	p.AdjustPlanDate()
	p.Priority()
	fmt.Printf("Fisrt Data %v", data)
}
