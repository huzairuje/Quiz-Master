/*
Copyright © 2022 muhammad huzair <muhammadhuzair@gmail.com>

*/
package main

import (
	"bufio"
	"fmt"
	"github.com/quipper/quiz_master/data"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/quipper/quiz_master/cmd"
)

func main() {
	_, err := fmt.Fprintln(os.Stdout, cmd.InitCmd)
	if err != nil {
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return
			}
		}
		err = runCommand(cmdString)
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return
			}
		}
	}
}

func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "exit":
		_, err := fmt.Fprintln(os.Stdout, "Are you sure you want to exit? (y/n)")
		if err != nil {
			return err
		}
		answer, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		answer = strings.TrimSuffix(answer, "\n")
		if answer != "y" {
			return nil
		}
		os.Exit(0)
	case "help":
		_, err := fmt.Fprintln(os.Stdout, cmd.HelpString)
		if err != nil {
			return err
		}
		return nil
	case "create_question":
		//id
		_, err := fmt.Fprintln(os.Stdout, "No: ")
		if err != nil {
			return err
		}
		id, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		id = strings.TrimSuffix(id, "\n")
		idInt, err := strconv.Atoi(id)
		//question
		_, err = fmt.Fprintln(os.Stdout, "Question: ")
		if err != nil {
			return err
		}
		question, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		question = strings.TrimSuffix(question, "\n")
		//answer
		_, err = fmt.Fprintln(os.Stdout, "Answer: ")
		if err != nil {
			return err
		}
		answer, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		answer = strings.TrimSuffix(answer, "\n")
		data.CreateQuestion(idInt, question, answer)
		_, err = fmt.Fprintln(os.Stdout)
		if err != nil {
			return err
		}
		return nil
	case "update_question":
		//id
		_, err := fmt.Fprintln(os.Stdout, "No: ")
		if err != nil {
			return err
		}
		id, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		id = strings.TrimSuffix(id, "\n")
		idInt, err := strconv.Atoi(id)
		//question
		_, err = fmt.Fprintln(os.Stdout, "Question: ")
		if err != nil {
			return err
		}
		question, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		question = strings.TrimSuffix(question, "\n")
		//answer
		_, err = fmt.Fprintln(os.Stdout, "Answer: ")
		if err != nil {
			return err
		}
		answer, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		answer = strings.TrimSuffix(answer, "\n")
		data.UpdateQuestion(idInt, question, answer)
		_, err = fmt.Fprintln(os.Stdout)
		if err != nil {
			return err
		}
		return nil
	case "questions":
		data.ListQuestion()
		return nil
	case "question":
		argsInt, err := strconv.Atoi(arrCommandStr[1])
		if err != nil {
			return err
		}
		data.GetDetailQuestion(argsInt)
		return nil
	case "delete_question":
		argsInt, err := strconv.Atoi(arrCommandStr[1])
		if err != nil {
			return err
		}
		data.DeleteQuestion(argsInt)
		return nil
	case "answer_question":
		argsInt, err := strconv.Atoi(arrCommandStr[1])
		if err != nil {
			return err
		}
		if arrCommandStr[2] == "" {
			_, err := fmt.Fprintln(os.Stderr, "Answer cannot be empty")
			if err != nil {
				return err
			}
		}
		data.CheckTheAnswerIsCorrect(argsInt, arrCommandStr[2])
		return nil
	default:
		_, err := fmt.Fprintf(os.Stdout, "%s is unknown command\n", arrCommandStr[0])
		if err != nil {
			return err
		}
	}
	command := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	return command.Run()
}
