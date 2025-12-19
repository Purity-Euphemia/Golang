package piscine

func FromTo(from int, to int) string {
	if from < 0 || to < 0 || from > 99 || to > 99 {
		return "Invalid\n"
	}

	result := ""

	step := 1
	if from > to {
		step = -1
	}

	for i := from; ; i += step {
		if i < 10 {
			result += "0"
		}
		result += itoa(i)

		if i == to {
			break
		}
		result += ", "
	}

	return result + "\n"
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	str := ""
	for n > 0 {
		str = string('0'+n%10) + str
		n /= 10
	}
	return str
}
