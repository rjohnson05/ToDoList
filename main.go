package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Serves as the controller for the TaskList console-based program. Receives all user input from the console and
displays it back to the user. Tasks may be added and removed from the task list, as well as having their name
date modified. No tasks with duplicate names may be created.

Author: Ryan Johnson
*/

// Scans the text input from the user, used whenever information is needed regarding how to change a task within
// task list.
func getInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	return scanner.Text()
}

// IsValidInput Determines if a string is a valid input. Valid input strings include "view", "add", "del",
// "chname", "chdate", and "quit". Returns true if the string is equal to one of these valid inputs and false
// otherwise.
func IsValidInput(userInput string) bool {
	validInput := []string{"view", "add", "del", "chname", "chdate", "quit", "help"}

	userInput = strings.TrimSpace(userInput)
	for _, input := range validInput {
		if input == userInput {
			return true
		}
	}
	return false
}

// DisplayList Prints the contents of the task list to the console.
func DisplayList(taskList TaskList) {
	fmt.Println("###########\n TASK LIST \n###########")
	if len(taskList) == 0 {
		fmt.Println("Empty List")
	}
	for task := range taskList {
		fmt.Printf("%s  -  Due: %s\n", task.Name, task.DueDate)
	}
}

// AddTaskInput Prompts the user to supply the information for a new task to be added to the task list.
func AddTaskInput(taskList TaskList) {
	var nameInput string
	var dateInput string

	for {
		// Loops until a valid name is received
		nameInput = getInput("Enter a name for the task (press 'q' to exit): ")
		nameInput = strings.TrimSpace(nameInput)

		// Quits the function
		if nameInput == "q" {
			return
		}
		// Repeats if no name is provided
		if nameInput == "" {
			fmt.Println("ERROR: You must enter a name.")
			continue
		}
		// Repeats if the name provided already exists within the task list
		if taskList.Contains(nameInput) {
			fmt.Println("ERROR: A task with that name already exists in your list")
			continue
		}
		break
	}

	for {
		// Loops until a valid date is received
		dateInput = getInput("Enter a date for the task (press 'q' to exit): ")
		dateInput = strings.TrimSpace(dateInput)

		// Quits the function
		if dateInput == "q" {
			return
		}
		// Repeats if no date is provided
		if dateInput == "" {
			fmt.Println("ERROR: You must enter a date.")
			continue
		}
		break
	}

	task := Task{nameInput, dateInput}
	taskList.AddTask(task)
}

// RemoveTaskInput Prompts the user to supply information about which task to remove from the task list.
func RemoveTaskInput(taskList TaskList) {
	var nameInput string
	for {
		nameInput = getInput("Enter the name of the task to remove (press 'q' to exit): ")
		nameInput = strings.TrimSpace(nameInput)

		if nameInput == "q" {
			return
		}
		if nameInput == "" {
			fmt.Println("ERROR: You must enter a name.")
			continue
		}
		if !taskList.Contains(nameInput) {
			fmt.Println("ERROR: That name isn't in your task list")
			continue
		}
		break
	}
	task := taskList.GetTask(nameInput)
	taskList.RemoveTask(task)
}

// ChangeNameInput Prompts the user to supply information about which task's name is to be changed, as well
// as what the new name is to be.
func ChangeNameInput(taskList TaskList) {
	var nameInput string
	var newNameInput string

	for {
		nameInput = getInput("Enter the name of the task to be changed (press 'q' to exit): ")
		nameInput = strings.TrimSpace(nameInput)

		if nameInput == "q" {
			return
		}
		if nameInput == "" {
			fmt.Println("ERROR: You must enter a name.")
			continue
		}
		if !taskList.Contains(nameInput) {
			fmt.Println("ERROR: That name isn't in your task list.")
			continue
		}
		break
	}

	for {
		newNameInput = getInput("Enter the new name for the task (press 'q' to exit): ")
		newNameInput = strings.TrimSpace(newNameInput)

		if newNameInput == "q" {
			return
		}
		if newNameInput == "" {
			fmt.Println("ERROR: You must enter a new name.")
			continue
		}
		if strings.ToLower(newNameInput) == strings.ToLower(nameInput) {
			fmt.Println("ERROR: The new name cannot be the same as the original name.")
			continue
		}
		break
	}

	task := taskList.GetTask(nameInput)
	taskList.ChangeTaskName(task, newNameInput)
}

// ChangeDateInput Prompts the user to supply information about which task's date is to be changed, as well
// as what the new date is to be.
func ChangeDateInput(taskList TaskList) {
	var nameInput string
	var newDateInput string

	for {
		nameInput = getInput("Enter the name of the task to be changed (press 'q' to exit): ")
		nameInput = strings.TrimSpace(nameInput)

		if nameInput == "q" {
			return
		}
		if nameInput == "" {
			fmt.Println("ERROR: You must enter a name.")
			continue
		}
		if !taskList.Contains(nameInput) {
			fmt.Println("ERROR: That name isn't in your task list.")
			continue
		}
		break
	}

	task := taskList.GetTask(nameInput)
	for {
		newDateInput = getInput("Enter the new date for the task (press 'q' to exit): ")
		newDateInput = strings.TrimSpace(newDateInput)

		if newDateInput == "q" {
			return
		}
		if newDateInput == "" {
			fmt.Println("ERROR: You must enter a new date.")
			continue
		}
		if strings.ToLower(newDateInput) == task.DueDate {
			fmt.Println("ERROR: The new date cannot be the same date as the original date.")
			continue
		}
		break
	}

	taskList.ChangeTaskDate(task, newDateInput)
}

// ShowCommands Prints all valid commands to the console.
func ShowCommands() {
	fmt.Println("###################\n POSSIBLE COMMANDS \n###################\n" +
		"- view: View To-Do List\n" +
		"- add: Add Task\n" +
		"- del: Remove Task\n" +
		"- chname: Change Name of Task\n" +
		"- chdate: Change Due Date of Task\n" +
		"- help: Show a list of possible commands\n" +
		"- quit: Exit the Program")
}

func main() {
	fmt.Println("Welcome to your To-Do List Manager! Enter a command to modify your list" +
		" (type 'help' to see possible commands).")

	taskList := make(TaskList)

	var running = true
	for running {
		input := getInput("-> ")

		validInput := false
		for !validInput {
			if !IsValidInput(input) {
				fmt.Println("Command not recognized. Please try again (type 'help' to see possible commands).")
				break
			}
			validInput = true
		}

		switch strings.TrimSpace(input) {
		case "view":
			DisplayList(taskList)
		case "add":
			AddTaskInput(taskList)
		case "del":
			RemoveTaskInput(taskList)
		case "chname":
			ChangeNameInput(taskList)
		case "chdate":
			ChangeDateInput(taskList)
		case "help":
			ShowCommands()
		case "quit":
			running = false
		}
	}
}
