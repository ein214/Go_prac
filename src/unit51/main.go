// main project main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Food struct {
	Name     string
	FoodType string
	Snack    Snack
}

type Snack struct {
	Name  string
	Count uint64
}

func main() {
	//json -> go
	sample := `[{
		"Name" : "허니콤보",
		"FoodType" : "chicken"
	}, {
		"Name" : "후렌치파이",
		"FoodType" : "snack",
		"Snack": {
			"Name" : "French Pie",
			"Count" : 100
		}
	}]`

	var data []Food
	json.Unmarshal([]byte(sample), &data)
	fmt.Println(data)

	//go -> json
	go_sample := make([]Food, 3)

	go_sample[0].Name = "짬뽕"
	go_sample[0].FoodType = "Chinese Food"

	go_sample[1].Name = "짜장면"
	go_sample[1].FoodType = "Chinese Food"

	go_sample[2].Name = "새우깡"
	go_sample[2].FoodType = "Snack"
	go_sample[2].Snack.Name = "Saewoo"
	go_sample[2].Snack.Count = 100

	data2, _ := json.MarshalIndent(go_sample, "", "  ")
	fmt.Println(string(data2))

	ioutil.WriteFile("./food.json", data2, os.FileMode(0644))
}
