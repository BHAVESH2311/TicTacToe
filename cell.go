package cell

import (
	"fmt"
)

type Cell struct {
	Mark string
}

func NewCell() *Cell {
	return &Cell{
		Mark:"N",
	}
}

func (cell *Cell) MarkCell(mark string){
	cell.Mark = mark
}

func (cell *Cell) IsEmpty() bool {
	return cell.Mark=="N"
}

func (cell *Cell) GetMark() string {
	return cell.Mark
}

func (cell1 *Cell) IsEqual(cell2 *Cell) bool {
	return cell1.Mark == cell2.Mark
}

func (cell *Cell) DisplayCell() {
	fmt.Print(cell.Mark+" ")
}
