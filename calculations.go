package main

import "C"

import (
	"math"
	"math/rand"
)

// Taken from Exercise 8
//export MeanStddev
func MeanStddev(arr []float64, chunks int) (mean, stddev float64) {
	if len(arr)%chunks != 0 {
		panic("You promised that chunks would divide slice size!")
	}
	// TODO: calculate the mean and population standard deviation of the array, breaking the array into chunks segments
	// and calculating on them in parallel.
	partialSum := make(chan PartialSum)
	size := len(arr) / chunks
	for i := 0; i < len(arr); i += size {
		temp := arr[i : i+size]
		go sigma(temp, partialSum)
	}

	var sum float64 = 0
	var ss float64 = 0

	for i := 0; i < len(arr); i += size {
		temp := <-partialSum
		sum += temp.sum
		ss += temp.sumSquare
	}

	n := float64(len(arr))
	mu := sum / n
	sd := math.Sqrt(ss/n - math.Pow(sum/n, 2))
	return mu, sd
}

func sigma(arr []float64, result chan PartialSum) {
	var s float64 = 0
	var sumSq float64 = 0
	for i := 0; i < len(arr); i++ {
		s += arr[i]
		sumSq += arr[i] * arr[i]
	}
	result <- PartialSum{s, sumSq}
}

type PartialSum struct {
	sum       float64
	sumSquare float64
}

// Partition the slice arr around a random pivot (in-place), and return the pivot location.
func partition(arr []float64) int {
	// Adapted from https://stackoverflow.com/a/15803401/6871666
	left := 0
	right := len(arr) - 1

	// Choose random pivot
	pivotIndex := rand.Intn(len(arr))

	// Stash pivot at the right of the slice
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Move elements smaller than the pivot to the left
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the last-smaller element
	arr[left], arr[right] = arr[right], arr[left]
	return left
}

func InsertionSort(arr []float64) {
	// TODO: implement insertion sort
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j = j - 1
		}
	}
}

const insertionSortCutoff = 50

//export QuickSort
func QuickSort(arr []float64) {
	// TODO: implement Quicksort:
	//   do nothing for length < 2
	//   do insertion sort for length < insertionSortCutoff
	//   do Quicksort otherwise.
	// TODO: decide on a good value for insertionSortCutoff
	if len(arr) < 2 {
		return
	}
	if len(arr) < insertionSortCutoff {
		InsertionSort(arr)
		return
	} else {
		pivot := partition(arr)
		if pivot > 0 {
			QuickSort(arr[:pivot-1])
		}
		if pivot+1 < len(arr) {
			QuickSort(arr[pivot+1:])
		}
	}
}

// ---------------------------------------------------------------------
// New code
//export Quartile
func Quartile(arr []float64) (float64, float64, float64) {
	if len(arr) < 1 {
		panic("The arry is of unappropiate size!")
	}
	QuickSort(arr)
	q1Index := 0.25*float64(len(arr)+1) - 1
	q2Index := 0.5*float64(len(arr)+1) - 1
	q3Index := 0.75*float64(len(arr)+1) - 1
	q1 := arr[uint64(q1Index)]
	q2 := arr[uint64(q2Index)]
	q3 := arr[uint64(q3Index)]

	if len(arr)%2 == 0 {
		if q1Index > math.Floor(q1Index) {
			q1 = (arr[uint(math.Floor(q1Index))] + arr[uint(math.Ceil(q1Index))]) / 2
		}

		if q2Index > math.Floor(q2Index) {
			q2 = (arr[uint(math.Floor(q2Index))] + arr[uint(math.Ceil(q2Index))]) / 2
		}

		if q3Index > math.Floor(q3Index) && int(q3Index+1) <= len(arr)-1 {
			q3 = (arr[uint(math.Floor(q3Index))] + arr[uint(math.Ceil(q3Index))]) / 2
		}
	}

	return q1, q2, q3
}

// Based on https://www.geeksforgeeks.org/frequent-element-array/
//export Mode
func Mode(arr []float64) (float64, int) {
	max_count := 1
	res := arr[0]
	curr_count := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			curr_count++
		} else {
			if curr_count > max_count {
				max_count = curr_count
				res = arr[i-1]
			}
			curr_count = 1
		}
	}

	if curr_count > max_count {
		max_count = curr_count
		res = arr[len(arr)-1]
	}

	return res, max_count
}

func main() {
}
