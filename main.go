package main

import (
	"bufio"
	"fmt"
	"os"
	"quiz/csv"
	"sync"
)

func main() {

	ch := make(chan string, 10)
	ch1 := make(chan string, 10)
	wg := &sync.WaitGroup{}

	wg.Add(1)

	go func(ch *chan string, ch1 *chan string, wg *sync.WaitGroup) {
		var ans [10]string
		question, answers := csv.ReadFile()
		trueAns := 0
		falseAns := 0
		reader := bufio.NewReader(os.Stdin)
		for i := range question {
			fmt.Printf("========Question%d======= \n%v\n", i+1, question[i])
			fmt.Println("Enter your answer : ")
			ans[i], _ = reader.ReadString('\n')
			*ch <- answers[i]
			*ch1 <- ans[i]
			if ch == ch1 {
				trueAns++
			} else {
				fmt.Printf("Correct Ans : %v but your answer is : %v\n", <-*ch, <-*ch1)
				falseAns++
			}
		}

		fmt.Println(ans)

		fmt.Printf("Correct Answers Are : %d\n", trueAns)
		fmt.Printf("False Answers Are : %d\n", falseAns)
		wg.Done()
	}(&ch, &ch1, wg)
	wg.Wait()
}
