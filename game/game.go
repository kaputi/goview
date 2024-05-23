package game

import (
	"fmt"
	"log"
	"time"

	"github.com/kaputi/sudokugo/gui"
	"github.com/kaputi/sudokugo/theme"
	"github.com/kaputi/sudokugo/view"
	"github.com/mattn/go-tty"
)

const (
	_ = iota
	PLAY
	NOTE1
	NOTE2
)

type Game struct {
	view        view.View
	theme       *theme.Theme
	gridImage   *view.Image
	layer       int
	lastLayer   int
	position    [2]int
	boardOffset [2]int
	testBoard   [9][9]int
}

func NewGame() Game {
	game := Game{
		view:      view.New(19, 40),
		theme:     theme.New(),
		layer:     PLAY,
		lastLayer: PLAY,
		position:  [2]int{0, 0},
	}

	gridImage := gui.GridImage("grid", game.theme)
	game.gridImage = gridImage

	// TEST:
	game.boardOffset = [2]int{0, 3}

	return game
}

func (g *Game) Start() {
	go g.listenForKeyPress()

	fmt.Print("\033[?25l") // hide cursor

	g.view.UpdateWithImage(g.gridImage, g.boardOffset[0], g.boardOffset[1])

	g.loop()
	fmt.Print("\033[?25h") // show cursor
}

func (g *Game) loop() {
	// for range time.Tick(16 * time.Millisecond) { // 60 fps
	for range time.Tick(33 * time.Millisecond) { // 30 fps
		g.update()
		g.render()
	}
}

func (g *Game) render() {
	g.view.Draw()
}

func (g *Game) update() {
	if g.layer != g.lastLayer {
		classColor := "grid"
		switch g.layer {
		case NOTE1:
			classColor = "note1"
		case NOTE2:
			classColor = "note2"
		}
		color := g.theme.GetClassColor(classColor)

		g.gridImage.SetImageColor(color)
		g.view.UpdateWithImage(g.gridImage, g.boardOffset[0], g.boardOffset[1])

		g.lastLayer = g.layer
	}

	g.drawNumbersAndPos()
}

func (g *Game) drawNumbersAndPos() {
	for rowI := 0; rowI < 9; rowI++ {
		for colI := 0; colI < 9; colI++ {
			viewRow, viewCol := gui.SudokuToViewCoord(rowI, colI, g.boardOffset)
			val := fmt.Sprintf("%d", g.testBoard[rowI][colI])
			if g.position[0] == rowI && g.position[1] == colI {
				g.view.UpdateCell(val, "black,white", viewRow, viewCol)
			} else {
				g.view.UpdateCell(val, "", viewRow, viewCol)
			}
		}
	}
}

func (g *Game) placeNumber(num int) {
	g.testBoard[g.position[0]][g.position[1]] = num
}

func (g *Game) changeLayer() {
	g.layer++
	if g.layer > NOTE2 {
		g.layer = PLAY
	}
}

func (g *Game) move(direction string) {
	switch direction {
	case "up":
		g.position[0]--
	case "down":
		g.position[0]++
	case "right":
		g.position[1]++
	case "left":
		g.position[1]--
	}

	g.position[0] = min(max(g.position[0], 0), 8)
	g.position[1] = min(max(g.position[1], 0), 8)
}

func (g *Game) listenForKeyPress() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		char, err := tty.ReadRune()
		if err != nil {
			panic(err)
		}

		// UP, DOWN, RIGHT, LEFT == [A, [B, [C, [D
		// we ignore the escape character [
		switch char {
		case 'A':
			g.move("up")
		case 'k':
			g.move("up")
		case 'B':
			g.move("down")
		case 'j':
			g.move("down")
		case 'C':
			g.move("right")
		case 'l':
			g.move("right")
		case 'D':
			g.move("left")
		case 'h':
			g.move("left")
		case 'n':
			g.changeLayer()
		case '0':
			g.placeNumber(0)
		case '1':
			g.placeNumber(1)
		case '2':
			g.placeNumber(2)
		case '3':
			g.placeNumber(3)
		case '4':
			g.placeNumber(4)
		case '5':
			g.placeNumber(5)
		case '6':
			g.placeNumber(6)
		case '7':
			g.placeNumber(7)
		case '8':
			g.placeNumber(8)
		case '9':
			g.placeNumber(9)
		}
	}
}
