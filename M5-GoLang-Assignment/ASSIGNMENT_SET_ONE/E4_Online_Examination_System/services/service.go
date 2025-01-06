package services

import (
	"examination/models"
	"fmt"
	"strconv"
	"time"
)

var questions = []models.Question{
	{
		ID:           1,
		QuestionText: "Which country won the first Cricket World Cup in 1975?",
		Options:      [4]string{"India", "Australia", "West Indies", "England"},
		Answer:       3,
	},
	{
		ID:           2,
		QuestionText: "Who is known as the 'God of Cricket'?",
		Options:      [4]string{"Ricky Ponting", "Sachin Tendulkar", "Virat Kohli", "M.S. Dhoni"},
		Answer:       2,
	},
	{
		ID:           3,
		QuestionText: "What is the maximum number of overs allowed per bowler in an ODI match?",
		Options:      [4]string{"5", "10", "15", "20"},
		Answer:       2,
	},
	{
		ID:           4,
		QuestionText: "Which cricket ground is known as the 'Mecca of Cricket'?",
		Options:      [4]string{"Eden Gardens", "Lord's", "MCG", "Old Trafford"},
		Answer:       2,
	},
	{
		ID:           5,
		QuestionText: "Who holds the record for the highest individual score in Test cricket?",
		Options:      [4]string{"Brian Lara", "Don Bradman", "Virender Sehwag", "Chris Gayle"},
		Answer:       1,
	},
	{
		ID:           6,
		QuestionText: "Which Indian cricketer has taken the most wickets in Test cricket?",
		Options:      [4]string{"Anil Kumble", "Kapil Dev", "Harbhajan Singh", "Ravichandran Ashwin"},
		Answer:       1,
	},
	{
		ID:           7,
		QuestionText: "Who was the captain of the Indian cricket team during the 2011 World Cup victory?",
		Options:      [4]string{"Sourav Ganguly", "Rahul Dravid", "M.S. Dhoni", "Virat Kohli"},
		Answer:       3,
	},
	{
		ID:           8,
		QuestionText: "How many players are there in a cricket team on the field during a match?",
		Options:      [4]string{"9", "10", "11", "12"},
		Answer:       3,
	},
	{
		ID:           9,
		QuestionText: "Who holds the record for the fastest century in ODI cricket?",
		Options:      [4]string{"AB de Villiers", "Shahid Afridi", "Chris Gayle", "Virat Kohli"},
		Answer:       1,
	},
	{
		ID:           10,
		QuestionText: "Which bowler has taken the most wickets in T20 internationals?",
		Options:      [4]string{"Rashid Khan", "Lasith Malinga", "Shakib Al Hasan", "Tim Southee"},
		Answer:       4,
	},
}

func StartQuiz() error {
	fmt.Println("Welcome to the Online Examination System!")
	fmt.Println("Enter the option number (1-4) for each question.")
	fmt.Println("Type 'exit' anytime to quit the quiz.")
	var score int

	for _, question := range questions {
		fmt.Printf("Question %d: %s\n", question.ID, question.QuestionText)
		for _, option := range question.Options {
			fmt.Println(option)
		}

		answer := -1
		start := time.Now()

		for {
			// Prompt user for input
			fmt.Print("Your answer: ")
			var input string
			_, err := fmt.Scan(&input)
			if err != nil {
				fmt.Println("Error reading input. Try again.")
				continue
			}

			// Check for "exit" command
			if input == "exit" {
				fmt.Println("\nExiting the quiz. Thanks for playing!")
				return nil
			}

			// Convert input to integer
			answer, err = strconv.Atoi(input)
			if err != nil {
				fmt.Println("invalid input. please enter a valid answer")
				continue
			}

			// Validate answer range
			if answer < 1 || answer > 4 {
				fmt.Println("Invalid option. Please enter a number between 1 and 4.")
				continue
			}

			// Check time limit (10 seconds per question)
			if time.Since(start) > 10*time.Second {
				fmt.Println("\nTime's up for this question!")
				break
			}

			break
		}

		// Validate answer
		if answer == question.Answer {
			score++
		}
	}

	// Display final score and performance classification
	displayScore(score)
	return nil
}

func displayScore(score int) {
	fmt.Printf("\nQuiz Completed! Your Score: %d/%d\n", score, len(questions))

	switch {
	case score == len(questions):
		fmt.Println("Performance: Excellent!")
	case score >= len(questions)/2:
		fmt.Println("Performance: Good!")
	default:
		fmt.Println("Performance: Needs Improvement.")
	}
}
