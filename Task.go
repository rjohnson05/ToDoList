package main

/*
Holds the information for a task item. Many of these tasks can be stored within a TaskList item. Each task has
a name and due date, both of which may be modified by the user.
*/

// Task Class containing data for a single task within the task list. Each task must have a name and the date it
// must be accomplished by (in string format).
type Task struct {
	Name    string
	DueDate string
}

// ChangeName Changes the name of the task to the specified value. Returns true if the name is successfully
// // changed and false if the specified name is not different from the original name.
func (task *Task) ChangeName(newName string) bool {
	task.Name = newName
	return true
}

// ChangeDate Modifies the due date of the task to the specified value. Returns true if the name is successfully
// changed and false if the specified name is not different from the original name.
func (task *Task) ChangeDate(newDate string) bool {
	if newDate == task.DueDate {
		return false
	}
	task.DueDate = newDate
	return true
}
