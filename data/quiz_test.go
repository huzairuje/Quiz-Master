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

//=====================================================DETAIL===============================================

// TestDetailQuestionSuccess calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestDetailQuestionSuccess(t *testing.T) {
	no := "10"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	responseCreate := CreateQuestion(no, question, answer)

	detailNo := no
	responseDetail := GetDetailQuestion(detailNo)

	expectedResult := fmt.Sprintf(
		`Question no %v:
Q: %s
A: %s`, no, question, answer)

	if responseDetail != expectedResult {
		t.Fatalf(`Fatal Error the resutl %s is not match for the want match %s`, responseDetail, responseCreate)
	}
}

// TestDetailQuestionSuccess calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestDetailQuestionAndNoIsEmpty(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	detailNo := ""
	responseDetail := GetDetailQuestion(detailNo)

	trueResult := fmt.Sprintf(`Value of No: %v cannot be empty or zero`, detailNo)

	if responseDetail != trueResult {
		t.Fatalf(`Fatal Error the resutl %s is not match for the want match %s`, responseDetail, responseDetail)
	}
}

// TestDetailQuestionSuccess calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestDetailQuestionAndNoIsZero(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	detailNo := "0"
	responseDetail := GetDetailQuestion(detailNo)

	trueResult := fmt.Sprintf(`Value of No: %v cannot be empty or zero`, detailNo)

	if responseDetail != trueResult {
		t.Fatalf(`Fatal Error the resutl %s is not match for the want match %s`, responseDetail, responseDetail)
	}
}

// TestDetailQuestionSuccess calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestDetailQuestionAndQuestionIsNotFound(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	detailNo := "2"
	responseDetail := GetDetailQuestion(detailNo)

	trueResult := fmt.Sprintf(`Question no %v not found!`, detailNo)

	if responseDetail != trueResult {
		t.Fatalf(`Fatal Error the resutl %s is not match for the want match %s`, responseDetail, responseDetail)
	}
}

//=====================================================DELETE===============================================

// TestDetailQuestionSuccess calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestDeletelQuestionSuccess(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	deleteNo := no
	responseDelete := DeleteQuestion(deleteNo)

	trueResult := fmt.Sprintf(`Question no %v was deleted!`, deleteNo)

	if responseDelete != trueResult {
		t.Fatalf(`Fatal Error the result %s is not match for the want match %s`, responseDelete, trueResult)
	}
}

// TestDetailQuestionSuccess calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestDeleteQuestionAndNoIsEmpty(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	deleteNo := ""
	responseDelete := DeleteQuestion(deleteNo)

	trueResult := fmt.Sprintf(`Value of No: %v cannot be empty or zero`, deleteNo)

	if responseDelete != trueResult {
		t.Fatalf(`Fatal Error the result %s is not match for the want match %s`, responseDelete, trueResult)
	}
}

// TestDetailQuestionSuccess calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestDeleteQuestionAndNoIsZero(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	deleteNo := "0"
	responseDelete := DeleteQuestion(deleteNo)

	trueResult := fmt.Sprintf(`Value of No: %v cannot be empty or zero`, deleteNo)

	if responseDelete != trueResult {
		t.Fatalf(`Fatal Error the result %s is not match for the want match %s`, responseDelete, trueResult)
	}
}

// TestDetailQuestionSuccess calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestDeleteQuestionAndQuestionIsNotFound(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	deleteNo := "2"
	responseDelete := DeleteQuestion(deleteNo)

	trueResult := fmt.Sprintf(`Question no %v not found!`, deleteNo)

	if responseDelete != trueResult {
		t.Fatalf(`Fatal Error the result %s is not match for the want match %s`, responseDelete, trueResult)
	}
}

//=================================================CHECK QUESTION IS EXIST===============================================

// TestDetailQuestionSuccess calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestCheckQuestionIsExistIsSuccess(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	noInt, _ := strconv.Atoi(no)
	responseCheck := CheckTheQuestionIsAlreadyExist(noInt)

	trueResult := true

	if responseCheck != trueResult {
		t.Fatalf(`Fatal Error the result %t is not match for the want match %t`, responseCheck, trueResult)
	}
}

