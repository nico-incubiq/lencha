package problems

import (
	"bytes"
	"clem/lencha/models"
	"encoding/gob"
	"encoding/json"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	ctx "github.com/gorilla/context"
)

const (
	StatusStarting   = "starting"
	StatusInProgress = "in progress"
	StatusLate       = "toolate"
	StatusFailed     = "failed"
	StatusSucess     = "success"

	StatusSuccessMessage = "Challenge Done! Congratulations. We updated your profile."
	StatusLateMessage    = "Too late ! Try to be faster! Wait a few seconds before trying again."
	StatusFailedMessage  = "Challenge Failed! Wait a few seconds before trying again."
)

type Problem struct {
	Name                string        `json:"name"`
	SolvingTime         time.Duration `json:"solvingTime"`
	DurationBeforeRetry time.Duration `json:"secondsBeforeRetry"`
	Data                interface{}
	InProgressHandler   InProgressHandler
	StartingHandler     StartingHandler
}

type ProblemState struct {
	Status    string      `json:"status"`
	StartedAt time.Time   `json:"startedAt"`
	EndingAt  time.Time   `json:"endingAt"`
	Data      interface{} `json:"data"`
}

type ProblemResponse struct {
	Name      string      `json:"name"`
	Status    string      `json:"status"`
	StartedAt time.Time   `json:"startedAt"`
	EndingAt  time.Time   `json:"endingAt"`
	Message   interface{} `json:"message"`
}

type ProblemStateHandler func(r *http.Request, state *ProblemState) (interface{}, error)
type StartingHandler func(state *ProblemState) (interface{}, error)
type InProgressHandler func(r *http.Request, state *ProblemState) (interface{}, error)

func (problem Problem) GetState(user models.User) (*ProblemState, error) {
	redisC := models.RedisPool.Get()
	defer redisC.Close()

	stateBytes, err := redis.Bytes(redisC.Do("GET", problem.Name+":"+string(user.Id)))
	if err == redis.ErrNil {
		return &ProblemState{Status: StatusStarting, StartedAt: time.Now(), EndingAt: time.Now().Add(problem.SolvingTime)}, nil
	}
	if err != nil {
		return nil, err
	}

	var state ProblemState
	pCache := bytes.NewBuffer(stateBytes)
	decCache := gob.NewDecoder(pCache)
	gob.Register(ReverseData{})
	err = decCache.Decode(&state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

func (problem Problem) SetState(user models.User, state *ProblemState) error {
	redisC := models.RedisPool.Get()
	defer redisC.Close()

	mCache := new(bytes.Buffer)
	encCache := gob.NewEncoder(mCache)
	gob.Register(problem.Data)
	err := encCache.Encode(state)
	if err != nil {
		return err
	}

	_, err = redisC.Do("SET", problem.Name+":"+string(user.Id), mCache.Bytes())
	return err
}

func (problem Problem) DeleteState(user models.User) error {
	redisC := models.RedisPool.Get()
	defer redisC.Close()

	_, err := redisC.Do("DEL", problem.Name+":"+string(user.Id))
	return err
}

// Rate Limit By Ip
// Find Problem State
func HandlerFromStateHandler(problem Problem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the user and state of the problem
		user := ctx.Get(r, "user").(models.User)
		state, err := problem.GetState(user)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Warn("Problem Error")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		var message interface{}
		if state.Status == StatusStarting {
			message, err = problem.StartingHandler(state)
			go func(user models.User, problem Problem) {
				time.Sleep(problem.DurationBeforeRetry)
				problem.DeleteState(user)
			}(user, problem)
		} else if state.Status == StatusInProgress {
			if state.EndingAt.Before(time.Now()) {
				state.Status = StatusLate
			} else {
				message, err = problem.InProgressHandler(r, state)
				if err != nil {
					state.Status = StatusFailed
				}
			}
		}

		// Send message if problem is over
		switch state.Status {
		case StatusSucess:
			message = StatusSuccessMessage
		case StatusFailed:
			message = StatusFailedMessage
		case StatusLate:
			message = StatusLateMessage
		}

		err = problem.SetState(user, state)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := ProblemResponse{
			Name:      problem.Name,
			Status:    state.Status,
			StartedAt: state.StartedAt,
			EndingAt:  state.EndingAt,
			Message:   message,
		}
		JSONResponse(w, response)
	})
}

func JSONResponse(w http.ResponseWriter, response ProblemResponse) {
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Info("Error creating JSON response")
		panic(err)
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
