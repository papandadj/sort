package sort

import "time"

//SolarBitflip .
// Check if the array is sorted.
// If the array is sorted, return the array.
// If the array is not sorted, wait for 10 seconds and pray for having bit flips caused by solar radiation, just in the correct order then repeat from step 1.
func SolarBitflip(arr ...int) (orders []int) {
	for {
		if checkSorted(arr...) {
			return
		}
		time.Sleep(time.Second * 10)
	}
}
