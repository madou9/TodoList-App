package main

import (
	"fmt"
	"net/http"
)

var short = "Watch Go course"
var full = "Watch Golang Full course"

var taskItems = []string{short, full}

func main() {
	fmt.Println("#### Welcome to our TodoList App! ###")

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":3000", nil)

}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Welcome to our TodoList App!"

	fmt.Fprintln(writer, greeting)
}
func showTasks(writer http.ResponseWriter, request *http.Request) {

	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}
