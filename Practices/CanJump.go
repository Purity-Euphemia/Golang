package piscine

func CanJump(arr []uint) bool {
	if len(arr) == 0 {
		return false
	}

	if len(arr) == 1 {
		return true
	}

	pos := 0
	last := len(arr) - 1

	for pos < last {
		step := int(arr[pos])

		if step == 0 {
			return false
		}

		pos += step

		if pos > last {
			return false
		}
	}

	return pos == last
}
