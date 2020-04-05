package sort

//Quick .
//快速排序（Quicksort）是对冒泡排序的一种改进
// 通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，
// 然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列
func Quick(arr ...int) []int {
	if len(arr) <= 1 {
		return arr
	}
	quick(arr, 0, len(arr)-1)
	return arr
}

func quick(arr []int, left, right int) {
	temp := arr[left]
	pointer := left

	i, j := left, right

	for i < j {
		for j >= pointer && arr[j] >= temp {
			j--
		}

		if j >= pointer {
			arr[pointer] = arr[j]
			pointer = j
		}

		for i <= pointer && arr[i] <= temp {
			i++
		}
		if i <= pointer {
			arr[pointer] = arr[i]
			pointer = i
		}
	}
	arr[pointer] = temp

	if pointer-left > 1 {
		quick(arr, left, pointer-1)
	}

	if right-pointer > 1 {
		quick(arr, pointer+1, right)
	}

}
