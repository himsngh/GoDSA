package searching

/*
	Binary Search : Binary search is the search technique which works efficiently on the sorted arrays.
					Binary search follows divide and conquer approach in which, the list is divided into
					two halves and the item is compared with the middle element of the list. If the match
					is found then, the location of middle element is returned otherwise, we search into
					either of the halves depending upon the result produced through the match.
	Time Complexicity : O(log(n))

*/

func binarySearch(arr []int, target int) int {

	start := 0
	end := len(arr) - 1

	for start <= end {

		mid := start + (end-start)/2
		if arr[mid] == target {
			return mid

		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}

	return -1
}
