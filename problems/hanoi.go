package problems

import (
	"encoding/gob"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

var Hanoi = Problem{
	Id:                  6,
	Name:                "hanoi",
	SolvingTime:         5 * time.Second,
	DurationBeforeRetry: 5 * time.Second,
	InProgressHandler:   HanoiProgressHandler,
	StartingHandler:     HanoiStartingHandler,
}

type Move struct {
	Start int
	End   int
}

type HanoiClientAnswer struct {
	Moves []Move
}

const (
	failedIncorrectMove = "One of the moves is not authorized."
	failedUnfinished    = "All blocks have not reached the end."
	messageHanoiStart   = "Move all numbers from the first tower to the last one, in the same order."
	messageHanoiSuccess = "You successfully moved all numbers."
)

type HanoiData struct {
	Size   int
	Towers [3][]int
}

type HanoiMesage struct {
	Size    int      `json:"size"`
	Message string   `json:"message"`
	Towers  []string `json:"towers"`
}

func init() {
	gob.Register(HanoiData{})
}

func HanoiStartingHandler(state *ProblemState) (interface{}, error) {
	towers := NewHanoi(3)
	state.Data = towers
	state.Status = StatusInProgress
	return towers.JSONStruct(messageHanoiStart), nil
}

func NewHanoi(size int) *HanoiData {
	var m HanoiData

	m.Size = size

	// Create first tower.
	m.Towers[0] = make([]int, size)

	// Initialize first tower data.
	for y := 0; y < size; y++ {
		m.Towers[0][y] = size - y
	}

	return &m
}

func (m *HanoiData) ToString() []string {
	res := make([]string, m.Size)
	for i := 0; i < m.Size; i++ {
		res[i] = ""
		for j := 0; j < 3; j++ {
			if j > 0 {
				res[i] += ","
			}
			if len(m.Towers[j]) > (m.Size-1)-i {
				res[i] += strconv.Itoa(m.Towers[j][(m.Size-1)-i])
			} else {
				res[i] += " "
			}
		}
	}
	return res
}

func (m *HanoiData) JSONStruct(message string) HanoiMesage {
	mess := HanoiMesage{
		Message: message,
		Size:    m.Size,
		Towers:  m.ToString(),
	}

	return mess
}

func HanoiProgressHandler(r *http.Request, state *ProblemState) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	var answer HanoiClientAnswer
	err := decoder.Decode(&answer)
	if err != nil {
		return failedJSON, err
	}

	tower := state.Data.(HanoiData)

	if message, correct := tower.CheckMoves(answer.Moves); correct {
		state.Status = StatusSuccess
		return tower.JSONStruct(message), nil
	} else {
		state.Status = StatusFailed
		return tower.JSONStruct(message), nil
	}
}

func (m *HanoiData) CheckMoves(moves []Move) (string, bool) {
	for _, move := range moves {
		if !m.CheckMove(move) {
			return failedIncorrectMove, false
		}
	}

	if !m.IsFinished() {
		return failedUnfinished, false
	}

	return messageHanoiSuccess, true
}

func (m *HanoiData) CheckMove(move Move) bool {
	// Check move is between existing towers.
	if move.Start < 0 || move.Start >= 3 || move.End < 0 || move.End >= 3 {
		return false
	}

	// Check move takes a number on a non empty tower.
	if len(m.Towers[move.Start]) < 1 {
		return false
	}

	// Check that the destination tower either empty or its top number is smaller than the number being moved.
	if len(m.Towers[move.End]) > 0 && m.Towers[move.Start][len(m.Towers[move.Start])-1] > m.Towers[move.End][len(m.Towers[move.End])-1] {
		return false
	}

	// Actually move the number.
	m.Towers[move.End] = append(m.Towers[move.End], m.Towers[move.Start][len(m.Towers[move.Start])-1])
	m.Towers[move.Start] = m.Towers[move.Start][:len(m.Towers[move.Start])-1]
	return true
}

func (m *HanoiData) IsFinished() bool {
	return len(m.Towers[0]) == 0 && len(m.Towers[1]) == 0 && len(m.Towers[2]) == m.Size
}
