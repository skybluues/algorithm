package algorithm

import "testing"

func TestBubbleSort(t *testing.T) {
	array := []int{9, 22, 4, 52, 75, 6, 735, 35, 77, 532, 75, 732, 15, 852}
	t.Log("source array: ", array)
	res := BubbleSort(array)
	t.Log("result array: ", res)
}

func TestInsertionSort(t *testing.T) {
	array := []int{9, 22, 4, 52, 75, 6, 735, 35, 77, 532, 75, 732, 15, 852}
	t.Log("source array: ", array)
	res := InsertionSort(array)
	t.Log("result array: ", res)
}
