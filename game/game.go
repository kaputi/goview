package game

import (
	"fmt"
	"time"

	"github.com/kaputi/sudokugo/gui"
	"github.com/kaputi/sudokugo/theme"
	"github.com/kaputi/sudokugo/view"
)

type Game struct {
	view      view.View
	theme     *theme.Theme
	gridImage *view.Image
}

func NewGame() Game {
	game := Game{
		view:  view.New(19, 37),
		theme: theme.New(),
	}

	game.theme.UseAscii(true)

	gridImage := gui.GridImage("grid", game.theme)
	game.gridImage = gridImage

	return game
}

func (g *Game) Start() {
	fmt.Print("\033[?25l") // hide cursor

	// TODO: this is all for testing
	g.view.UpdateWithImage(g.gridImage, 0, 0)

	sudokuTest := [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 5, 0},
		{0, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 4, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 6},
		{0, 3, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 0},
	}
	for rowI, row := range sudokuTest {
		for colI, val := range row {
			if val != 0 {
				viewRow, viewCol := gui.SudokuToViewCoord(rowI, colI)
				g.view.UpdateCell(fmt.Sprint(val), "yellow", viewRow, viewCol)
			}
		}
	}

	g.loop()
	fmt.Print("\033[?25h") // show cursor
}

func (g *Game) loop() {
	for range time.Tick(16 * time.Millisecond) {
		// getInput
		// update
		g.render()
	}
}

func (g *Game) render() {
	g.view.Draw()
}
