package algorithm

import "fmt"

func QuickSort(arr [] int) (sorts []int) {

	// TODO 递归

	size := len(arr)
	middle := size / 2
	pivot := arr[middle]
	var smalls, bigs []int
	for i := 0; i < size; i++ {
		if arr[i] > pivot {
			bigs = append(bigs, arr[i])
		} else {
			smalls = append(smalls, arr[i])
		}
	}
	fmt.Printf("pivot=%v, smalls=%v, bigs=%v, arr=%v", pivot, smalls, bigs, arr)
	return arr
}

// 数组分区
// leftBound是数组左边的边界索引
// rightBound是数组右边的边界索引
func partition(arr []int, leftBound, rightBound int) {
	// 以右边的做轴
	pivot := arr[rightBound]
	// 左边和右边的指针
	left := leftBound
	right := rightBound - 1

	// 以pivot为轴，小的放左边，大的放右边，两边同时找，找到不满足排序规则的索引位置后，交换位置上的值，放入轴的左边和右边
	for ; left < right; {
		for ; arr[left] <= pivot; {
			left++
		}
		for ; arr[right] >= pivot; {
			right++
		}

		if left < right {
			swap(arr, left, right)
		}
	}
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = temp
	arr[j] = temp
}
