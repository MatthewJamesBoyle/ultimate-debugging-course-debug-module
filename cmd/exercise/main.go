package main

import (
	"fmt"
	"log"
	"net/http"
	"ultimate-debugging-course-debug-module/internal/todo"
)

func main() {

	svc, err := todo.NewService()
	if err != nil {
		log.Fatal(err)
	}

	svr, err := todo.NewServer(svc)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/todos", svr.CreateToDo)
	http.HandleFunc("/todos/", svr.GetToDoHandler)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
