package main

func median(arr []int) float64 {
	n := len(arr)
	if n%2 == 1 {
		return float64(arr[n/2])
	}
	return float64(arr[(n/2)-1]+arr[n/2]) / float64(2)
}

func mergeSort(arr1, arr2 []int) []int {
	var arr []int
	i1, i2 := 0, 0
	for i1 < len(arr1) && i2 < len(arr2) {
		if arr1[i1] < arr2[i2] {
			arr = append(arr, arr1[i1])
			i1++
		} else {
			arr = append(arr, arr2[i2])
			i2++
		}
	}
	if i1 < len(arr1) {
		arr = append(arr, arr1[i1:]...)
	}
	if i2 < len(arr2) {
		arr = append(arr, arr2[i2:]...)
	}
	return arr
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return median(mergeSort(nums1, nums2))
}
