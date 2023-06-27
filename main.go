package main

import (
	"bufio"
	"fmt"
	"os"
	"quiz/csv"
)

func main() {
	var ans [10]string
	question, answers := csv.ReadFile()
	trueAns := 0
	falseAns := 0
	reader := bufio.NewReader(os.Stdin)
	for i := range question {
		fmt.Printf("========Question%d======= \n%v\n", i+1, question[i])
		fmt.Println("Enter your answer : ")
		ans[i], _ = reader.ReadString('\n')
		if answers[i] != ans[i] {
			trueAns++
		} else {
			fmt.Printf("Correct Ans : %v but your answer is : %v", answers[i], ans[i])
			falseAns++
		}
	}

	for j := range ans {
		fmt.Println(ans[j])
	}

	fmt.Printf("Correct Answers Are : %d\n", trueAns)
	fmt.Printf("False Answers Are : %d\n", falseAns)
}
