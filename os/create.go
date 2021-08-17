package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// file : name + foramt
	var file_name string
	var file_format string = ".txt"

	// three ways of input : 1) user input 2) command line args 3) read external file
	fmt.Print("Enter your file name: ")
	fmt.Scanln(&file_name) // 1) user input

	CreateFile(file_name + file_format)
	OpenFile(file_name + file_format)
	ReadFile(file_name + file_format) // 3) read external file
	GetFileDir(file_name)

}

var myErr error = errors.New("this is my error handling!")

func CreateFile(file_name string) {
	// os.Create : parameter - name : string, returns : *os.File, error(*PathError)
	file, err := os.Create(file_name)

	// error handling
	if err != nil {
		log.Println(myErr)
	}
	// Close closes the File, rendering it unusable for I/O. Close will return an error if it has already been called.
	defer file.Close()

	file.WriteString("this is my txt file!")

}

// Read the contents of a file and display them on the terminal:
func OpenFile(file_name string) {
	file, err := os.Open(file_name)

	if err != nil {
		log.Print(myErr)
	}

	defer file.Close()

	// Stat returns the FileInfo structure describing file. If there is an error, it will be of type *PathError.
	stat, err := file.Stat()

	if err != nil {
		log.Print(myErr)
	}

	// Read reads up to len(b) bytes from the File.
	// It returns the number of bytes read and any error encountered. At end of file, Read returns 0, io.EOF.
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)

	if err != nil {
		log.Print(myErr)
	}

	str := string(bs) + "read from function OpenFile!"
	fmt.Println(str)

}

func ReadFile(file_name string) {
	// ReadFile reads the file named by filename and returns the contents. A successful call returns err == nil,
	// not err == EOF. Because ReadFile reads the whole file, it does not treat an EOF from Read as an error to be reported.
	//  Reading files is very common, so there’s a shorter way to do this using ioutil package
	bs, err := ioutil.ReadFile(file_name)

	if err != nil {
		log.Print(myErr)
	}

	// binary to string encoding
	str := string(bs) + "read from function ReadFile!"
	fmt.Println(str)

	return
}

func GetFileDir(file_name string) {
	dir, err := os.Open("./")

	if err != nil {
		log.Print(myErr)
	}

	defer dir.Close()

	// ReadDir reads the contents of the directory associated with the file f and
	// returns a slice of DirEntry values in directory order.
	// If n > 0, ReadDir returns at most n DirEntry records. In this case, if ReadDir returns an empty slice,
	// it will return an error explaining why. At the end of a directory, the error is io.EOF.
	// If n <= 0, ReadDir returns all the DirEntry records remaining in the directory.
	// When it succeeds, it returns a nil error (not io.EOF).
	fileInfos, err := dir.ReadDir(-1)

	if err != nil {
		log.Print(myErr)
	}

	fmt.Println(" ========================== In this directory, there are : ")
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
