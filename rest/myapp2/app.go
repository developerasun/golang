package myapp2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var userMap map[int]*User
var lastID int

// User struct
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get UserInfo by /users/{id}")
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	user, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID: ", id)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
	// Vars returns the route variables for the current request, if any.
	// vars := mux.Vars(r)
	// fmt.Fprint(w, "User Id: ", vars["id"])
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// he new built-in function allocates memory. The first argument is a type, not a value,
	// and the value returned is a pointer to a newly allocated zero value of that type.
	user := new(User)

	// Decode reads the next JSON-encoded value from its input
	// and stores it in the value pointed to by v.
	err := json.NewDecoder(r.Body).Decode(user)

	// Check if there is error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// Created User
	lastID++
	user.ID = lastID
	user.CreatedAt = time.Now()

	userMap[user.ID] = user

	w.WriteHeader(http.StatusCreated)

	// Marshal returns the JSON encoding of v. Marshal traverses the value v recursively.
	// If an encountered value implements the Marshaler interface and is not a nil pointer,
	// Marshal calls its MarshalJSON method to produce JSON.
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))

}

func NewHandler() http.Handler {

	userMap = make(map[int]*User)
	lastID = 0
	// Use Gorilla mux instead of http.NewServeMux
	mux := mux.NewRouter()
	// mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

	// Handler should be differentiated by method even when it is about the same path
	// Methods adds a matcher for HTTP methods.
	// It accepts a sequence of one or more methods to be matched, e.g.: "GET", "POST", "PUT".
	mux.HandleFunc("/users", usersHandler).Methods("GET")
	mux.HandleFunc("/users", createUserHandler).Methods("POST")
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler)

	return mux
}
