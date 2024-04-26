package utils

import (
	"math/rand"
	"time"
)

func RandSlice(n, max int) []int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rng.Intn(max)
	}

	return numbers
}

func Shuffle[T any](arr []T) {
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
}
