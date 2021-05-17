/*
	Remove Element (Leetcode- 27) (https://leetcode.com/problems/remove-element/)
	Given an array nums and a value val, remove all instances of that value in-place and return the new length.
	Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
	The order of elements can be changed. It doesn't matter what you leave beyond the new length.
	
	Clarification:
	Confused why the returned value is an integer but your answer is an array?
	Note that the input array is passed in by reference, which means a modification to the input array will be known to the caller as well.
	Internally you can think of this:	

	// nums is passed in by reference. (i.e., without making a copy)
	int len = removeElement(nums, val);

	// any modification to nums in your function would be known by the caller.
	// using the length returned by your function, it prints the first len elements.
	for (int i = 0; i < len; i++) {
    	print(nums[i]);
	}	
*/

/*
		Dry Run := 
		array = [3, 2, 2, 3]  	, 	val = 3
		
	Solu :  Use two pointers to monitor the val element in the array . Whenever we find a position where the array value is 
	not equal to the val . Shift that value to the start of the array and increment the first pointer . 
	
	i=0, j=0;   array[j] = 3    array[j] == val; // [3, 2, 2, 3]

	i=0, j=1; 	arraa[j] = 2, val = 3; 	array[j] != val =>   array[i] = array[0] = array[j]; i++ => // [2, 2, 2, 3], i = 1
	i=1, j= 2;  array[j] = 2, val = 3;  array[j] != val =>   array[i] = array[1] = array[j]; i++ => // [2, 2, 2, 3], i = 2

	i=2, j=3;   array[j] = 3    array[j] == val; // [2, 2, 2, 3]

	therefore : the length of the required array is : 2 // i.e => [2, 2]

*/

package arrays	

func removeElement(array []int, val int) int {

	i := 0
	for j:=0; j<len(array); j++ {

		if array[j] != val {

			array[i] = array[j]
			i++
		}
	}

	return i
}

