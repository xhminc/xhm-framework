package common

import (
	"math/rand"
	"time"
)

const (
	RAND_NUM   = 0
	RAND_LOWER = 1
	RAND_UPPER = 2
	RAND_ALL   = 3
)

func KeyRand(size int, randType int) []byte {

	kind, kinds, result := randType, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := randType > 2 || randType < 0
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		if isAll {
			kind = rand.Intn(3)
		}
		scope, base := kinds[kind][0], kinds[kind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}

	return result
}
