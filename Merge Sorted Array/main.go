package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
	}
	i, j := m-1, n-1

	for i >= 0 && j >= 0 {
		if nums2[j] > nums1[i] {
			nums1[m+n-1] = nums2[j]
			j--
			n--
		} else {
			nums1[m+n-1] = nums1[i]
			i--
			m--
		}
	}
	if j >= 0 {
		copy(nums1[:j+1], nums2[:j+1])
	}
}
