package main

import (
	"fmt"
	"log"
)

func main() {


	minuites := []int{500, 330}
	for i := 0; i < len(minuites); i++ {
		data := getTime(minuites[i])
		log.Println(data)
	}
}

func getTime(minuite int) string {
	var currentMinuite string
	var hour int
	var minuites int
	hour = time / 60
	minuites = time % 60
	currentMinuite = fmt.Sprint(minuites)

	if len(currentMinuite) == 1 {
		currentMinuite = "0" + currentMinuite
	}
	check := false
	if hour >= 12 {
		check = true
		hour = hour - 12
	}
	hoursData := fmt.Sprint(hour)

	if len(hoursData) == 1 {
		hoursData = "0" + hoursData
	}
	finalData := hoursData + ":" + currentMinuite + "PM"
	if check {
		finalData = hoursData + ":" + currentMinuite + "AM"
	}

	return finalData
}
