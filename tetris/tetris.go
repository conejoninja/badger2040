package tetris

import (
	"image/color"

	"math/rand"

	"tinygo.org/x/drivers"
	"tinygo.org/x/tinydraw"
)

type point [2]int16
type piece [4]point
type mypiece struct {
	p    piece
	x, y int16
}

var myp mypiece

const size = 16
const wb = 18 // 21 ~= 250/12
const hb = 8  // 10 ~= 122/12

var board [wb][hb]bool

var pieces = [7]piece{
	piece{
		point{1, 0},
		point{1, 1},
		point{1, 2},
		point{0, 0},
	},
	piece{
		point{1, 0},
		point{1, 1},
		point{1, 2},
		point{1, 3},
	},
	piece{
		point{1, 0},
		point{1, 1},
		point{0, 0},
		point{0, 1},
	},
	piece{
		point{1, 0},
		point{1, 1},
		point{1, 2},
		point{0, 1},
	},
	piece{
		point{0, 0},
		point{0, 1},
		point{1, 1},
		point{1, 2},
	},
	piece{
		point{1, 0},
		point{1, 1},
		point{0, 1},
		point{0, 2},
	},
	piece{
		point{1, 0},
		point{0, 1},
		point{0, 0},
		point{0, 2},
	},
}

var offset = int16(1)
var black = color.RGBA{1, 1, 1, 255}

func NewBoard() {
	for i := int16(0); i < wb; i++ {
		for j := int16(0); j < hb; j++ {
			board[i][j] = false
		}
	}
}

func DrawBoard(d drivers.Displayer) {
	for i := int16(0); i < wb; i++ {
		for j := int16(0); j < hb; j++ {
			if board[i][j] {
				tinydraw.FilledRectangle(d, 3+i*size, j*size, size, size, black)
			}
		}
	}
}

func NewPiece() {
	p := rand.Int31n(7) // random piece
	r := rand.Int31n(4) // random rotation
	for i := 0; i < 4; i++ {
		myp.p[i] = pieces[p][i]
	}
	for j := int32(0); j < r; j++ {
		pivot := myp.p[0]
		for i := 1; i < 4; i++ {
			dx := pivot[0] - myp.p[i][0]
			dy := pivot[1] - myp.p[i][1]
			myp.p[i][0] = pivot[0] - dy
			myp.p[i][1] = pivot[1] + dx
		}
	}
	myp.x = int16(rand.Int31n(21)) // random x location
	myp.y = 0
}

func DrawPiece(d drivers.Displayer) {
	tinydraw.FilledRectangle(d, 3+(myp.x+myp.p[0][0])*size, (myp.y+myp.p[0][1])*size, size, size, black)
	tinydraw.FilledRectangle(d, 3+(myp.x+myp.p[1][0])*size, (myp.y+myp.p[1][1])*size, size, size, black)
	tinydraw.FilledRectangle(d, 3+(myp.x+myp.p[2][0])*size, (myp.y+myp.p[2][1])*size, size, size, black)
	tinydraw.FilledRectangle(d, 3+(myp.x+myp.p[3][0])*size, (myp.y+myp.p[3][1])*size, size, size, black)
}

func MovePiece() bool {
	stopped := false
	for i := 0; i < 4; i++ {
		if myp.y+myp.p[i][1]+1 >= hb ||
			(myp.x+myp.p[i][0] >= 0 && myp.x+myp.p[i][0] < wb &&
				myp.y+myp.p[i][1] >= 0 && myp.y+myp.p[i][1]+1 < hb &&
				board[myp.x+myp.p[i][0]][myp.y+myp.p[i][1]+1]) {
			stopped = true
			break
		}
	}
	if stopped {
		for i := 0; i < 4; i++ {
			if myp.x+myp.p[i][0] >= 0 && myp.x+myp.p[i][0] < wb &&
				myp.y+myp.p[i][1] >= 0 && myp.y+myp.p[i][1] < hb {
				board[myp.x+myp.p[i][0]][myp.y+myp.p[i][1]] = true
			}
		}
		return false
	}
	myp.y++
	return true
}
