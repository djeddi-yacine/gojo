package utils

func RemoveDuplicatesTitles(Titles []string) []string {
	seen := make(map[string]struct{})

	result := make([]string, 0, len(Titles))

	for _, v := range Titles {
		if _, ok := seen[v]; !ok {
			result = append(result, v)

			seen[v] = struct{}{}
		}
	}

	return result
}

func RemoveDuplicatesInt64(Nums []int64) []int64 {
	seen := make(map[int64]struct{})

	result := make([]int64, 0, len(Nums))

	for _, v := range Nums {
		if _, ok := seen[v]; !ok {
			result = append(result, v)

			seen[v] = struct{}{}
		}
	}

	return result
}
