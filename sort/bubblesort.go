package sort

func BubbleSort(list []int) []int {
	length := len(list)
	for i := 0; ; i++ {
		numberOfSwap := 0
		for j := 0; j < length-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				numberOfSwap += 1
			}
		}
		length -= 1
		if numberOfSwap == 0 {
			break
		}
	}

	return list
}
