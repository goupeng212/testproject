package main

import (
	"fmt"
	//"time"
)

func bubbleSort(arr []int) {
	length := len(arr)
	var tmp int
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	fmt.Print(" ")
}
func cockTailSort(arr []int) {
	left := 0
	right := len(arr)
	var tmp int
	for left < right {
		for i := left; i < right; i++ {
			if arr[i] < arr[i+1] {
				tmp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
			}
		}

		right++
		for j := right; j > left; j-- {
			if arr[j-1] > arr[j] {
				tmp = arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = tmp
			}
		}
		left++
	}

}
func InsertionSort(arr []int) {
	length := len(arr)
	var preIndex, current int
	for i := 1; i < length; i++ {
		preIndex = i - 1
		current = arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = current
	}
}

//归并排序
//https://emacsist.github.io/2016/11/22/golang-%E5%BD%92%E5%B9%B6%E6%8E%92%E5%BA%8Fmergesort/

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	length := len(arr)
	middle := length / 2
	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])
	result := merge(left, right)
	return result
}
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

func QuickSort(arr []int, start, end int) {
	if start < end {
		pivot := quickpartition(arr, start, end)
		QuickSort(arr, start, pivot-1)
		QuickSort(arr, pivot+1, end)
	}
}
func quickpartition(arr []int, pivot, end int) int {
	//小于等于往左 大于往右，第一个的值作为轴值

	start := pivot //左边从开始的下一个启动
	pivotValue := arr[pivot]
	for start < end {
		//从右往左，找到小于等于轴值的位置
		for arr[end] > pivotValue && start < end {
			end--
		}
		//从最左侧的下一个从左往右找到大于轴值的位置
		for arr[start] <= pivotValue && start < end {
			start++
		}
		if start < end {
			arr[start], arr[end] = arr[end], arr[start]
		}

	}
	arr[start], arr[pivot] = arr[pivot], arr[start]
	return start
}

//func main() {
//	arr := []int{8, 20, 9, 3, 85, 4, 10, 7, 6, 15, 40}
//	fmt.Println(arr[:5])
//	fmt.Print(arr[5:])
//	result := MergeSort(arr)
//	fmt.Print("Merge sourt:", result)

//}
