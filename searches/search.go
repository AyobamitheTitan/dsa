package searches

func BinarySearch(arr []int, item int) (bool, int) {
	l, r := 0, len(arr)-1
	for l < r {
		mid := l + r/2
		if arr[mid] == item {
			return true, mid
		} else if arr[mid] > item {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false, -1
}
