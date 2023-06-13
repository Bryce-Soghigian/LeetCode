package main

import "fmt"

/*
I do like binary search, but i also feel its like many parts of math,
where people do not understand why its important to learn it, but is very useful for other algorithms.


 Understanding divide and conquer is useful for many algorithms, and the pivot concept in the recursive solution is good to know.
*/

func binarySearch(arr []int, target int) int {
	/*
		Lets go through a couple of test cases here.
		We have a couple of pieces of state here.
		Start = 0
		End = len(arr) - 1
		Mid = (start + end) / 2

		arr = [1,2,3,4,5,6,7,8,9,10], target = 5
		          s 	 m         e    arr[s] = 0, arr[m] = 5, arr[e] = 10

				arr = [1,2,3,4,5,6,7,8,9,10], target = 3
						// iter 1
						s = 0
						e = 10
						m = 5
						// iter 2
						s = 0
						e = 5
						m = 2
						// iter 3
						s = 3
						e = 5
						m = 4
						// iter 4
						s = 3
						e = 4
						m = 3 -> found end here, but just for understanding the not found case continue
						// iter 5
						s = 3
						e = 3
						m = 3
			We see for the not found case, we end up with start and end being the same, and mid being the same. We end up converging if we do not find it

	*/

	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1 // not found
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(arr, 5))
	fmt.Println(binarySearch(arr, 3))
}
