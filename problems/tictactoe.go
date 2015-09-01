package problems

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var TicTacToe = Problem{
	Id:                  1,
	Name:                "tictactoe",
	SolvingTime:         1000 * time.Second, //8 * time.Second,
	DurationBeforeRetry: 1000 * time.Second, //DefaultTimeBeforeRetry,
	InProgressHandler:   TicTacToeInProgressHandler,
	StartingHandler:     TicTacToeStartingHandler,
}

type TicTacToeData struct {
	Board [][]byte
}

type TicTacToeMessage struct {
	Board []string `json:"board"`
}

func (tm *TicTacToeMessage) ToTicTacToeData() (*TicTacToeData, error) {
	t := NewTicTacToe()
	errorLength := errors.New("Your Board is not correctly sized.")

	if len(tm.Board) != 3 {
		return nil, errorLength
	}

	for i, str := range tm.Board {
		if len(str) != 3 {
			return nil, errorLength
		}

		for j, c := range str {
			b := byte(c)
			if b == empty || b == player1Mark || b == player2Mark {
				t.Board[i][j] = b
			} else {
				return nil, errors.New("Wrong character!")
			}
		}
	}

	return t, nil
}

const (
	player1Mark         = 'X'
	player2Mark         = 'O'
	failedFullTTC       = "The game is over and you didn't win!"
	failedCharTTC       = "Wrong charater in you message (Only 'X', '-', 'O') are allowed."
	failedWrongStateTTC = "You sent a wrong state! You only have the right for a single move."
)

type Player byte

type TicTacToePos struct {
	x int
	y int
}

func init() {
	gob.Register(TicTacToeData{})
}

func TicTacToeStartingHandler(state *ProblemState) (interface{}, error) {
	t := NewTicTacToe()
	state.Data = t
	state.Status = StatusInProgress
	return t.JSONStruct(), nil
}

func TicTacToeInProgressHandler(r *http.Request, state *ProblemState) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	var answer TicTacToeMessage
	err := decoder.Decode(&answer)
	if err != nil {
		return "The Body of your request is not valid JSON or is not what we expect.", err
	}

	previousT := state.Data.(TicTacToeData)
	responseT, err := answer.ToTicTacToeData()
	if err != nil {
		state.Status = StatusFailed
		return failedCharTTC, nil
	}

	if responseT.IsPossibleNextBoard(previousT) {
		if win, _ := responseT.HasWinner(); win {
			state.Status = StatusSuccess
			return nil, nil
		} else if responseT.IsFull() {
			state.Status = StatusFailed
			return failedFullTTC, nil
		} else if responseT.IsFirstMove() {
			responseT.PlayRandom()
			state.Data = responseT
			return responseT.JSONStruct(), nil
		} else {
			// Ignore error we already checked if full
			responseT.Play()

			win, _ := responseT.HasWinner()
			if win || responseT.IsFull() {
				state.Status = StatusFailed
				return failedFullTTC, nil
			} else {
				state.Data = responseT
				return responseT.JSONStruct(), nil
			}
		}
	}

	return failedWrongStateTTC, nil
}

func NewTicTacToe() *TicTacToeData {
	var t TicTacToeData

	t.Board = make([][]byte, 3)

	for i := 0; i < 3; i++ {
		t.Board[i] = make([]byte, 3)
		for j := 0; j < 3; j++ {
			t.Board[i][j] = empty
		}
	}

	return &t
}

func (t *TicTacToeData) Display() {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Printf("%c", t.Board[x][y])
		}
		fmt.Println("")
	}
}

func (t *TicTacToeData) IsEmpty() bool {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if t.Board[x][y] != empty {
				return false
			}
		}
	}
	return true
}

func (t *TicTacToeData) IsFull() bool {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if t.Board[x][y] == empty {
				return false
			}
		}
	}
	return true
}

