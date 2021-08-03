package main

import (
	"net/http"

	"github.com/designerasun/golang/rest/myapp2"
)

func main() {
	http.ListenAndServe(":3001", myapp2.NewHandler())
}
