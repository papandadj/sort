package sort

//Merge .
//归并排序（MERGE-SORT）是利用归并的思想实现的排序方法，
//该算法采用经典的分治（divide-and-conquer）策略（分治法将问题分(divide)成一些小的问题然后递归求解，
//而治(conquer)的阶段则将分的阶段得到的各答案"修补"在一起，即分而治之)。
func Merge(arr ...int) []int {
	divide(arr, 0, len(arr)-1)
	return arr
}

//治理数据
func conquer(arr []int, left, mid, right int) {
	var tempArr []int
	pl := left
	pr := mid + 1
	for pl <= mid && pr <= right {
		var item int
		if arr[pl] < arr[pr] {
			item = arr[pl]
			pl++
		} else {
			item = arr[pr]
			pr++
		}
		tempArr = append(tempArr, item)
	}

	for pl <= mid {
		tempArr = append(tempArr, arr[pl])
		pl++
	}

	for pr <= right {
		tempArr = append(tempArr, arr[pr])
		pr++
	}
	for i := range tempArr {
		arr[left+i] = tempArr[i]
	}
}

//分离数据
func divide(arr []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		divide(arr, left, mid)
		divide(arr, mid+1, right)
		conquer(arr, left, mid, right)
	}
}
