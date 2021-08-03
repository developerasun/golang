package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	Name string `json:"name"`

	// `json:"-"` will ignore a corresponding struct field. Information such as password should not be disclosed.
	Password string `json:"-"`

	// `json:" ,omitempty"` will omit a corresponding struct field.
	PrefferedFish []string  `json:"preferredFish,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
}

func main() {
	u := &User{
		Name:      "Sammy the Shark",
		Password:  "fish are great",
		CreatedAt: time.Now(),
	}

	result1, err := json.MarshalIndent(u, "", " ")

	if err != nil {
		log.Println(err)
		// Exit causes the current program to exit with the given status code.
		// Conventionally, code zero indicates success, non-zero an error.
		// The program terminates immediately; deferred functions are not run.
		os.Exit(2)
	}

	fmt.Println(string(result1))

}
