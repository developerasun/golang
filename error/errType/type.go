package main

import (
	"fmt"
	"log"
)

func errorCheck(n int) (string, error) {
	if n == 1 {
		s := "Goorm"
		return s, nil
	}
	return "", fmt.Errorf("%d는 1이 아닙니다.", n)
}

func main() {
	s, err := errorCheck(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)

	s, err = errorCheck(2)
	if err != nil {
		// log.Print only prints error message.
		// Use this error in a less important issue.
		log.Print(err)
	}
	fmt.Println(s)

	defer func() {
		s, err = errorCheck(4)
		if err != nil {
			// log.Fatal completely finishes program.
			// Use this error in an important issue.
			log.Fatal(err)
		}
		fmt.Println(s)
	}()

	s, err = errorCheck(3)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(s)

	fmt.Println(s)

	fmt.Println("Hello, world!")
}

// Error importance order
// log.Print -> log.Panic -> log.Fatal
