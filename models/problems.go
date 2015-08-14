package models

import "time"

type Problem struct {
	Id               int       `json:"id"`
	Title            string    `json:"title"`
	SmallDescription string    `json:"smallDescription" db:"small_description"`
	Description      string    `json:"description"`
	ApiUrl           string    `json:"apiUrl" db:"api_url"`
	SolvedTotal      int       `json:"solvedTotal" db:"solved_total"`
	CreatedAt        time.Time `json:"createdAt" db:"created_at"`
}

func GetAllProblems() ([]Problem, error) {
	problems := []Problem{}
	err := db.Select(&problems, "SELECT * FROM problems")

	return problems, err
}
