package sort

func BubbleSort(list []int) []int {
	l := len(list)
	for i := 0; ; i++ {
		nswap := 0
		for j := 0; j < l-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				nswap += 1
			}
		}
		l -= 1
		if nswap == 0 {
			break
		}
	}

	return list
}
