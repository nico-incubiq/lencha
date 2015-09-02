package problems

import (
	"encoding/gob"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

var TicTacToe = Problem{
	Id:                  5,
	Name:                "tic tac toe",
	SolvingTime:         1000 * time.Second,
	DurationBeforeRetry: 1000 * time.Second,
	InProgressHandler:   TicTacToeInProgressHandler,
	StartingHandler:     TicTacToeStartingHandler,
}

type TicTacToeClientAnswer struct {
	X int
	Y int
}

const (
	blank                 = '-'
	computerPon           = 'O'
	playerPon             = 'X'
	messagePlayerBegins   = "You begin."
	messagePlayerTurn     = "Your turn."
	messageComputerBegins = "I began, your turn."
	failedJSON            = "The Body of your request is not valid JSON or is not what we expect."
	failedUnavailable     = "Your answer overrides a position already taken."
	endLoose              = "You failed to beat the computer!"
	endWin                = "You won!"
	endTie                = "This game is a tie!"
)

type Coord struct {
	X int
	Y int
}

type TicTacToeData struct {
	Size          int
	ComputerBegan bool
	Board         [][]byte
}

type TicTacToeMesage struct {
	Size    int      `json:"size"`
	Message string   `json:"message"`
	Board   []string `json:"board"`
}

func init() {
	gob.Register(TicTacToeData{})
}

func TicTacToeStartingHandler(state *ProblemState) (interface{}, error) {
	board := NewTicTacToe(3)
	message := board.PickStart()
	state.Data = board
	state.Status = StatusInProgress
	return board.JSONStruct(message), nil
}

func NewTicTacToe(size int) *TicTacToeData {
	var m TicTacToeData

	m.Size = size
	m.Board = make([][]byte, size)

	for x := 0; x < size; x++ {
		m.Board[x] = make([]byte, size)
		for y := 0; y < size; y++ {
			m.Board[x][y] = blank
		}
	}

	return &m
}

func (m *TicTacToeData) PickStart() string {
	computerBegins := rand.Intn(2)
	if computerBegins == 1 {
		X := rand.Intn(m.Size)
		Y := rand.Intn(m.Size)

		m.Board[X][Y] = computerPon
		m.ComputerBegan = true

		return messageComputerBegins
	}

	m.ComputerBegan = false

	return messagePlayerBegins
}

func TicTacToeInProgressHandler(r *http.Request, state *ProblemState) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	var answer TicTacToeClientAnswer
	err := decoder.Decode(&answer)
	if err != nil {
		return failedJSON, err
	}

	board := state.Data.(TicTacToeData)

	if message, finished := board.DoPlayerMovePlusFightBack(answer.X, answer.Y); finished {
		switch message {
		case endWin, endTie:
			state.Status = StatusSuccess
			return board.JSONStruct(message), nil
		default:
			state.Status = StatusFailed
			return board.JSONStruct(message), nil
		}
	} else {
		state.Status = StatusInProgress
		return board.JSONStruct(message), nil
	}
}

func (m *TicTacToeData) ToString() []string {
	res := make([]string, m.Size)
	for i := 0; i < m.Size; i++ {
		res[i] = string(m.Board[i])
	}
	return res
}

func (m *TicTacToeData) JSONStruct(message string) TicTacToeMesage {
	mess := TicTacToeMesage{
		Message: message,
		Size:    m.Size,
		Board:   m.ToString(),
	}

	return mess
}

func (m *TicTacToeData) DoMove(moveX int, moveY int, pon byte) bool {
	if moveX < 0 || moveX >= m.Size || moveY < 0 || moveY >= m.Size || m.Board[moveX][moveY] != blank {
		return false
	}

	m.Board[moveX][moveY] = pon

	return true
}

func (m *TicTacToeData) DoPlayerMovePlusFightBack(moveX int, moveY int) (string, bool) {
	// Inverted because otherwise this is counterintuitive.
	if !m.DoMove(moveY, moveX, playerPon) {
		return failedUnavailable, true
	}

	if m.CheckWins(playerPon) {
		return endWin, true
	} else if _, availability := m.GetAvailableMoves(); !availability {
		return endTie, true
	}

	m.ComputerPlays()

	if m.CheckWins(computerPon) {
		return endLoose, true
	} else if _, availability := m.GetAvailableMoves(); !availability {
		return endTie, true
	}

	return messagePlayerTurn, false
}

