package sort

func InsertionSort(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]

			for j := i; j >= 0; j-- {
				if list[j] > list[j+1] {
					list[j], list[j+1] = list[j+1], list[j]
				}
			}
		}
	}
	return list
}
