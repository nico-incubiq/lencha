package problems

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/claisne/lencha/utils"
)

var Reverse = Problem{
	Name:                "reverse",
	SolvingTime:         3 * time.Second,
	DurationBeforeRetry: 5 * time.Second,
	Data:                ReverseData{},
	InProgressHandler:   ReverseInProgressHandler,
	StartingHandler:     ReverseStartingHandler,
}

type ReverseData struct {
	Str string `json:"string"`
}

type ReverseClientAnswer struct {
	Reversed string `json:"reversed"`
}

func ReverseStartingHandler(state *ProblemState) (interface{}, error) {
	str := utils.RandStringBytesMaskImprSrc(20)
	state.Data = ReverseData{Str: str}
	state.Status = StatusInProgress
	return str, nil
}

func ReverseInProgressHandler(r *http.Request, state *ProblemState) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	var answer ReverseClientAnswer
	err := decoder.Decode(&answer)
	if err != nil {
		return "The Body of your request is not valid JSON or is not what we expect.", err
	}

	reverseData := state.Data.(ReverseData)
	if answer.Reversed != utils.Reverse(reverseData.Str) {
		state.Status = StatusFailed
	} else {
		state.Status = StatusSucess
	}
	return nil, nil
}
