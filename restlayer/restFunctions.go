package restlayer

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"todo/database/dbtools"
	"todo/database/model"

	"github.com/gorilla/mux"
)

// ---
// Rest Layer API Todo
// ---

func SelectTodoBasedId(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	id, ok := vars["id"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "Todo item not found.")
	}

	todo, err := dbtools.SelectTodoBasedId(id)

	if err != nil {
		json.NewEncoder(response).Encode("Todo item not found.")
	} else {
		json.NewEncoder(response).Encode(todo)
	}
}

func SelectAllTodos(response http.ResponseWriter, request *http.Request) {

	todos := dbtools.SelectAllTodos()

	json.NewEncoder(response).Encode(todos)
}

func SaveTodo(response http.ResponseWriter, request *http.Request) {

	var todo model.Todo

	err := json.NewDecoder(request.Body).Decode(&todo)

	if err != nil {
		fmt.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Could not add new Website by error: %v.", err)
		return
	}

	todoCheck := dbtools.SaveTodo(todo)

	json.NewEncoder(response).Encode(todoCheck)

}

func UpdateTodo(response http.ResponseWriter, request *http.Request) {

	var todo model.Todo

	err := json.NewDecoder(request.Body).Decode(&todo)

	if err != nil {
		fmt.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Could not update website by error: %v.", err)
		return
	}

	todoCheck := dbtools.UpdateTodo(todo)

	json.NewEncoder(response).Encode(todoCheck)

}

func DeleteTodoId(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	id, ok := vars["id"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(response, "Id todo not found.")
	}

	// convert string to int
	idTodo, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Cannot convert string to int.")
	}

	todo := dbtools.DeleteTodoId(idTodo)

	json.NewEncoder(response).Encode(todo)

}

func DeleteAllTodos(response http.ResponseWriter, request *http.Request) {

	todo := dbtools.DeleteAllTodos()

	json.NewEncoder(response).Encode(todo)

}

func SaveMultipleTodos(response http.ResponseWriter, request *http.Request) {

	var todos []model.Todo

	err := json.NewDecoder(request.Body).Decode(&todos)

	if err != nil {
		fmt.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Could not update todo by error: %v.", err)
		return
	}

	dbtools.SaveMultipleTodos(todos)

}

func UpdateCheckTodo(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	id, ok := vars["id"]

	if !ok {
		fmt.Println(ok)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Could not update todo by error: %v.", ok)
		return
	}

	idCheck, err := strconv.Atoi(id)

	if err != nil {
		log.Fatal(err.Error())
	}

	todoCheck := dbtools.UpdateCheck(idCheck)

	json.NewEncoder(response).Encode(todoCheck)

}

// ---
// HTML Files Template Handling
// ---

// --- These functions allow to show the data in HTML format on the same domain (CORS policy avoid) ---

func ShowHtmlFile(response http.ResponseWriter, request *http.Request) {

	http.ServeFile(response, request, request.URL.Path[1:])

}

// ---
// HTML Template Handling
// ---

func indexTemplateHandling(response http.ResponseWriter, request *http.Request) {

	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	files := []string{
		"templates/base_template.html",
		"templates/index_template.html",
	}

	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we can pass the slice of file
	// paths as a variadic parameter?
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.
	err = ts.ExecuteTemplate(response, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}

// ---
// Website HTML Template Handling
// ---

func createTodoTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/todo/create_todo_template.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(response, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}

func deleteAllTodosTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/todo/delete_all_todo_template.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(response, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}

func deleteIdTodoTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/todo/delete_id_todo_template.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(response, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}

func editIdTodoTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/todo/edit_todo_template.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(response, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}

func showAllTodosTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/todo/show_all_todo_template.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(response, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}

func indexHomeTemplateHandling(response http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/base_template.html",
		"templates/index_home_template.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(response, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}
