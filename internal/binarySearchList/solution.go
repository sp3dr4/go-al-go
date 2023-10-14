package binarysearchlist

func BinarySearchList(haystack []int, needle int) bool {
	lo := 0
	hi := len(haystack)

	for {
		m := lo + (hi-lo)/2
		v := haystack[m]

		if v == needle {
			return true
		} else if v > needle {
			hi = m
		} else {
			lo = m + 1
		}

		if !(lo < hi) {
			break
		}
	}
	return false
}
