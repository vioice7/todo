package restlayer

import (
	"net/http"

	"github.com/gorilla/mux"
)

func restConfig(router *mux.Router) {

	restRouter := router.PathPrefix("/todo/api").Subrouter()

	httpRouter := router.PathPrefix("").Subrouter()

	// ------------
	// Todos API
	// ------------

	// localhost:8080/todo/api/todos
	restRouter.Methods("GET").Path("/todos").HandlerFunc(SelectAllTodos)

	//localhost:8080/todo/api/todo/id/{id}
	restRouter.Methods("GET").Path("/todo/id/{id}").HandlerFunc(SelectTodoBasedId)

	//localhost:8080/todo/api/todo/add
	restRouter.Methods("PUT").Path("/todo/add").HandlerFunc(SaveTodo)

	//localhost:8080/todo/api/todo/edit
	restRouter.Methods("POST").Path("/todo/edit").HandlerFunc(UpdateTodo)

	//localhost:8080/todo/api/todo/deleteid/{id}
	restRouter.Methods("DELETE").Path("/todo/deleteid/{id}").HandlerFunc(DeleteTodoId)

	//localhost:8080/todo/api/todo/deleteall
	restRouter.Methods("DELETE").Path("/todo/deleteall").HandlerFunc(DeleteAllTodos)

	//localhost:8080/todo/api/todo/addmultiple
	restRouter.Methods("POST").Path("/todo/addmultiple").HandlerFunc(SaveMultipleTodos)

	//localhost:8080/todo/api/todo/check/{id}
	restRouter.Methods("POST").Path("/todo/check/{id}").HandlerFunc(UpdateCheckTodo)

	// ---
	// File Show System
	// ---

	//localhost:8080/configHtmlServer.json
	httpRouter.Methods("GET").Path("/configHtmlServer.json").HandlerFunc(ShowHtmlFile)

	// ---
	// HTML Files Show
	// ---

	//localhost:8080/admin/index.html
	httpRouter.HandleFunc("/admin", indexTemplateHandling)

	//localhost:8080/index.html
	httpRouter.HandleFunc("/", indexHomeTemplateHandling)

	// ---
	// HTML Files Show Todo
	// ---

	//localhost:8080/admin/create_todo.html
	httpRouter.HandleFunc("/admin/create_todo.html", createTodoTemplateHandling)

	//localhost:8080/admin/delete_all_todo.html
	httpRouter.HandleFunc("/admin/delete_all_todo.html", deleteAllTodosTemplateHandling)

	//localhost:8080/admin/delete_id_todo.html
	httpRouter.HandleFunc("/admin/delete_id_todo.html", deleteIdTodoTemplateHandling)

	//localhost:8080/admin/edit_id_todo.html
	httpRouter.HandleFunc("/admin/edit_id_todo.html", editIdTodoTemplateHandling)

	//localhost:8080/admin/show_all_todo.html
	httpRouter.HandleFunc("/admin/show_all_todo.html", showAllTodosTemplateHandling)

}

func RestStart(endpoint string) error {

	router := mux.NewRouter()

	restConfig(router)

	return http.ListenAndServe(endpoint, router)
}