// TestDetailQuestionSuccess calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestCheckQuestionIsExistIsFalse(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	noInt, _ := strconv.Atoi("2")
	responseCheck := CheckTheQuestionIsAlreadyExist(noInt)

	trueResult := false

	if responseCheck != trueResult {
		t.Fatalf(`Fatal Error the result %t is not match for the want match %t`, responseCheck, trueResult)
	}
}

//=======================================CHECK ANSWER==========================================================

// TestUpdateQuestionAndAnswerIsEmpty calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestCheckAnswerQuestionAndCorrectIsSuccess(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	noAnswer := no
	answerUpdate := "2"
	responseAnswer := CheckTheAnswerIsCorrect(noAnswer, answerUpdate)

	trueResult := fmt.Sprintf(`Correct!`)

	if responseAnswer != trueResult {
		t.Fatalf(`Fatal Error the result %s is match for the want match %s`, responseAnswer, trueResult)
	}
}

// TestUpdateQuestionAndAnswerIsEmpty calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestCheckAnswerQuestionAndInCorrectIsSuccess(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	noAnswer := no
	answerUpdate := "3"
	responseAnswer := CheckTheAnswerIsCorrect(noAnswer, answerUpdate)

	trueResult := fmt.Sprintf(`Incorrect!`)

	if responseAnswer != trueResult {
		t.Fatalf(`Fatal Error the result %s is match for the want match %s`, responseAnswer, trueResult)
	}
}

// TestUpdateQuestionAndAnswerIsEmpty calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestCheckAnswerQuestionAndNoIsEmpty(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	noAnswer := ""
	answerUpdate := "2"
	responseAnswer := CheckTheAnswerIsCorrect(noAnswer, answerUpdate)

	trueResult := fmt.Sprintf(`No %v cannot be empty or zero`, noAnswer)

	if responseAnswer != trueResult {
		t.Fatalf(`Fatal Error the result %s is match for the want match %s`, responseAnswer, trueResult)
	}
}

// TestUpdateQuestionAndAnswerIsEmpty calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestCheckAnswerQuestionAndNoIsZero(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	noAnswer := "gh"
	answerUpdate := "2"
	responseAnswer := CheckTheAnswerIsCorrect(noAnswer, answerUpdate)

	trueResult := fmt.Sprintf(`No %v cannot be other than number!`, noAnswer)

	if responseAnswer != trueResult {
		t.Fatalf(`Fatal Error the result %s is match for the want match %s`, responseAnswer, trueResult)
	}
}

// TestUpdateQuestionAndAnswerIsEmpty calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestCheckAnswerQuestionAndQuestionIsNotFound(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	noAnswer := "2"
	answerUpdate := "2"
	responseAnswer := CheckTheAnswerIsCorrect(noAnswer, answerUpdate)

	trueResult := fmt.Sprintf(`Question no %v not found!`, noAnswer)

	if responseAnswer != trueResult {
		t.Fatalf(`Fatal Error the result %s is match for the want match %s`, responseAnswer, trueResult)
	}
}

// TestUpdateQuestionAndAnswerIsEmpty calls data.CreateQuestionSuccess
// with a No:, Question , and Answer checking
// for a valid return value.
func TestCheckAnswerQuestionAndAnswerIsEmpty(t *testing.T) {
	no := "1"
	question := "what are the result of 1 + 1 ?"
	answer := "2"
	_ = CreateQuestion(no, question, answer)

	noAnswer := "1"
	answerUpdate := ""
	responseAnswer := CheckTheAnswerIsCorrect(noAnswer, answerUpdate)

	trueResult := fmt.Sprintf(`Answer %v cannot be empty or zero`, noAnswer)

	if responseAnswer != trueResult {
		t.Fatalf(`Fatal Error the result %s is match for the want match %s`, responseAnswer, trueResult)
	}
}
