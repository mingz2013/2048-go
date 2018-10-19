package board

import (
	"log"
	"math/rand"
	"time"
)

var (
	SIZE int
)

func Init() {
	SIZE = 4
}

//type Board [SIZE][SIZE]int
type Board [SIZE * SIZE]int

func NewBoard() *Board {
	b := &Board{}
	b.Init()
	return b
}

func (b *Board) Init() {
	// 随机出现两个2

	for i := 0; i < 2; i++ {
		b.randomOne()
	}

}

func (b *Board) randomOne() {

	// 随机放一个2或者4在空白的地方

	// 首先选择随机的是2还是4
	// 然后在所有空白的地方随机一个
	whitespace := b.getAllWhitespace()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(whitespace))

	b.setValue(index, 2)

}

func (b *Board) setValue(index int, value int) {
	b[index] = value
}

func (b *Board) getAllWhitespace() (lst []int) {

	// 获取一个所有空白的下标数组
	for index, value := range b {
		if value == 0 {
			lst = append(lst, index)
		}
	}
	return
}

func (b *Board) DoActionList(actions string) {
	// 上下左右滑动
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
