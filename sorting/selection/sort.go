package selectsort

func Sort(arr []int) {
	curr := 0
	for curr < len(arr) {
		minInd := curr
		minVal := arr[minInd]
		for i := curr + 1; i < len(arr); i++ {
			if arr[i] < minVal {
				minVal = arr[i]
				minInd = i
			}
		}
		if minInd == curr {
			curr++
			continue
		}
		arr[curr], arr[minInd] = arr[minInd], arr[curr]
		curr++
	}
}
