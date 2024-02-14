package main

func removeDuplicates(nums []int) int {
	slow := 0
	firstVal := true
	for i := 0; i < len(nums); i++ {

		if nums[i] != nums[slow] {
			slow++
			firstVal = true
			nums[slow] = nums[i]
		} else if firstVal && slow != i {
			slow++
			firstVal = false
			nums[slow] = nums[i]
		}

		// fmt.Printf("%v\n", nums)
	}

	return slow + 1
}
