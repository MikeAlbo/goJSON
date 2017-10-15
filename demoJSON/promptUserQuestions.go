package main

import (
	"bufio"
	"os"
	"fmt"
)

// prompt user issues a set of questions with multiple fields and keeps track of the user's responses

// question type
type Question struct {
	text string
}

func getAddInfoQuestions()[]Question  {
	q1 := Question{text:"What is your first name?",}
	q2 := Question{text: "What is your Last Name?",}
	q3 := Question{text: "What is your address?",}
	q4 := Question{text: "What City?",}
	q5 := Question{text: "What State?",}
	q6 := Question{text: "What is your zip-code?",}

	return []Question{q1,q2,q3,q4,q5,q6}
}

// set and return the questions for search
func getSearchQuestions() []Question  {
	return []Question{{"Please enter a search term\n"}}
}

func activateScanner() (string, error)  {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		return scanner.Text(), nil
	}
	if  err := scanner.Err(); err != nil {
		return "", err
	}

	return scanner.Text(), nil
}

func PromptUser(questions[]Question) []string {

	var answers []string

	for _,q := range questions {
		fmt.Println(q.text)
		a, err := activateScanner()
		ExitIfError(err)
		answers = append(answers,a)
	}
	return answers
}
