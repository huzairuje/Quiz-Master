/*
Package data
Copyright © 2022 muhammad huzair <muhammadhuzair@gmail.com>
*/
package data

import (
	"fmt"
	"strconv"
)

type Quiz struct {
	No       int
	Question string
	Answer   string
}

var listQuiz []Quiz

// CreateQuestion : create question with parameter No, question and the answer
func CreateQuestion(No string, question string, answer string) string {
	//validate No: value
	if No == "" || No == "0" {
		response := fmt.Sprintf(`Value of No: %v cannot be empty or zero`, No)
		println(response)
		return response
	}
	noInt, err := strconv.Atoi(No)
	if err != nil {
		response := fmt.Sprintf(`Value of No: %v cannot be other than number!`, No)
		println(response)
		return response
	}

	// Validate Question: Value
	if question == "" {
		response := fmt.Sprintf(`Question %v cannot be empty`, question)
		println(response)
		return response
	}

	// Validate Answer: Value
	if answer == "" {
		response := fmt.Sprintf(`Answer %v cannot be empty`, answer)
		println(response)
		return response
	}

	isExist := CheckTheQuestionIsAlreadyExist(noInt)
	if isExist {
		response := fmt.Sprintf(`Question no %v already existed!`, No)
		println(response)
		return response
	}

	newQuiz := Quiz{
		No:       noInt,
		Question: question,
		Answer:   answer,
	}

	listQuiz = append(listQuiz, newQuiz)
	response := fmt.Sprintf(
		`Question no %v created:
Q: %s
A: %s`, No, question, answer)
	println(response)
	return response
}

// UpdateQuestion : update question with parameter No, question and the answer
func UpdateQuestion(No string, question string, answer string) string {
	if No == "" || No == "0" {
		response := fmt.Sprintf(`Value of No: %v cannot be empty or zero`, No)
		println(response)
		return response
	}
	noInt, err := strconv.Atoi(No)
	if err != nil {
		response := fmt.Sprintf(`Value of No: %v cannot be other than number!`, No)
		println(response)
		return response
	}

	// Validate Question: Value
	if question == "" {
		response := fmt.Sprintf(`Question %v cannot be empty`, question)
		println(response)
		return response
	}

	// Validate Answer: Value
	if answer == "" {
		response := fmt.Sprintf(`Answer %v cannot be empty`, answer)
		println(response)
		return response
	}

	isExist := CheckTheQuestionIsAlreadyExist(noInt)
	if !isExist {
		response := fmt.Sprintf(`Question no %v not found!`, No)
		println(response)
		return response
	}
	//delete first
	DeleteQuestion(No)
	newQuiz := Quiz{
		No:       noInt,
		Question: question,
		Answer:   answer,
	}
	listQuiz = append(listQuiz, newQuiz)
	response := fmt.Sprintf(
		`Question no %v updated:
Q: %s
A: %s`, No, question, answer)
	println(response)
	return response
}

// ListQuestion : list all question
func ListQuestion() {
	headerTable := `No | Question | Answer`
	listQuestion := ``
	println(headerTable)
	if len(listQuiz) > 0 {
		for _, v := range listQuiz {
			listQuestion = fmt.Sprintf(`%v	%s	%s`, v.No, v.Question, v.Answer)
			println(listQuestion)
		}
		return
	}
}

//GetDetailQuestion : get detail question by No
func GetDetailQuestion(No string) string {
	if No == "" || No == "0" {
		response := fmt.Sprintf(`Value of No: %v cannot be empty or zero`, No)
		println(response)
		return response
	}

	noInt, err := strconv.Atoi(No)
	if err != nil {
		response := fmt.Sprintf(`Value of No: %v cannot be other than number!`, No)
		println(response)
		return response
	}

	isExist := CheckTheQuestionIsAlreadyExist(noInt)
	if !isExist {
		response := fmt.Sprintf(`Question no %v not found!`, No)
		println(response)
		return response
	}
	response := ""
	if len(listQuiz) > 0 {
		for _, v := range listQuiz {
			if v.No == noInt {
				response = fmt.Sprintf(
					`Question no %v:
Q: %s
A: %s`, v.No, v.Question, v.Answer)
			}
		}

	}
	println(response)
	return response
}

//DeleteQuestion : delete question by No
func DeleteQuestion(No string) string {
	if No == "" || No == "0" {
		response := fmt.Sprintf(`Value of No: %v cannot be empty or zero`, No)
		println(response)
		return response
	}

	noInt, err := strconv.Atoi(No)
	if err != nil {
		response := fmt.Sprintf(`Value of No: %v cannot be other than number!`, No)
		println(response)
		return response
	}

	isExist := CheckTheQuestionIsAlreadyExist(noInt)
	if !isExist {
		response := fmt.Sprintf(`Question no %v not found!`, No)
		println(response)
		return response
	}
	response := ""
	if len(listQuiz) > 0 {
		for i, v := range listQuiz {
			if v.No == noInt {
				listQuiz = append(listQuiz[:i], listQuiz[i+1:]...)
				response := fmt.Sprintf(`Question no %v was deleted!`, No)
				println(response)
				return response
			}
		}
	}
	return response
}

//CheckTheAnswerIsCorrect : check the answer is correct
func CheckTheAnswerIsCorrect(No string, answer string) string {
	// Validate No : Value
	if No == "" || No == "0" {
		response := fmt.Sprintf(`No %v cannot be empty or zero`, No)
		println(response)
		return response
	}

	noInt, err := strconv.Atoi(No)
	if err != nil {
		response := fmt.Sprintf(`No %v cannot be other than number!`, No)
		println(response)
		return response
	}
	// Validate Answer Value
	if answer == "" {
		response := fmt.Sprintf(`Answer %v cannot be empty or zero`, No)
		println(response)
		return response
	}

	isExist := CheckTheQuestionIsAlreadyExist(noInt)
	if !isExist {
		response := fmt.Sprintf(`Question no %v not found!`, No)
		println(response)
		return response
	}
	response := ""
	if len(listQuiz) > 0 {
		for _, v := range listQuiz {
			if v.No == noInt {
				if v.Answer == answer {
					response = fmt.Sprintf(`Correct!`)
					println(response)
					return response
				} else {
					response = fmt.Sprintf(`Incorrect!`)
					println(response)
					return response
				}
			}
		}
	}
	return response
}

func CheckTheQuestionIsAlreadyExist(No int) bool {
	if len(listQuiz) > 0 {
		for _, v := range listQuiz {
			if v.No == No {
				return true
			}
		}
	}
	return false
}
