package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"slices"
	"strconv"
)

type Todo struct {
	Id   int
	Task string
	Done bool
}

var data = map[string][]Todo{
	"Todos": {}}

var createdTodos = 0

func main() {
	log.Print("Server started and serving on port 8000.")

	http.HandleFunc("/", listTodos)
	http.HandleFunc("/add-todo/", addTodo)
	http.HandleFunc("/remove-todo/", removeTodo)
	http.HandleFunc("/update-todo/", updateTodo)
	http.HandleFunc("/edit-task/", startEdit)
	http.HandleFunc("/save-edit/", saveEdit)
	http.HandleFunc("/cancel-edit/", cancelEdit)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func saveEdit(w http.ResponseWriter, r *http.Request) {
	task := r.PostFormValue("task")
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	var todo *Todo = &data["Todos"][slices.IndexFunc(data["Todos"], func(t Todo) bool { return t.Id == id })]

	todo.Task = task

	tmpl := template.Must(template.ParseFiles("templates/task_view_tmpl.html"))
	tmpl.ExecuteTemplate(w, "todo-list-row", todo)

	log.Print("Todo updated: ", todo)
}

func cancelEdit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	var todo Todo = data["Todos"][slices.IndexFunc(data["Todos"], func(t Todo) bool { return t.Id == id })]
	tmpl := template.Must(template.ParseFiles("templates/task_view_tmpl.html"))
	tmpl.ExecuteTemplate(w, "todo-list-row", todo)
}

func startEdit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	var todo Todo = data["Todos"][slices.IndexFunc(data["Todos"], func(t Todo) bool { return t.Id == id })]
	tmpl := template.Must(template.ParseFiles("templates/task_input_tmpl.html"))
	tmpl.ExecuteTemplate(w, "todo-list-row", todo)
}

func listTodos(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, data)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	//time.Sleep(1 * time.Second)
	task := r.PostFormValue("task")

	createdTodos = createdTodos + 1
	newTodo := Todo{Id: createdTodos, Task: task, Done: false}

	data["Todos"] = append(data["Todos"], newTodo)

	tmpl := template.Must(template.ParseFiles("templates/task_view_tmpl.html"))
	tmpl.ExecuteTemplate(w, "todo-list-row", newTodo)
	log.Print("Created new todo: ", newTodo)
}

func removeTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	data["Todos"] = slices.DeleteFunc(data["Todos"], func(t Todo) bool { return t.Id == id })
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	var todo *Todo = &data["Todos"][slices.IndexFunc(data["Todos"], func(t Todo) bool { return t.Id == id })]

	if todo.Done {
		todo.Done = false
	} else {
		todo.Done = true
	}

	log.Print("Updated Todo: ", *todo)
}
