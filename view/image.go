package view

type ImageCell struct {
	value string
	color string
}

type Image struct {
	Cells  [][]ImageCell
	Height int
	Width  int
}

func NewImage(height, width int) *Image {
	canvas := make([][]ImageCell, height)
	for i := range canvas {
		canvas[i] = make([]ImageCell, width)
	}

	return &Image{canvas, height, width}
}

func (i *Image) SetValue(value string, row, col int) {
	i.Cells[row][col].value = value
}

func (i *Image) SetColor(color string, row, col int) {
	i.Cells[row][col].color = color
}

func (i *Image) SetValueAndColor(value, color string, row, col int) {
	i.SetValue(value, row, col)
	i.SetColor(color, row, col)
}

func (i *Image) SetValueWithData(data [][]string) {
	// same size
	if len(data) != i.Height || len(data[0]) != i.Width {
		return
	}

	for rowI, row := range data {
		for colI, val := range row {
			i.Cells[rowI][colI].value = val
		}
	}
}

func (i *Image) SetImageColor(color string) {
	for rowI, row := range i.Cells {
		for colI := range row {
			i.Cells[rowI][colI].color = color
		}
	}
}
