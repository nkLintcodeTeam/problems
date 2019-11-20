package search

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	low := 0
	high := len(nums) - 1
	// if the length of nums is 2
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}

		// should be <=, not <
		// what if low==mid? we need =
		if nums[low] <= nums[mid] {
			if target >= nums[low] && target <= nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
