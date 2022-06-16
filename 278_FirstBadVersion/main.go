package firstbadversion

func firstBadVersion(n int) int {
	low := 1
	high := n
	for low < high {
		mid := low + (high-low)/2

		if /*isBadVersion(mid)*/ true {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
