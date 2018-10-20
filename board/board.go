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

	b.setValueBase(index, 2)

}

func (b *Board) setValueBase(index int, value int) {
	b[index] = value
}

func (b *Board) setValue(x, y, v int) {
	b[x*SIZE+y] = v
}

func (b *Board) getValue(x, y int) int {
	return b[x*SIZE+y]
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

		if b.checkWin() {
			log.Println("win...")
			return
		}

		b.randomOne()

		if b.checkFail() {
			log.Println("fail...")
			return
		}

	}

}

func (b *Board) U() {

	for x := 0; x < SIZE; x++ {

		var newlst []int

		var last int

		for y := SIZE - 1; y >= 0; y-- {

			v := b.getValue(x, y)
			if v == 0 {
				continue
			}

			if last == 0 {
				last = v
				continue
			}

			if v == last {
				newlst = append(newlst, last*2)
				last = 0
				continue
			} else {
				newlst = append(newlst, last)
				last = v
				continue
			}

		}

		for y := SIZE - 1; y >= 0; y-- {
			if x < len(newlst) {
				b.setValue(x, y, newlst[y])
			} else {
				b.setValue(x, y, 0)
			}
		}

	}

}

func (b *Board) D() {

	for x := 0; x < SIZE; x++ {

		var newlst []int

		var last int

		for y := 0; y < SIZE; y++ {

			v := b.getValue(x, y)
			if v == 0 {
				continue
			}

			if last == 0 {
				last = v
				continue
			}

			if v == last {
				newlst = append(newlst, last*2)
				last = 0
				continue
			} else {
				newlst = append(newlst, last)
				last = v
				continue
			}

		}

		for y := 0; y < SIZE; y++ {
			if x < len(newlst) {
				b.setValue(x, y, newlst[y])
			} else {
				b.setValue(x, y, 0)
			}
		}

	}

}

func (b *Board) L() {
	// 往左滑动，每一行
	for y := 0; y < SIZE; y++ {

		var newlst []int

		var last int

		for x := 0; x < SIZE; x++ {

			v := b.getValue(x, y)
			if v == 0 {
				continue
			}

			if last == 0 {
				last = v
				continue
			}

			if v == last {
				newlst = append(newlst, last*2)
				last = 0
				continue
			} else {
				newlst = append(newlst, last)
				last = v
				continue
			}

		}

		for x := 0; x < SIZE; x++ {
			if x < len(newlst) {
				b.setValue(x, y, newlst[x])
			} else {
				b.setValue(x, y, 0)
			}
		}

	}

}

func (b *Board) R() {
	for y := 0; y < SIZE; y++ {

		var newlst []int

		var last int

		for x := SIZE - 1; x >= 0; x-- {

			v := b.getValue(x, y)
			if v == 0 {
				continue
			}

			if last == 0 {
				last = v
				continue
			}

			if v == last {
				newlst = append(newlst, last*2)
				last = 0
				continue
			} else {
				newlst = append(newlst, last)
				last = v
				continue
			}

		}

		for x := SIZE - 1; x >= 0; x-- {
			if x < len(newlst) {
				b.setValue(x, y, newlst[x])
			} else {
				b.setValue(x, y, 0)
			}
		}

	}
}

func (b *Board) getBestV() (best int) {
	for x := 0; x < SIZE; x++ {
		for y := 0; y < SIZE; y++ {
			v := b.getValue(x, y)
			if v > best {
				best = v
			}
		}
	}
	return
}

func (b *Board) checkWin() bool {
	// 出现2048
	bestV := b.getBestV()
	if bestV >= 2048 {
		return true
	}
	return false
}

func (b *Board) checkFail() bool {
	// 不能滑动
	// 所有格子都是满的，上下左右相邻的没有一样的

	for x := 0; x < SIZE; x++ {

		var last int

		for y := 0; y < SIZE; y++ {
			v := b.getValue(x, y)

			if v == 0 {
				return false
			}
			if last == 0 {
				last = v
				continue
			}

			if last == v {
				return false
			} else {
				last = v
				continue
			}
		}

	}

	for y := 0; y < SIZE; y++ {

		var last int

		for x := 0; x < SIZE; x++ {
			v := b.getValue(x, y)

			if v == 0 {
				return false
			}
			if last == 0 {
				last = v
				continue
			}

			if last == v {
				return false
			} else {
				last = v
				continue
			}
		}

	}

	return true
}
