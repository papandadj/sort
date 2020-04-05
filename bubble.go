package sort

//Bubble .
func Bubble(arr ...int) []int {
	for i := len(arr); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
