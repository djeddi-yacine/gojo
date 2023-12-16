package utils

func RemoveDuplicatesINT64(nums []int64) []int64 {
	// Create a map to track unique values
	seen := make(map[int64]struct{})

	// Create a new slice for the result
	result := make([]int64, 0, len(nums))

	// Iterate through the input slice
	for _, num := range nums {
		// Check if the value is unique
		if _, ok := seen[num]; !ok {
			// Add the value to the result slice
			result = append(result, num)

			// Mark the value as seen in the map
			seen[num] = struct{}{}
		}
	}

	return result
}