func (t *TicTacToeData) HasWinner() (bool, Player) {
	// Horizontal
	for x := 0; x < 3; x++ {
		if t.Board[x][0] == t.Board[x][1] && t.Board[x][1] == t.Board[x][2] &&
			t.Board[x][0] != empty {
			return true, Player(t.Board[x][0])
		}
	}

	// Vertical
	for y := 0; y < 3; y++ {
		if t.Board[0][y] == t.Board[1][y] && t.Board[1][y] == t.Board[2][y] &&
			t.Board[0][y] != empty {
			return true, Player(t.Board[0][y])
		}
	}

	// Diagonals
	if t.Board[0][0] == t.Board[1][1] && t.Board[1][1] == t.Board[2][2] &&
		t.Board[0][0] != empty {
		return true, Player(t.Board[0][0])
	}

	if t.Board[2][0] == t.Board[1][1] && t.Board[1][1] == t.Board[0][2] &&
		t.Board[2][0] != empty {
		return true, Player(t.Board[2][0])
	}

	return false, Player('-')
}

func (t *TicTacToeData) IsOver() bool {
	win, _ := t.HasWinner()
	return t.IsFull() || win
}

func (t *TicTacToeData) IsPossibleNextBoard(p TicTacToeData) bool {
	newMark := false
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if t.Board[x][y] != p.Board[x][y] {
				if newMark == false && t.Board[x][y] == 'O' {
					newMark = true
				} else {
					return false
				}
			}
		}
	}
	return newMark
}

func (t *TicTacToeData) PlayRandom() {

	search := true
	var x, y int

	for search {
		x = rand.Intn(3)
		y = rand.Intn(3)
		if t.Board[x][y] == empty {
			search = false
		}
	}

	t.Board[x][y] = player1Mark
}

func (t *TicTacToeData) score(depth int) int {
	if win, winner := t.HasWinner(); win {
		if winner == player1Mark {
			return 10 - depth
		} else {
			return depth - 10
		}
	} else {
		return 0
	}
}

func (t *TicTacToeData) EmptyPositions() []TicTacToePos {
	positions := make([]TicTacToePos, 0, 9)

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if t.Board[x][y] == empty {
				positions = append(positions, TicTacToePos{x: x, y: y})
			}
		}
	}

	return positions
}

func oppositeMark(mark byte) byte {
	if mark == player1Mark {
		return player2Mark
	} else {
		return player1Mark
	}
}

func (t *TicTacToeData) Play() (*TicTacToePos, error) {
	pos, err := t.negamaxWithPos(1)
	t.Board[pos.x][pos.y] = player1Mark
	return pos, err
}

// color: 1 player1 -1 player2
func (t *TicTacToeData) negamaxWithPos(color int) (*TicTacToePos, error) {
	if t.IsOver() {
		return nil, errors.New("Game already over.")
	}

	bestValue := -10
	bestPosition := TicTacToePos{x: -1, y: -1}

	for _, emptyPos := range t.EmptyPositions() {
		t.Board[emptyPos.x][emptyPos.y] = t.getMark(color)
		val := -t.negamax(-color, 0)
		if val >= bestValue {
			bestValue = val
			bestPosition = emptyPos
		}
		t.Board[emptyPos.x][emptyPos.y] = empty
	}

	return &bestPosition, nil
}

func (t *TicTacToeData) getMark(color int) byte {
	if color == 1 {
		return player1Mark
	} else {
		return player2Mark
	}
}

func (t *TicTacToeData) negamax(color int, depth int) int {
	if t.IsOver() {
		return color * t.score(depth)
	}

	bestValue := -10

	for _, emptyPos := range t.EmptyPositions() {
		t.Board[emptyPos.x][emptyPos.y] = t.getMark(color)
		val := -t.negamax(-color, depth+1)
		if val >= bestValue {
			bestValue = val
		}
		t.Board[emptyPos.x][emptyPos.y] = empty
	}

	return bestValue
}

func (t *TicTacToeData) NumberEmptyCases() int {
	res := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if t.Board[i][j] == empty {
				res += 1
			}
		}
	}
	return res
}

func (t *TicTacToeData) IsFirstMove() bool {
	return t.NumberEmptyCases() == 8
}

func (t *TicTacToeData) JSONStruct() (mess TicTacToeMessage) {
	mess.Board = make([]string, 3)
	for i := 0; i < 3; i++ {
		mess.Board[i] = string(t.Board[i])
	}

	return mess
}
