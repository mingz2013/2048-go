package game_2048_go

import "github.com/mingz2013/game-2048-go/board"

func main() {
	actions := board.RandomActions()

	board.NewBoard().DoActionList(actions)
}
