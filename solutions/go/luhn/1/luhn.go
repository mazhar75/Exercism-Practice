package luhn

func Valid(id string) bool {
	var a []int

	// Remove spaces and validate digits
	for i := 0; i < len(id); i++ {
		if id[i] == ' ' {
			continue
		}
		if id[i] < '0' || id[i] > '9' {
			return false 
		}
		a = append(a, int(id[i]-'0'))
	}

	// Must have at least 2 digits
	if len(a) <= 1 {
		return false
	}

	// Double every second digit from the right
	for i := len(a) - 2; i >= 0; i -= 2 {
		a[i] *= 2
		if a[i] > 9 {
			a[i] -= 9
		}
	}

	// Sum and check mod 10
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}

	return sum%10 == 0
}