package sort

import (
	"math/rand"
	"time"
)

//Random Sort.
// Create and compile a random program.
// Run the random program on the input array.
// If the program produces sorted output, you are done.
// Else repeat from the beginning.
func Random(arr ...int) []int {
	for {
		shuffle(arr...)
		if checkSorted(arr...) {
			break
		}
	}
	return arr
}

func shuffle(arr ...int) []int {
	rand.Seed(time.Now().UnixNano())
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func checkSorted(arr ...int) bool {
	var sorted = false
	var index = 0
	for ; index < len(arr)-1; index++ {
		if arr[index] < arr[index+1] {
			continue
		}
		break
	}

	if index == len(arr)-1 {
		sorted = true
	}
	return sorted
}
