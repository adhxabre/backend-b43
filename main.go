package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// {id: "1"}

// Declare Todos Struct here ...
type Todos struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

// Declare Todos Global Variable ...
var Todo = []Todos{
	{
		Id:     "1",
		Title:  "Cuci Tangan",
		IsDone: false,
	},
	{
		Id:     "2",
		Title:  "Jaga Jarak",
		IsDone: true,
	},
}

func main() {
	r := mux.NewRouter()

	// Create routes here ...
	r.HandleFunc("/todos", FindTodos).Methods("GET")
	r.HandleFunc("/todo/{id}", GetTodo).Methods("GET")
	r.HandleFunc("/todo", CreateTodo).Methods("POST")
	r.HandleFunc("/todo/{id}", UpdateTodo).Methods("PATCH")
	r.HandleFunc("/todo/{id}", DeleteTodo).Methods("DELETE")

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}

// Create FindTodos Function here ...
func FindTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Todo)
}

// Create GetTodo Function here ...
func GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var todoData Todos
	var isGetTodo = false

	for _, todo := range Todo {
		if id == todo.Id {
			isGetTodo = true
			todoData = todo
		}
	}

	if isGetTodo == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID: " + id + " not found!")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todoData)
}

// Create CreateTodo Function here ...
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data Todos

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Data isn't filled!")
		return
	}

	Todo = append(Todo, data)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Todo)
}

// Create UpdateTodo Function here ...
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var data Todos
	var isGetTodo = false

	json.NewDecoder(r.Body).Decode(&data)

	for idx, todo := range Todo {
		if id == todo.Id {
			isGetTodo = true
			Todo[idx] = data
		}
	}

	if isGetTodo == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID: " + id + " not found!")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Todo)
}

// Create DeleteTodo Function here ...
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var isGetTodo = false
	var index = 0

	for idx, todo := range Todo {
		if id == todo.Id {
			isGetTodo = true
			index = idx
		}
	}

	if isGetTodo == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID: " + id + " not found!")
		return
	}

	Todo = append(Todo[:index], Todo[index+1:]...)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("ID: " + id + " successfully deleted!")
}
