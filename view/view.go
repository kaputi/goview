package view

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

type View struct {
	width      int
	height     int
	canvas     *Image
	drawBuf    *bytes.Buffer
	lastUpdate time.Time
	lastDraw   time.Time
}

func New(height, width int) View {
	return View{
		width:      width,
		height:     height,
		canvas:     NewImage(height, width),
		drawBuf:    new(bytes.Buffer),
		lastUpdate: time.Now(),
	}
}

func (v *View) Draw() {
	if v.lastUpdate != v.lastDraw {
		v.updateBuf()
		v.lastDraw = v.lastUpdate
	}

	// clean screen
	fmt.Print("\033[H\033[2J")

	fmt.Fprint(os.Stdout, v.drawBuf.String())

	v.lastDraw = v.lastUpdate
}

func (v *View) UpdateCellValue(value string, row, col int) {
	v.canvas.SetValue(value, row, col)
	v.lastUpdate = time.Now()
}

func (v *View) UpdateCellColor(color string, row, col int) {
	v.canvas.SetColor(color, row, col)
	v.lastUpdate = time.Now()
}

func (v *View) UpdateCell(value, color string, row, col int) {
	v.canvas.SetValueAndColor(value, color, row, col)
	v.lastUpdate = time.Now()
}

func (v *View) UpdateWithImage(img *Image, rowOffset, colOffset int) {
	// TODO: check for out of bounds

	// add to canvas
	for rowI := 0; rowI < img.Height; rowI++ {
		for colI := 0; colI < img.Width; colI++ {
			v.canvas.Cells[rowOffset+rowI][colOffset+colI] = img.Cells[rowI][colI]
		}
	}

	v.lastUpdate = time.Now()
}

func (v *View) ChangeImageColor(color string) {
	v.canvas.SetImageColor(color)
	v.lastUpdate = time.Now()
}

func (v *View) ClearCanvas() {
	for _, row := range v.canvas.Cells {
		for _, cell := range row {
			cell.value = ""
			cell.color = ""
		}
	}
	v.drawBuf.Reset()
}

func (v *View) updateBuf() {
  v.drawBuf.Reset()
	for _, row := range v.canvas.Cells {
		for colI, cell := range row {
			str := ""
			prevColor := ""
			nextColor := ""
			if colI != 0 {
				prevColor = row[colI-1].color
			}
			if colI < len(row)-1 {
				nextColor = row[colI+1].color
			}
			if cell.color != prevColor {
				str += GetColorCode(cell.color)
			}
			if cell.value == "" {
				str += " "
			} else {
				str += cell.value
			}
			if cell.color != nextColor {
				str += ResetColor
			}
			v.drawBuf.WriteString(str)
		}
		v.drawBuf.WriteString("\n")
	}
}
