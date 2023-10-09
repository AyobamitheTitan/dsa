package sort

//Sorts the slice in place using the selection sort algorithm
func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		swap(&arr[i], &arr[min])
	}
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
