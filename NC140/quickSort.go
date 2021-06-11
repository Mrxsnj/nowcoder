package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 将给定数组排序
 * @param arr int整型一维数组 待排序的数组
 * @return int整型一维数组
 */
func MySort(arr []int) []int {
	// write code here
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}

	i := low
	j := high
	index := arr[low]

	for i < j {
		for j > i && arr[j] >= index {
			j--
		}

		if i < j {
			arr[i] = arr[j]
			i++
		}

		for i < j && arr[i] <= index {
			i++
		}

		if i < j {
			arr[j] = arr[i]
			j--
		}
	}

	arr[i] = index

	quickSort(arr, low, i-1)
	quickSort(arr, i+1, high)
}
