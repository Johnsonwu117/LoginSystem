package tool

import (
	"math/rand"
	"time"
)

var Letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = Letters[r.Intn(62)]
	}

	return string(b)
}

func RandSeq1(n int) string {
	b := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = Letters[r.Intn(62)]
	}

	return string(b)
}
