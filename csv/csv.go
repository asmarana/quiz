package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

// var wg sync.WaitGroup

type Problem struct {
	question string
	solution string
}

func CreateProblemsList(records [][]string) (QuestionsList []string, AnswersList []string) {

	var problemList []Problem
	// var QuestionsList []string
	for i, line := range records {
		if i > 0 {
			var rec Problem
			for j, feild := range line {
				if j == 0 {
					rec.question = feild
				} else if j == 1 {
					rec.solution = feild
				}
			}
			problemList = append(problemList, rec)
			QuestionsList = append(QuestionsList, rec.question)
			AnswersList = append(AnswersList, rec.solution)
		}
	}

	fmt.Println(problemList)
	return QuestionsList, AnswersList

}

func ReadFile() (questions []string, answers []string) {
	// defer wg.Done()
	file, error := os.Open("csv/problems.csv")

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println("Successfully open csv file")
	defer file.Close()

	fileReader := csv.NewReader(file)
	records, error := fileReader.ReadAll()

	if error != nil {
		fmt.Println(error)
	}

	Questions, Answers := CreateProblemsList(records)

	return Questions, Answers
}
