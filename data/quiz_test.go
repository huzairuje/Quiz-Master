package data

import (
	"fmt"
	"strconv"
	"testing"
)

// TestCreateQuestionSuccess calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestCreateQuestionSuccess(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	NoInt, _ := strconv.Atoi(no)
	response := CreateQuestion(no, question, answer)
	trueResult := fmt.Sprintf(
		`Question no %v created:
Q: %s
A: %s`, NoInt, question, answer)

	if response != trueResult {
		t.Fatalf(`Fatal Error the resutl %s is not match for the want match %s`, response, trueResult)
	}
}

// TestCreateQuestionAndNoIsEmpty calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestCreateQuestionAndNoIsEmpty(t *testing.T) {
	No := ""
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	response := CreateQuestion(No, question, answer)
	trueResult := fmt.Sprintf(`Value of No: %v cannot be empty or zero`, No)

	if response != trueResult {
		t.Fatalf(`Fatal Error the resutl %s is not match for the want match %s`, response, trueResult)
	}
}

// TestCreateQuestionAndNoZero calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestCreateQuestionAndNoZero(t *testing.T) {
	No := "0"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	response := CreateQuestion(No, question, answer)
	trueResult := fmt.Sprintf(`Value of No: %v cannot be empty or zero`, No)

	if response != trueResult {
		t.Fatalf(`Fatal Error the resutl %s is not match for the want match %s`, response, trueResult)
	}
}

// TestCreateQuestionAndNoIsNotNumber calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestCreateQuestionAndNoIsNotNumber(t *testing.T) {
	No := "gh"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	response := CreateQuestion(No, question, answer)
	trueResult := fmt.Sprintf(`Value of No: %v cannot be other than number!`, No)

	if response != trueResult {
		t.Fatalf(`Fatal Error the resutl %s is not match for the want match %s`, response, trueResult)
	}
}

// TestCreateQuestionAndQuestionIsEmpty calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestCreateQuestionAndQuestionIsEmpty(t *testing.T) {
	No := "1"
	question := ""
	answer := "2"
	response := CreateQuestion(No, question, answer)
	trueResult := fmt.Sprintf(`Question %v cannot be empty`, question)

	if response != trueResult {
		t.Fatalf(`Fatal Error the resutl %s is not match for the want match %s`, response, trueResult)
	}
}

// TestCreateQuestionAndAnswerIsEmpty calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestCreateQuestionAndAnswerIsEmpty(t *testing.T) {
	No := "1"
	question := "what are the result of 1 + 1 ?"
	answer := ""
	response := CreateQuestion(No, question, answer)
	trueResult := fmt.Sprintf(`Answer %v cannot be empty`, answer)

	if response != trueResult {
		t.Fatalf(`Fatal Error the resutl %s is not match for the want match %s`, response, trueResult)
	}
}

//=====================================================UPDATE===============================================

// TestUpdateQuestionSuccess calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestUpdateQuestionSuccess(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	responseCreate := CreateQuestion(no, question, answer)

	noUpdate := "1"
	questionUpdate := "what are the result is one plus one ?"
	answerUpdate := "2"
	responseUpdate := UpdateQuestion(noUpdate, questionUpdate, answerUpdate)

	if responseCreate == responseUpdate {
		t.Fatalf(`Fatal Error the result %s is match for the want match %s`, responseCreate, responseUpdate)
	}
}

// TestUpdateQuestionAndNoIsEmpty calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestUpdateQuestionAndNoIsEmpty(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	noUpdate := ""
	questionUpdate := "what are the result is one plus one ?"
	answerUpdate := "2"
	responseUpdate := UpdateQuestion(noUpdate, questionUpdate, answerUpdate)

	trueResult := fmt.Sprintf(`Value of No: %v cannot be empty or zero`, noUpdate)

	if responseUpdate != trueResult {
		t.Fatalf(`Fatal Error the result %s is match for the want match %s`, responseUpdate, trueResult)
	}
}

// TestUpdateQuestionAndNoZero calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestUpdateQuestionAndNoZero(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	noUpdate := "0"
	questionUpdate := "what are the result is one plus one ?"
	answerUpdate := "2"
	responseUpdate := UpdateQuestion(noUpdate, questionUpdate, answerUpdate)

	trueResult := fmt.Sprintf(`Value of No: %v cannot be empty or zero`, noUpdate)

	if responseUpdate != trueResult {
		t.Fatalf(`Fatal Error the result %s is match for the want match %s`, responseUpdate, trueResult)
	}
}

// TestUpdateQuestionAndNoIsNotNumber calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestUpdateQuestionAndNoIsNotNumber(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	noUpdate := "gh"
	questionUpdate := "what are the result is one plus one ?"
	answerUpdate := "2"
	responseUpdate := UpdateQuestion(noUpdate, questionUpdate, answerUpdate)

	trueResult := fmt.Sprintf(`Value of No: %v cannot be other than number!`, noUpdate)

	if responseUpdate != trueResult {
		t.Fatalf(`Fatal Error the result %s is match for the want match %s`, responseUpdate, trueResult)
	}
}

// TestUpdateQuestionAndQuestionIsEmpty calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestUpdateQuestionAndQuestionIsEmpty(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	noUpdate := "1"
	questionUpdate := ""
	answerUpdate := "2"
	responseUpdate := UpdateQuestion(noUpdate, questionUpdate, answerUpdate)

	trueResult := fmt.Sprintf(`Question %v cannot be empty`, questionUpdate)

	if responseUpdate != trueResult {
		t.Fatalf(`Fatal Error the result %s is match for the want match %s`, responseUpdate, trueResult)
	}
}

// TestUpdateQuestionAndAnswerIsEmpty calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestUpdateQuestionAndAnswerIsEmpty(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	noUpdate := "1"
	questionUpdate := "what are the result is one plus one ?"
	answerUpdate := ""
	responseUpdate := UpdateQuestion(noUpdate, questionUpdate, answerUpdate)

	trueResult := fmt.Sprintf(`Answer %v cannot be empty`, answerUpdate)

	if responseUpdate != trueResult {
		t.Fatalf(`Fatal Error the result %s is match for the want match %s`, responseUpdate, trueResult)
	}
}
