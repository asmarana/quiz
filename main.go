package main

import (
	"fmt"
	"quiz/csv"
)

func main() {
	var ans [10]string
	question, answers := csv.ReadFile()
	trueAns := 0
	falseAns := 0
	for i := range question {
		fmt.Printf("========Question%d======= \n%v\n", i+1, question[i])
		fmt.Scanf("Enter answer : ", ans[i])
		if ans[i] == answers[i] {
			trueAns++
		} else {
			falseAns++
		}
	}

	fmt.Printf("Correct Answers Are : %d\n", trueAns)
	fmt.Printf("False Answers Are : %d\n", falseAns)
}
