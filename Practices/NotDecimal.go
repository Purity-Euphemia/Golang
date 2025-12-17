package main

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	dotCount := 0

	for i, ch := range dec {
		if ch == '.' {
			dotCount++
			if dotCount > 1 {
				return dec + "\n"
			}
		} else if ch == '-' {
			if i != 0 {
				return dec + "\n"
			}
		} else if ch < '0' || ch > '9' {
			return dec + "\n"
		}
	}

	if dotCount == 0 {
		return dec + "\n"
	}

	dotIndex := -1
	for i := 0; i < len(dec); i++ {
		if dec[i] == '.' {
			dotIndex = i
			break
		}
	}

	allZero := true
	for i := dotIndex + 1; i < len(dec); i++ {
		if dec[i] != '0' {
			allZero = false
			break
		}
	}

	if allZero {
		return dec + "\n"
	}

	result := ""
	for i := 0; i < len(dec); i++ {
		if dec[i] != '.' {
			result += string(dec[i])
		}
	}

	return result + "\n"
}
