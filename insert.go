package sort

//Insert .
//每次选择一个元素，并且将这个元素和整个数组中的所有元素进行比较，然后插入到合适的位置，
func Insert(arr ...int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr
}
