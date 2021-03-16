package main

func quick_arr(arr []int, l int, r int) int {
	i := l
	j := r
	num := arr[i]
	for i < j {
		 for i < j && arr[j] >= num {
			  j--
		 }
		 if i < j {
			  arr[i] = arr[j]
			  i++
		 }
		 for i < j && arr[i] < num {
			  i++
		 }
		 if i < j {
			  arr[j] = arr[i]
			  j--
		 }
	}
	arr[i] = num
	return i
}

func sort(arr []int, l int, r int) {
	if l < r {
		 i := quick_arr(arr[:], l, r)
		 sort(arr, l, i - 1)
		 sort(arr, i + 1, r)
	}
}