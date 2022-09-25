package util

func GetTopLastLenByCount(topCount, dataCount int) int {

	var count int

	switch {
	case dataCount > topCount:
		count = topCount
	case dataCount <= topCount:
		count = dataCount
	}

	return count
}
