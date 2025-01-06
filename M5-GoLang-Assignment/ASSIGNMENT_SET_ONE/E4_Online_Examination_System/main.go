package main

import (
	"examination/services"
	"fmt"
)

func main() {
	err := services.StartQuiz()
	if err != nil {
		fmt.Println("Failed to start because of error: ", err)
	} else {
		fmt.Println("Thank you for using Online Quiz!")
	}
}
