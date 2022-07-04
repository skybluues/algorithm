package algorithm

func BubbleSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	for i := 0; i < len(array)-1; i++ {
		for j := 1; j < len(array)-i; j++ {
			if array[j-1] > array[j] {
				tmp := array[j-1]
				array[j-1] = array[j]
				array[j] = tmp
			}
		}
	}
	return array
}

func InsertionSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	for i := 1; i < len(array); i++ {
		loc := -1
		for j := i; j > 0; j-- {
			if array[j-1] < array[i] {
				loc = j
				break
			}
		}
		if loc == -1 {
			loc = 0
		}
		tmp := array[i]
		for k := i; k > loc; k-- {
			array[k] = array[k-1]
		}
		array[loc] = tmp
	}
	return array
}
