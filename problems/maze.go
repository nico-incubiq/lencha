package problems

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var Maze = Problem{
	Id:                  4,
	Name:                "maze",
	SolvingTime:         5 * time.Second,
	DurationBeforeRetry: DefaultTimeBeforeRetry,
	InProgressHandler:   MazeInProgressHandler,
	StartingHandler:     MazeStartingHandler,
}

type MazeData struct {
	Height, Width int
	Prof          int
	SolX, SolY    int
	Data          [][]byte
}

type MazeMesage struct {
	Height int      `json:"width"`
	Width  int      `json:"height"`
	Maze   []string `json:"maze"`
}

type MazeClientAnswer struct {
	Solution string
}

const (
	wall         = 'X'
	empty        = '-'
	start        = 'S'
	end          = 'E'
	visited      = 'v'
	up           = 'U'
	down         = 'D'
	left         = 'L'
	right        = 'R'
	failedLength = "The length of your answer is not correct."
	failedWall   = "Your answer goes trough a wall."
	failedOut    = "Your answer goes out of the maze."
	failedChar   = "Your answer contains other charater than U D L R."
	failedNotEnd = "Your answer does not finish on the End of the maze.."
)

func init() {
	gob.Register(MazeData{})
}

func MazeStartingHandler(state *ProblemState) (interface{}, error) {
	maze := NewMaze(50, 120)
	maze.GenExplo()
	state.Data = maze
	state.Status = StatusInProgress
	return maze.JSONStruct(), nil
}

func MazeInProgressHandler(r *http.Request, state *ProblemState) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	var answer MazeClientAnswer
	err := decoder.Decode(&answer)
	if err != nil {
		return "The Body of your request is not valid JSON or is not what we expect.", err
	}

	maze := state.Data.(MazeData)
	if mess, ok := maze.IsSolution(answer.Solution); ok {
		state.Status = StatusSuccess
		return mess, nil
	} else {
		state.Status = StatusFailed
		return mess, nil
	}
}

func NewMaze(width int, height int) *MazeData {
	var m MazeData

	m.Height = height
	m.Width = width
	m.Prof = 0

	m.Data = make([][]byte, width)

	for x := 0; x < width; x++ {
		m.Data[x] = make([]byte, height)
		for y := 0; y < height; y++ {
			m.Data[x][y] = wall
		}
	}

	return &m
}

func (m *MazeData) GenExplo() {
	m.Data[0][0] = empty
	m.genExploRec(0, 0, 0)
	m.Data[0][0] = start
	m.Data[m.SolX][m.SolY] = end
}

// We can open a case if there is only one case
// open, ie the previous one (to avoid creating loops)
func (m *MazeData) canOpen(x int, y int) bool {

	// If already opened dont open it again
	if m.Data[x][y] == empty {
		return false
	}

	countOpen := 0
	if y < m.Height-1 && m.Data[x][y+1] == empty {
		countOpen += 1
	}

	if y > 0 && m.Data[x][y-1] == empty {
		countOpen += 1
	}

	if x < m.Width-1 && m.Data[x+1][y] == empty {
		countOpen += 1
	}

	if x > 0 && m.Data[x-1][y] == empty {
		countOpen += 1
	}

	return countOpen == 1
}

func (m *MazeData) genExploRec(x int, y int, prof int) {

	if m.Prof < prof {
		m.Prof = prof
		m.SolX = x
		m.SolY = y
	}

	for _, dir := range rand.Perm(4) {
		switch dir {
		case 0: // up
			if y < m.Height-1 && m.canOpen(x, y+1) {
				m.Data[x][y+1] = empty
				m.genExploRec(x, y+1, prof+1)
			}
		case 1: // down
			if y > 0 && m.canOpen(x, y-1) {
				m.Data[x][y-1] = empty
				m.genExploRec(x, y-1, prof+1)
			}
		case 2: // left
			if x > 0 && m.canOpen(x-1, y) {
				m.Data[x-1][y] = empty
				m.genExploRec(x-1, y, prof+1)
			}
		case 3: // right
			if x < m.Width-1 && m.canOpen(x+1, y) {
				m.Data[x+1][y] = empty
				m.genExploRec(x+1, y, prof+1)
			}
		}
	}
}

func (m *MazeData) Display() {
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			fmt.Printf("%c", m.Data[x][y])
		}
		fmt.Println("")
	}
}

func (m *MazeData) Correct() bool {
	return m.Data[0][0] == empty && m.Data[m.Width-1][m.Height-1] == empty
}

func (m *MazeData) NumberWalls() int {
	res := 0
	for i := 0; i < m.Width; i++ {
		for j := 0; j < m.Height; j++ {
			if m.Data[i][j] == wall {
				res += 1
			}
		}
	}

	return res
}

func (m *MazeData) ToString() []string {
	res := make([]string, m.Width)
	for i := 0; i < m.Width; i++ {
		res[i] = string(m.Data[i])
	}
	return res
}

func (m *MazeData) JSONStruct() MazeMesage {
	mess := MazeMesage{
		Height: m.Height,
		Width:  m.Width,
		Maze:   m.ToString(),
	}

	return mess
}

func (m *MazeData) IsSolution(answer string) (string, bool) {
	x := 0
	y := 0

	// We already know the length of the answer
	if len(answer) != m.Prof {
		return failedLength, false
	}

	for _, c := range answer {
		// Update the position in the maze
		switch c {
		case up:
			x += -1
		case down:
			x += 1
		case left:
			y += -1
		case right:
			y += 1
		default:
			return failedChar, false
		}

		// Wrong if out of the maze
		if x < 0 || x > m.Width || y < 0 || y > m.Height {
			return failedOut, false
		}

		// Wrong if Wall
		if m.Data[x][y] == wall {
			return failedWall, false
		}
	}

	// Good if end (we should check outside of the for)
	// it will do for now
	if m.Data[x][y] == end {
		return StatusSuccessMessage, true
	} else {
		return failedNotEnd, false
	}
}
