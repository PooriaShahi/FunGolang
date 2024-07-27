package sort

func MergeSort(list []int) []int {
	if len(list)-1 <= 0 {
		return list
	}
	list1 := MergeSort(list[:len(list)/2])
	list2 := MergeSort(list[len(list)/2:])
	return conquer(list1, list2)
}

func conquer(list1, list2 []int) []int {
	result := make([]int, 0, len(list1)+len(list2))
	i, j := 0, 0
	for i < len(list1) && j < len(list2) {
		if list1[i] < list2[j] {
			result = append(result, list1[i])
			i++
		} else {
			result = append(result, list2[j])
			j++
		}
	}

	for i < len(list1) {
		result = append(result, list1[i])
		i++
	}

	for j < len(list2) {
		result = append(result, list2[j])
		j++
	}
	return result
}
