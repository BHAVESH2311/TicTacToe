package board

import (
	"tictactoe/cell"
	"errors"
	"fmt"
)

type Board struct {
	Cells [9]*cell.Cell
}

func NewBoard() *Board {
	var cells [9]*cell.Cell
	for i := 0; i < 9; i++ {
		cells[i] = cell.NewCell()
	}
	return &Board{
		Cells: cells,
	}
}

func (board *Board) IsCellEmpty(cellNumber uint) (bool,error) {
	if !board.ValidateCellNumber(cellNumber) {
		return true,errors.New("Cell number is invalid")
	}
	return board.Cells[cellNumber].IsEmpty(),nil
}

func (board *Board) ValidateCellNumber(cellNumber uint) bool{
	// length,_ := uint(len(board.Cells)-1)
	if cellNumber > uint(len(board.Cells)-1) {
		return false
	}
	return true
}

// func GetCell

func (board *Board) MarkCell(cellNumber uint,symbol string) {
	if !board.ValidateCellNumber(cellNumber) {
		fmt.Println(" Cell Number is invalid")
		return 
	}
	if isEmpty,_:=board.IsCellEmpty(cellNumber); !isEmpty{
		fmt.Println("Cell is not empty")
		return
	}
	board.Cells[cellNumber].MarkCell(symbol)
}

func (board *Board) CheckDiagonals() (bool) {
	if !board.Cells[0].IsEmpty() && board.Cells[0].IsEqual(board.Cells[4]) && board.Cells[0].IsEqual(board.Cells[8]) {
		return true
	}
	if !board.Cells[2].IsEmpty() && board.Cells[2].IsEqual(board.Cells[4]) && board.Cells[4].IsEqual(board.Cells[6]) {
		return true
	}
	return false
} 

func (board *Board) CheckLinear() (bool) {
	for i := 0; i < 3; i++ {
		if !board.Cells[i*3].IsEmpty() && board.Cells[i*3].IsEqual(board.Cells[i*3+1]) && board.Cells[i*3].IsEqual(board.Cells[i*3+2]) {
			return true
		}
		if !board.Cells[i].IsEmpty() && board.Cells[i].IsEqual(board.Cells[i+3]) && board.Cells[i].IsEqual(board.Cells[i+6]) {
			return true
		}
	}
	return false
}

func (board *Board) IsDraw() (bool) {
	for _, cell := range board.Cells {
		if cell.IsEmpty() {
			return false
		}
	}
	return true
}

func (board *Board) DisplayBoard() {
	for i := 0; i < 3; i++ {
		board.Cells[i*3].DisplayCell()
		board.Cells[i*3+1].DisplayCell()
		board.Cells[i*3+2].DisplayCell()
		fmt.Print("\n")
	}
}