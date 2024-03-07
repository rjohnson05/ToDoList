package main

import (
	"fmt"
	"strings"
)

/*
Holds the logic concerning the task list. This list has several methods for adding/removing tasks and modifying
tasks within the list.

Author: Ryan Johnson
*/

// TaskList Class containing the task list, which holds all of the tasks added by the user.
type TaskList map[Task]struct{}

// Contains Determines whether a task is already contained within the task list. The same name may not be held
// by multiple tasks.
func (set TaskList) Contains(name string) bool {
	for task := range set {
		if strings.ToLower(task.Name) == strings.ToLower(name) {
			return true
		}
	}
	return false
}

// GetTask Returns the task with the specified name.
func (set TaskList) GetTask(name string) Task {
	for task := range set {
		if task.Name == name {
			return task
		}
	}
	return Task{}
}

// AddTask Appends a task to the task list if it is not already contained within the list. Returns true if the task is
// successfully added to the task list and false if the task is already in the list.
func (set TaskList) AddTask(task Task) bool {
	if set.Contains(task.Name) {
		fmt.Println("ERROR: A task with that name already exists in your list")
		return false
	}
	set[task] = struct{}{}
	fmt.Println("Task successfully added")
	return true
}

// RemoveTask Removes the specified task from the task list. Returns true if the task is successfully removed from the task
// list and false if the task is not contained within the list.
func (set TaskList) RemoveTask(task Task) bool {
	if !set.Contains(task.Name) {
		fmt.Println("That name isn't in your task list")
		return false
	}
	delete(set, task)
	fmt.Println("Task successfully removed")
	return true
}

// ChangeTaskName Changes the name of the task to the specified value. Returns true if the name is successfully
// changed and false if the specified name is not different from the original name.
func (set TaskList) ChangeTaskName(task Task, newName string) bool {
	oldName := task.Name
	if !set.Contains(oldName) {
		fmt.Println("ERROR: That task is not in your task list")
		return false
	}

	if newName == oldName {
		fmt.Println("ERROR: You entered the same name for both the old and new name. A different name must be" +
			"provided to change the task's name.")
		return false
	}

	delete(set, task)        // Remove the task with the old name from the task list
	task.ChangeName(newName) // Change the task to have a new name
	set[task] = struct{}{}   // Add the newly named task to the task list
	fmt.Println("Name successfully changed")
	return true
}

// ChangeTaskDate Changes the name of the task to the specified value. Returns true if the name is successfully
// changed and false if the specified name is not different from the original name.
func (set TaskList) ChangeTaskDate(task Task, newDate string) bool {
	if !set.Contains(task.Name) {
		fmt.Println("ERROR: That task is not in your task list.")
		return false
	}

	if task.DueDate == newDate {
		fmt.Println("ERROR: That task already has this date. A different date must be" +
			"provided to change its value.")
		return false
	}

	delete(set, task)
	task.ChangeDate(newDate)
	set[task] = struct{}{}
	fmt.Println("Date successfully changed")
	return true
}
