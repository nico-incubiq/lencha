package models

import (
	"time"
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

func (u User) GetSolvedProblems() ([]Problem, error) {
	problems := []Problem{}
	err := db.Select(&problems, "SELECT pb.* FROM problemssolved solved_pb JOIN problems pb ON (pb.id = solved_pb.id) WHERE solved_pb.user_id=$1;", u.Id)

	return problems, err
}
