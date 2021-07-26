package main

import (
	"net/http"

	"github.com/designerasun/golang/web/myapp"
)

func main() {

	// Wait port number and serve.
	http.ListenAndServe(":3000", myapp.NewHttpHandler())

}
