package main

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	newArray := []int{nums[0]}
	prev := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			newArray = append(newArray, nums[i])
		}
		prev = nums[i]
	}
	return len(newArray)
}

func main() {
	nums := []int{1, 1, 2}
	removeDuplicates(nums)
}
