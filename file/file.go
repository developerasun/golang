package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	// FormFile returns the first file for the provided form key.
	uploadFile, header, err := r.FormFile("upload_file")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	dirName := "./uploads"

	// Refer to linux commands to understand what is permission bit: 0777
	os.MkdirAll(dirName, 0777)
	filepath := fmt.Sprintf("%s/%s", dirName, header.Filename)
	file, err := os.Create(filepath)

	// Check error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	// golang conventionally recommends to check if there is error before using defer.
	defer file.Close()

	// Copy copies from src to dst until either EOF is reached on src or an error occurs.
	// It returns the number of bytes copied and the first error encountered while copying, if any.
	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)

	// Print out where the uploaded file is located
	fmt.Fprint(w, filepath)

}

func main() {

	// HandleFunc registers the handler function for the given pattern
	// in the DefaultServeMux.
	// The documentation for ServeMux explains how patterns are matched.
	http.HandleFunc("/uploads", uploadsHandler)

	http.Handle("/", http.FileServer(http.Dir("public")))

	// ListenAndServe listens on the TCP network address addr
	// and then calls Serve with handler to handle requests on incoming connections.
	http.ListenAndServe(":3000", nil)
}
