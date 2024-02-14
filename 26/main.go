package main

func removeDuplicates(nums []int) int {
	unique := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] != nums[unique] {
			unique++
			nums[unique] = nums[i]
			// fmt.Printf("%v\n", nums)
		}
	}

	return unique + 1
}
