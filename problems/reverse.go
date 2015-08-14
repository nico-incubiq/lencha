package problems

import (
	"clem/lencha/utils"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

var Reverse = Problem{
	Name:                "reverse",
	SolvingTime:         5 * time.Second,
	DurationBeforeRetry: 20 * time.Second,
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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func ReverseStartingHandler(state *ProblemState) (interface{}, error) {
	str := RandStringBytesMaskImprSrc(20)
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
