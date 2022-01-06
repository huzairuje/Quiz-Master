/*
Package data
Copyright © 2022 muhammad huzair <muhammadhuzair@gmail.com>
*/
package data

import "fmt"

type Quiz struct {
	No       int
	Question string
	Answer   string
}

var listQuiz []Quiz

// CreateQuestion : create question with parameter No, question and the answer
func CreateQuestion(No int, question string, answer string) {
	isExist := CheckTheQuestionIsAlreadyExist(No)
	if isExist {
		response := fmt.Sprintf(`Question no %v already existed!`, No)
		println(response)
		return
	}
	newQuiz := Quiz{
		No:      No,
		Question: question,
		Answer:   answer,
	}
	listQuiz = append(listQuiz, newQuiz)
	response := fmt.Sprintf(
		`Question no %v created:
Q: %s
A: %s`, No, question, answer)
	println(response)
	return
}

// UpdateQuestion : update question with parameter No, question and the answer
func UpdateQuestion(No int, question string, answer string) {
	isExist := CheckTheQuestionIsAlreadyExist(No)
	if !isExist {
		response := fmt.Sprintf(`Question no %v not found!`, No)
		println(response)
		return
	}
	//delete first
	DeleteQuestion(No)
	newQuiz := Quiz{
		No:      No,
		Question: question,
		Answer:   answer,
	}
	listQuiz = append(listQuiz, newQuiz)
	response := fmt.Sprintf(
		`Question no %v updated:
Q: %s
A: %s`, No, question, answer)
	println(response)
	return
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
	return
}

//GetDetailQuestion : get detail question by No
func GetDetailQuestion(No int) {
	isExist := CheckTheQuestionIsAlreadyExist(No)
	if !isExist {
		response := fmt.Sprintf(`Question no %v not found!`, No)
		println(response)
		return
	}
	response := ""
	if len(listQuiz) > 0 {
		for _, v := range listQuiz {
			if v.No == No {
				response = fmt.Sprintf(
					`Question no %v:
Q: %s
A: %s`, v.No, v.Question, v.Answer)
			}
		}
		println(response)
		return
	}
}

//DeleteQuestion : delete question by No
func DeleteQuestion(No int)  {
	isExist := CheckTheQuestionIsAlreadyExist(No)
	if !isExist {
		response := fmt.Sprintf(`Question no %v not found!`, No)
		println(response)
		return
	}
	if len(listQuiz) > 0 {
		for i, v := range listQuiz {
			if v.No == No {
				listQuiz = append(listQuiz[:i], listQuiz[i+1:]...)
				response := fmt.Sprintf(`Question no %v was deleted!`, No)
				println(response)
				return
			}
		}
	}
}

//CheckTheAnswerIsCorrect : check the answer is correct
func CheckTheAnswerIsCorrect(No int, answer string) {
	isExist := CheckTheQuestionIsAlreadyExist(No)
	if !isExist {
		response := fmt.Sprintf(`Question no %v not found!`, No)
		println(response)
		return
	}
	if len(listQuiz) > 0 {
		for _, v := range listQuiz {
			if v.No == No {
				if v.Answer == answer {
					response := fmt.Sprintf(`Correct!`)
					println(response)
					return
				}
				response := fmt.Sprintf(`Incorrect!`)
				println(response)
				return
			}
		}
	}
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
