package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("#### Welcome to our TodoList App! ###")

	http.HandleFunc("/", helloUser)

	http.ListenAndServe(":3000", nil)

}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Welcome to our TodoList App!"

	fmt.Fprintln(writer, greeting)
}

// func printTasks(taskItems []string) {
// 	fmt.Println("List of my Todos")
// 	for index, task := range taskItems {
// 		fmt.Printf("%d: %s\n", index+1, task)
// 	}
// }

// func addTask(taskItems []string, newTask string) []string {
// 	var updatedTaskItems = append(taskItems, newTask)
// 	return updatedTaskItems
// }
