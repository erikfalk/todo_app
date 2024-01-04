package todo

import (
	"log"
	"net/http"
	"path"
	"slices"
	"strconv"
	"text/template"
)

type Todo struct {
	Id   int
	Task string
	Done bool
}

var data = map[string][]Todo{
	"Todos": {}}

var createdTodos = 0

func SaveUpdate(w http.ResponseWriter, r *http.Request) {
	task := r.PostFormValue("task")
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	var todo *Todo = &data["Todos"][slices.IndexFunc(data["Todos"], func(t Todo) bool { return t.Id == id })]

	todo.Task = task

	tmpl := template.Must(template.ParseFiles("todo/templates/task_view_tmpl.html"))
	tmpl.ExecuteTemplate(w, "todo-list-row", todo)

	log.Print("Todo updated: ", todo)
}

func CancelUpdate(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	var todo Todo = data["Todos"][slices.IndexFunc(data["Todos"], func(t Todo) bool { return t.Id == id })]
	tmpl := template.Must(template.ParseFiles("todo/templates/task_view_tmpl.html"))
	tmpl.ExecuteTemplate(w, "todo-list-row", todo)
}

func StartUpdate(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	var todo Todo = data["Todos"][slices.IndexFunc(data["Todos"], func(t Todo) bool { return t.Id == id })]
	tmpl := template.Must(template.ParseFiles("todo/templates/task_input_tmpl.html"))
	tmpl.ExecuteTemplate(w, "todo-list-row", todo)
}

func ListTodos(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, data)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	// for spinner testing purposes
	//time.Sleep(1 * time.Second)
	task := r.PostFormValue("task")

	createdTodos = createdTodos + 1
	newTodo := Todo{Id: createdTodos, Task: task, Done: false}

	data["Todos"] = append(data["Todos"], newTodo)

	tmpl := template.Must(template.ParseFiles("todo/templates/task_view_tmpl.html"))
	tmpl.ExecuteTemplate(w, "todo-list-row", newTodo)
	log.Print("Created new todo: ", newTodo)
}

func RemoveTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	data["Todos"] = slices.DeleteFunc(data["Todos"], func(t Todo) bool { return t.Id == id })

	log.Print("Deleted todo with id: ", id)
}

func SetStatus(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	var todo *Todo = &data["Todos"][slices.IndexFunc(data["Todos"], func(t Todo) bool { return t.Id == id })]

	if todo.Done {
		todo.Done = false
	} else {
		todo.Done = true
	}

	log.Print("Updated Todo: ", *todo)
}
