package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/erikfalk/todo_app/todo"
)

func main() {

	portp := flag.Int("port", 8000, "The port the server should use.")

	flag.Parse()
	log.Printf("Server started and serving on port %d.", *portp)

	http.HandleFunc("/", todo.ListTodos)
	http.HandleFunc("/add-todo/", todo.AddTodo)
	http.HandleFunc("/remove-todo/", todo.RemoveTodo)
	http.HandleFunc("/set-status/", todo.SetStatus)
	http.HandleFunc("/start-update/", todo.StartUpdate)
	http.HandleFunc("/save-update/", todo.SaveUpdate)
	http.HandleFunc("/cancel-update/", todo.CancelUpdate)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*portp), nil))
}
