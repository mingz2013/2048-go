package board

type Board [4][4]int

func NewBoard() *Board {
	b := &Board{}
	b.Init()
	return b
}

func (b *Board) Init() {

}

func (b *Board) DoActionList(actions string) {

}
