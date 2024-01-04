package main

import (
	"log"
	"net/http"

	"github.com/erikfalk/todo_app/todo"
)

func main() {
	log.Print("Server started and serving on port 8000.")

	http.HandleFunc("/", todo.ListTodos)
	http.HandleFunc("/add-todo/", todo.AddTodo)
	http.HandleFunc("/remove-todo/", todo.RemoveTodo)
	http.HandleFunc("/set-status/", todo.SetStatus)
	http.HandleFunc("/start-update/", todo.StartUpdate)
	http.HandleFunc("/save-update/", todo.SaveUpdate)
	http.HandleFunc("/cancel-update/", todo.CancelUpdate)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
