/* There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5 */

package main

import (
	"fmt"
	"testing"
)

func Example() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	// Output:
	// 2.5
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	}
}
