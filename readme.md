# Quiz Master
## Introduction
This is Shell Interactive CLI Apps Quiz Master.

## Getting Started
1.	Installation process
2.	Usage

### 1. Installation process
you can run the program with the following way:

 a. by running the program by this command (golang way):
   1. firstly, you can install go programming language.
   1. clone this project to your local machine, by this command
       ```
       $ git clone https://github.com/huzairuje/Quiz-Master.git
       ```
   2. go to the directory of the project, by this command
       ```
       $ cd Quiz-Master
       ```
   3. and install the dependencies by this command
       ```
       $ go mod download
       $ go mod tidy
       $ go mod verify
       ```
   4. run the program by this command
       ```
       $ go run main.go
       ```
    
 b. by running the program by this command (bash way):
   1. firstly, clone this project to your local machine, by this command
     ```
     $ git clone https://github.com/huzairuje/Quiz-Master.git
     ```
   2. go to the directory of the project, by this command
        ```
        $ cd Quiz-Master
        ```
   4. run the program by this command
        ```
        $ ./bin/quiz_master
        ```
 c. by running the program by this command (docker way):
 ```shell
$ docker build -f Dockerfile -t quiz_master .
```  
 ```shell
$ docker run quizmaster:latest
```

## 2. Usage
you can running the program itself by type this command

```shell
$ ./bin/quiz_master
```
or

```shell
$ cd ROOT_PROJECT_FOLDER/bin
```
and

```shell
$ ./quiz_master
```

you can list the command by typing this command

```shell
$ help
```
the result is
```shell
Command | Description
create_question | Creates a question
	No:<no> 
	Question: <question> 
	Answer: <answer>
update_question | Updates a question
	No:<no> 
	Question: <question> 
	Answer: <answer>
delete_question <no> | Deletes a question
question <no> | Shows a question
questions | Shows question list
```
