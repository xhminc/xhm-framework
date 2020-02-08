package common

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"math/rand"
	"time"
)

const (
	RAND_NUM   = 0
	RAND_LOWER = 1
	RAND_UPPER = 2
	RAND_ALL   = 3
)

func MD5(text string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}

func SHA1(text string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(text)))
}

func SHA256(text string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(text)))
}

func SHA512(text string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(text)))
}

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
