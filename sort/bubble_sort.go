package main

import "fmt"

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// 给定的数组
	arr := []int{34, 56, 12, 23, 9, 88}
	
	fmt.Println("排序前:", arr)
	bubbleSort(arr)
	fmt.Println("排序后:", arr)
}