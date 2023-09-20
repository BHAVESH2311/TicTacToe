package game
import (
	"tictactoe/board"
	"tictactoe/player"
	"fmt"
)

type Game struct {
	Board *board.Board
	Players [2]*player.Player
	turn uint
	isGameEnded bool
}

func NewGame(player0Name,player1Name string) *Game {
	var players [2]*player.Player
	players[0]=player.NewPlayer(player0Name,"X")
	players[1]=player.NewPlayer(player1Name,"O")
	return &Game{
		Board: board.NewBoard(),
		Players: players,
		turn: 0,
		isGameEnded: false,
	}
}

func (game *Game) HasGameEnded() bool {
	return game.isGameEnded
}

func (game *Game) Play(cellNumber uint) string  {
	if game.isGameEnded {
		return "The Game has Ended"
	}
	// validate cell number and is it empty?
	if !game.Board.ValidateCellNumber(cellNumber) {
		return "Invalid Cell Number"
	}
	if isEmpty,err:=game.Board.IsCellEmpty(cellNumber); err==nil && !isEmpty {
		return "The Selected Cell is not empty"
	}
	game.Board.MarkCell(cellNumber,game.Players[game.turn%2].GetChar())	
	game.turn++
	if game.IsWon(){
		game.isGameEnded = true
		return "Player: "+game.Players[(game.turn+1)%2].GetName()+" has won!!"
	}
	if game.IsDraw(){
		game.isGameEnded = true
		return "The game has drawn!!!"
	}
	return "The cell has been marked!!"
}

func (game *Game) GetPlayerBySymbol(symbol string) *player.Player{
	for i := 0; i < len(game.Players); i++ {
		if symbol == game.Players[i].GetChar(){
			return game.Players[i]
		}
	}
	return nil
}

func (game *Game) IsWon() (bool) {
	return game.Board.CheckDiagonals() || game.Board.CheckLinear()
}

func (game *Game) IsDraw() bool {
	return game.Board.IsDraw()
}

func (game *Game) DisplayBoard() {
	fmt.Println("-----------X and O board--------------")
	game.Board.DisplayBoard()
}