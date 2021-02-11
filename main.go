package main

import "game-2048-go/board"

func main() {
	actions := board.RandomActions()

	board.NewBoard().DoActionList(actions)
}
