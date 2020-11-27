package password

// import "math"

// GetDifferentPasswordsInRange returns a slice of
// possible password between low and high range
func GetDifferentPasswordsInRange(low, high int) []int {
	var passwords []int

	for i := low; i <= high; i++ {
		if PassMeetsConditions(i, [2]int{low, high}) {
			passwords = append(passwords, i)
		}
	}
	return passwords
}

// PassMeetsConditions reports if pass meets conditions to be a correct password
// passRange contains low at 0 idex and high range at 1 index
func PassMeetsConditions(pass int, passRange [2]int) bool {
	var hasTwoSameAdjacentDigits bool

	// Must be a 6-digit number
	if pass < 100000 || pass > 999999 {
		return false
	}
	// Must be in range
	if pass < passRange[0] || pass > passRange[1] {
		return false
	}

	for i := 0; i < 5; i++ {
		rightDigit := pass % 10
		pass = pass / 10
		leftDigit := pass % 10
		if rightDigit < leftDigit { // Right digit must be higher than left digit
			return false
		} else if rightDigit == leftDigit {
			hasTwoSameAdjacentDigits = true
		}
	}

	// Must have two adjacent digits equal
	if hasTwoSameAdjacentDigits == false {
		return false
	}
	return true
}
