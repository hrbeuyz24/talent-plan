package main

import (
	"sort"
	"sync"
)

var (
	depth = 3
)

// merge merges two sorted arrays.
func merge(src []int64) {
	len := len(src)
	mid := len / 2
	tmp := make([]int64, len)
	leftIndex := 0
	rightIndex := mid
	num := 0
	for leftIndex < mid || rightIndex < len {
		if leftIndex < mid && rightIndex < len {
			if src[leftIndex] < src[rightIndex] {
				tmp[num] = src[leftIndex]
				leftIndex++
				num++
			} else {
				tmp[num] = src[rightIndex]
				rightIndex++
				num++
			}
		} else if leftIndex >= mid {
			tmp[num] = src[rightIndex]
			rightIndex++
			num++
		} else {
			tmp[num] = src[leftIndex]
			leftIndex++
			num++
		}
	}
	copy(src, tmp)
}

// mergeSort performs the multipole merge sort if deep > 0, or uses quick sort.
func mergeSort(src []int64, deep int) {
	if deep == 0 {
		sort.Slice(src, func(i, j int) bool {
			return src[i] < src[j]
		})
		return
	}

	len := len(src)
	mid := len / 2
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		mergeSort(src[:mid], deep-1)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		mergeSort(src[mid:], deep-1)
		wg.Done()
	}()
	wg.Wait()

	merge(src)
}

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	mergeSort(src, depth)
}
