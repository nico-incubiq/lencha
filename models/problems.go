package models

import "time"

type Problem struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
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

func GetProblemByName(name string) (Problem, error) {
	problem := Problem{}
	err := db.Get(&problem, "SELECT * FROM problems WHERE api_url=$1", name)

	return problem, err
}
