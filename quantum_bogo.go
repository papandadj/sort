package sort

import (
	"math/rand"
	"time"
)

//QuantumBogo sort .
// Randomly shuffle the elements of the list.
// If the list is not sorted, destroy the universe.
// Any surviving universes will then have the sorted version of the list.
func QuantumBogo(arr ...int) []int {
	ch := make(chan []int, 0)
	rand.Seed(time.Now().UnixNano())
	//开启10000个宇宙
	for i := 0; i < 10000; i++ {
		arrInner := make([]int, len(arr))
		copy(arrInner, arr)
		go func(arrInner ...int) {
			shuffle(arrInner...)
			if checkSorted(arrInner...) {
				select {
				case ch <- arrInner:
				default:
				}
			}
		}(arrInner...)
	}
	arr = <-ch
	return arr
}
