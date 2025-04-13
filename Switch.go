package main

import "fmt"

func main() {
	fmt.Println("#### Welcome to our TodoList App! ###")

	var Short = "Watch Go crash course"
	var full = "Watch Golang Full course"

	var taskItems = []string{Short, full}

	printTasks(taskItems)
	fmt.Println()
	taskItems = addTask(taskItems, "Go for a run")
	taskItems = addTask(taskItems, "Practice coding in Go")

	fmt.Println("updated list")
	printTasks(taskItems)
}

func printTasks(taskItems []string) {
	fmt.Println("List of my Todos")
	for index, task := range taskItems {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}
