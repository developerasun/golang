package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint: print out what io.Writer writes.
	fmt.Fprint(w, "Hello Go Web Server")

}

type fooHandler struct{}

// ServeHTTP : Handler(interface)'s method - spelling-sensitive.
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}

	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)

	// fmt.Fprintln: can print out calculation as well.
	num1 := 3
	num2 := 4
	fmt.Fprintln(w, num1+num2)

}

func NewHttpHandler() http.Handler {
	// The name mux stands for "HTTP request multiplexer".
	// Like the standard http.ServeMux, mux.Router matches incoming requests
	// against a list of registered routes and calls a handler for the route
	// that matches the URL or other conditions.
	mux := http.NewServeMux()

	// http.HandleFunc: Register handler function
	// "/": absolute path, the first path of domain
	mux.HandleFunc("/", indexHandler)

	// http.HandleFunc: Register handler function
	mux.HandleFunc("/bar", barHandler)

	// http.Handle take &fooHandler{} as argument since it embodies ServeHTTP method.
	mux.Handle("/foo", &fooHandler{})

	return mux
}
