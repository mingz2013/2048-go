package board

import (
	"log"
)

type Board [4][4]int

func NewBoard() *Board {
	b := &Board{}
	b.Init()
	return b
}

func (b *Board) Init() {

}

func (b *Board) DoActionList(actions string) {

	for _, v := range actions {
		switch v {
		case U:
			b.U()
		case D:
			b.D()
		case L:
			b.L()
		case R:
			b.R()
		default:
			log.Println("err action: ", v)
			break

		}
	}

}

func (b *Board) U() {

}

func (b *Board) D() {

}

func (b *Board) L() {

}

func (b *Board) R() {

}

func (b *Board) checkWin() {
	// 出现2048
}

func (b *Board) checkFail() {
	// 不能滑动
}
