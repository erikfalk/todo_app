package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Todo struct {
	Task string
	Done bool
}

func main() {
	fmt.Println("Server started and serving on port 8000.")

	http.HandleFunc("/", listTodos)
	http.HandleFunc("/add-todo/", addTodo)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func listTodos(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	todos := map[string][]Todo{
		"Todos": {
			{Task: "wash dishes", Done: false},
			{Task: "do laundrey", Done: true},
		},
	}

	tmpl.Execute(w, todos)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	task := r.PostFormValue("task")

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "todo-list-element", Todo{Task: task, Done: false})
}
