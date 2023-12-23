package main

import (
	"fmt"
	"math"
	"time"
)

type Element struct {
	// parentID int
	// elementID int
	// content Content
	// priority int
	
	/*
	* 0 = Article type
	* 1 = Item type
	* 2 = Tasks
	*/
	//elementType int
	interval int
	grade int
	EF  float64
	//prevRepDate time.Time //TODO need to update prevRepDate and nextRepDate for every update
	nextRepDate time.Time
}

// type Content struct {
// 	imageIDs 	[]int
// 	templateIDs	[]int
// 	textIDs		[]int
// 	referenceIDs	[]int
// 	videoIDs	[]int
// }


func main() {
	date := time.Now()
	myElement := Element{
		interval: 4,
		grade:    3,
		EF:       1.7,
		nextRepDate: date,}
	fmt.Printf("%+v\n", myElement)
	fmt.Println()
	fmt.Println("updating element")
	updateLearningData(&myElement)
	fmt.Printf("%+v\n", myElement)
}

func updateLearningData(element *Element){
	element.EF = calculateNewEF(element.EF, element.grade)
	element.interval = calculateNewInterval(element.EF, element.interval)
	element.nextRepDate = calculateNextRepDate(element.nextRepDate, element.interval)
}
	

func calculateNewInterval(newEF float64, oldInterval int) int{
	newInterval := int(math.Round(newEF * float64(oldInterval)))
	return newInterval
}

	// takes in prev rep date and interval
	// returns next rep date
func calculateNextRepDate(prevRepDate time.Time, interval int) time.Time {
	nextRepDate := prevRepDate.Add(time.Duration(interval) * 24 * time.Hour)
	return nextRepDate
}

	// takes in oldEF and grade
	// returns newEF
func calculateNewEF(oldEF float64, grade int) float64 {
	newEF := oldEF - 0.8 + 0.28 * float64(grade) - 0.02 * float64(math.Pow(float64(grade), 2))
	return newEF
}
