package main

// import fmt, net/http, github.com/gorilla/mux here ...
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Create main function to show "hello world" here ...
func main() {
	fmt.Println("Hello world!")

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world!"))
	}).Methods("GET")

	fmt.Println("Server is running on port 5000")
	http.ListenAndServe("localhost:5000", r)
}
