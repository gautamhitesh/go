package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomArray(size int) []int {
	result := make([]int, 0)
	rand.Seed(time.Now().Unix())
	for {
		if len(result) == size {
			break
		} else {
			temp := rand.Intn(size)
			if !getItem(result, len(result), temp) {
				result = append(result, temp)
			}
		}
	}
	return result
}
func getItem(array []int, size int, item int) bool {
	for i := 0; i < size; i++ {
		if array[i] == item {
			return true
		}
	}
	return false
}

func main() {
	t := GenerateRandomArray(10)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", t[i])
	}
}