func (m *TicTacToeData) CheckWins(pon byte) bool {
	// Check rows.
	for i := 0; i < m.Size; i++ {
		if m.WinsRow(i, pon) {
			return true
		}
	}
	// Check lines.
	for j := 0; j < m.Size; j++ {
		if m.WinsColumn(j, pon) {
			return true
		}
	}
	// Check diagonal 1.
	if m.WinsDiagonals(pon) {
		return true
	}

	return false
}

func (m *TicTacToeData) GetAvailableMoves() ([]Coord, bool) {
	var moves []Coord
	availability := false

	// Check possibility to move.
	for i := 0; i < m.Size; i++ {
		for j := 0; j < m.Size; j++ {
			if m.Board[i][j] == blank {
				c := Coord{X: i, Y: j}
				moves = append(moves, c)
				availability = true
			}
		}
	}

	return moves, availability
}

func (m *TicTacToeData) WinsRow(number int, pon byte) bool {
	for i := 0; i < m.Size; i++ {
		if m.Board[number][i] != pon {
			return false
		}
	}

	return true
}

func (m *TicTacToeData) WinsColumn(number int, pon byte) bool {
	for j := 0; j < m.Size; j++ {
		if m.Board[j][number] != pon {
			return false
		}
	}

	return true
}

func (m *TicTacToeData) WinsDiagonals(pon byte) bool {
	diagA := true
	diagB := true

	for k := 0; k < m.Size; k++ {
		if m.Board[k][k] != pon {
			diagA = false
		}
	}

	for k := 0; k < m.Size; k++ {
		if m.Board[(m.Size-1)-k][k] != pon {
			diagB = false
		}
	}

	return diagA || diagB
}

func (m *TicTacToeData) ComputerPlays() {
	_, bestMove := m.Minimax(0)

	m.DoMove(bestMove.X, bestMove.Y, computerPon)
}

func (m *TicTacToeData) GetScore(recursion int) int {
	max := m.Size*m.Size + 1

	if m.CheckWins(computerPon) {
		return max - recursion
	} else if m.CheckWins(playerPon) {
		return recursion - max
	} else {
		return 0
	}
}

func GetMaxValueIndex(a []int) int {
	index := 0
	for i, value := range a {
		if value > a[index] {
			index = i
		}
	}

	return index
}

func GetMinValueIndex(a []int) int {
	index := 0
	for i, value := range a {
		if value < a[index] {
			index = i
		}
	}

	return index
}

func (m *TicTacToeData) CleanMoves(moves []Coord) {
	for _, move := range moves {
		m.Board[move.X][move.Y] = blank
	}
}

func (m *TicTacToeData) Minimax(recursion int) (int, Coord) {
	availableMoves, availability := m.GetAvailableMoves()
	switch {
	case !availability, recursion%2 == 1 && m.CheckWins(computerPon), recursion%2 == 0 && m.CheckWins(playerPon):
		// No more move, computer won, or player won.
		return m.GetScore(recursion), Coord{-1, -1}
	}

	var scores []int
	var moves []Coord

	for _, coord := range availableMoves {
		if recursion%2 == 0 {
			m.DoMove(coord.X, coord.Y, computerPon)
		} else {
			m.DoMove(coord.X, coord.Y, playerPon)
		}

		score, _ := m.Minimax(recursion + 1)

		m.CleanMoves(availableMoves)

		scores = append(scores, score)
		moves = append(moves, coord)
	}

	if recursion%2 == 0 {
		// It is the computer turn.
		maxIndex := GetMaxValueIndex(scores)
		return scores[maxIndex], moves[maxIndex]
	} else {
		// It is the player turn.
		minIndex := GetMinValueIndex(scores)
		return scores[minIndex], moves[minIndex]
	}
}
