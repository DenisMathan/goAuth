package router

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/denismathan/goAuth/src/entities"
	"github.com/denismathan/goAuth/src/router/interfaces"
	"github.com/gorilla/mux"
)

func todoRequests(router *mux.Router) {
	// router.HandleFunc("/api/createTodo", createTodo)
	router.HandleFunc("/todo", createTodo).Methods("POST")
	router.HandleFunc("/todos", readTodos).Methods("GET", "OPTIONS")
	router.HandleFunc("/todo", updateTodo).Methods("PUT", "OPTIONS")
	router.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE", "OPTIONS")
}
func readTodos(w http.ResponseWriter, request *http.Request) {
	sqlH := request.Context().Value("values").(map[string]interface{})["sqlHandler"].(interfaces.SqlHandler)
	userID := request.Context().Value("userID").(uint)
	var todos []entities.Todo
	sqlH.GetTodos(&todos, userID)
	json.NewEncoder(w).Encode(todos)
}

func createTodo(writer http.ResponseWriter, request *http.Request) {
	userID := request.Context().Value("userId").(uint)
	sqlH := request.Context().Value("sqlHandler").(interfaces.SqlHandler)
	var todo entities.Todo
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		io.WriteString(writer, err.Error())
		return
	}
	todo.OwnerID = userID
	sqlH.Create(&todo)
	var todos []entities.Todo
	sqlH.GetTodos(&todos, userID)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(todos)
}

func updateTodo(writer http.ResponseWriter, request *http.Request) {
	userID := request.Context().Value("userId").(uint)
	sqlH := request.Context().Value("sqlHandler").(interfaces.SqlHandler)
	var todo entities.Todo
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		io.WriteString(writer, err.Error())
		return
	}
	todo.OwnerID = userID
	sqlH.Update(&todo)
	var todos []entities.Todo
	sqlH.GetTodos(&todos, userID)
	json.NewEncoder(writer).Encode(todos)
}

func deleteTodo(writer http.ResponseWriter, request *http.Request) {
	userID := request.Context().Value("userId").(uint)
	sqlH := request.Context().Value("sqlHandler").(interfaces.SqlHandler)
	id := mux.Vars(request)["id"]
	sqlH.DeleteById(&entities.Todo{}, id)
	var todos []entities.Todo
	sqlH.GetTodos(&todos, userID)
	json.NewEncoder(writer).Encode(todos)
}
