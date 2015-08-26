package problems

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/claisne/lencha/utils"
)

var Equation = Problem{
	Id:                  2,
	Name:                "equation",
	SolvingTime:         3 * time.Second,
	DurationBeforeRetry: DefaultTimeBeforeRetry,
	InProgressHandler:   EquationInProgressHandler,
	StartingHandler:     EquationStartingHandler,
}

const (
	randRange       = 1000
	equationEpsilon = 0.0001
)

type EquationData struct {
	X1 int
	X2 int
}

type EquationClientAnswer struct {
	X1 float64 `json:"x1"`
	X2 float64 `json:"x2"`
}

func init() {
	gob.Register(EquationData{})
}

func EquationString(data EquationData) string {
	lambda := rand.Intn(2*randRange) - randRange
	a := -lambda * (data.X1 + data.X2)
	b := lambda * data.X1 * data.X2
	var buffer bytes.Buffer

	buffer.WriteString(strconv.Itoa(lambda))
	buffer.WriteString("*x^2")

	if a >= 0 {
		buffer.WriteString("+")
	}
	buffer.WriteString(strconv.Itoa(a))

	buffer.WriteString("*x")

	if b >= 0 {
		buffer.WriteString("+")
	}
	buffer.WriteString(strconv.Itoa(b))

	buffer.WriteString("=0")

	return buffer.String()
}

func EquationStartingHandler(state *ProblemState) (interface{}, error) {
	data := EquationData{
		X1: rand.Intn(2*randRange) - randRange,
		X2: rand.Intn(2*randRange) - randRange,
	}
	state.Data = data
	state.Status = StatusInProgress
	return EquationString(data), nil
}

func EquationInProgressHandler(r *http.Request, state *ProblemState) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	var answer EquationClientAnswer
	err := decoder.Decode(&answer)
	if err != nil {
		return "The Body of your request is not valid JSON or respecting the correct format.", err
	}

	equationData := state.Data.(EquationData)

	if utils.FloatEqualsWithPrecision(answer.X1, float64(equationData.X1), equationEpsilon) &&
		utils.FloatEqualsWithPrecision(answer.X2, float64(equationData.X2), equationEpsilon) {
		state.Status = StatusSuccess
	} else if utils.FloatEqualsWithPrecision(answer.X1, float64(equationData.X2), equationEpsilon) &&
		utils.FloatEqualsWithPrecision(answer.X2, float64(equationData.X1), equationEpsilon) {
		state.Status = StatusSuccess
	} else {
		state.Status = StatusFailed
	}

	return nil, nil
}
