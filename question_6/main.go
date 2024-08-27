package main

import (
	"fmt"
	"math"
	"sort"
)

type MiddleElement struct {
	arr1   []int
	arr2   []int
	merged []int
}

func (mE *MiddleElement) new(arr1 []int, arr2 []int) {
	mE.arr1 = arr1
	mE.arr2 = arr2
	mE.merged = append(mE.merged, arr1...)
	mE.merged = append(mE.merged, arr2...)
	sort.Slice(mE.merged, func(i, j int) bool {
		return mE.merged[i] < mE.merged[j]
	})
}

func findMedianSortedArrays(nums1 []int, nums2 []int) int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	imin, imax := 0, m
	halfLen := (m + n + 1) / 2

	for imin <= imax {
		i := (imin + imax) / 2
		j := halfLen - i

		if i < m && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else {
			var maxOfLeft int
			if i == 0 {
				maxOfLeft = nums2[j-1]
			} else if j == 0 {
				maxOfLeft = nums1[i-1]
			} else {
				maxOfLeft = int(math.Max(float64(nums1[i-1]), float64(nums2[j-1])))
			}

			if (m+n)%2 == 1 {
				return maxOfLeft
			}

			var minOfRight int
			if i == m {
				minOfRight = nums2[j]
			} else if j == n {
				minOfRight = nums1[i]
			} else {
				minOfRight = int(math.Min(float64(nums1[i]), float64(nums2[j])))
			}

			return (maxOfLeft + minOfRight) / 2
		}
	}

	return -1
}

func main() {
	arr1 := []int{1, 2, 5, 7, 8}
	arr2 := []int{3, 4, 6}

	mE := new(MiddleElement)
	mE.new(arr1, arr2)

	middle := findMedianSortedArrays(mE.arr1, mE.arr2)
	fmt.Println(middle)
}
