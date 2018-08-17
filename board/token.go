package board

import (
	"math/rand"
	"strconv"
	"time"
)

// 上下左右四个指令
const (
	U = 0
	D = 1
	L = 2
	R = 3
)

func RandomActions() (actions string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		j := r.Intn(4)
		actions += strconv.Itoa(j)
	}
	return
}
