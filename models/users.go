package models

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	Id             int       `json:"id"`
	Username       string    `json:"username"`
	Hash           string    `json:"hash"`
	Email          string    `json:"email"`
	ApiKey         string    `json:"api_key" db:"api_key"`
	ProblemsSolved int       `json:"problems_solved" db:"problems_solved"`
	Privilege      int       `json:"privilege"`
	Activated      bool      `json:"activated"`
	EmailUpdate    bool      `json:"activated" db:"email_update"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

func GetAllUsers() ([]User, error) {
	users := []User{}
	err := db.Select(&users, "SELECT * FROM users")

	return users, err
}

func GetUserById(id int) (User, error) {
	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE id=$1", id)

	return user, err
}

func GetUserByApiKey(apiKey string) (User, error) {
	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE api_key=$1", apiKey)

	return user, err
}

func GetUserByUsername(username string) (User, error) {
	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE username=$1", username)

	return user, err
}

func CreateUser(user *User) error {
	_, err := db.NamedExec(`INSERT INTO users (username, hash, email, api_key, email_update) VALUES (:username,:hash,:email,:api_key,:email_update)`, *user)
	return err
}

func GetSolvedProblemsIdById(id int) ([]int, error) {
	var problemsId []int
	err := db.Select(&problemsId, "SELECT problem_id FROM problems_solved WHERE user_id = $1", id)

	return problemsId, err
}

func (user User) GetSolvedProblems() ([]Problem, error) {
	problems := []Problem{}
	err := db.Select(&problems, "SELECT pb.* FROM problems_solved solved_pb JOIN problems pb ON (pb.id = solved_pb.problem_id) WHERE solved_pb.user_id=$1;", user.Id)

	return problems, err
}

func (user User) SaveSuccessProblem(problemId int) error {
	_, err := db.Exec(`INSERT INTO problems_solved (user_id, problem_id) VALUES ($1, $2)`, user.Id, problemId)
	if err, ok := err.(*pq.Error); ok {
		// Already solved
		if err.Code.Class() == "23" {
			return nil
		} else {
			return err
		}
	}

	// Update Profile and problem stats
	_, err = db.Exec(`UPDATE users SET problems_solved = problems_solved + 1 WHERE id = $1`, user.Id)
	if err != nil {
		return err
	}

	_, err = db.Exec(`UPDATE problems SET solved_total = solved_total + 1 WHERE id = $1`, problemId)
	return err
}

func (user User) SetActivated() error {
	_, err := db.Exec(`UPDATE users SET activated = true WHERE id = $1`, user.Id)
	return err
}
