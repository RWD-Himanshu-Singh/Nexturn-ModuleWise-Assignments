package models

type Question struct {
	ID           int
	QuestionText string
	Options      [4]string
	Answer       int
}
