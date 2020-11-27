package password

import (
	"testing"
)

func TestPassMeetsCondition(t *testing.T) {
	var expected, actual bool
	var pass int
	var passRange [2]int

	// Must be 6-digit
	expected = false
	pass = 1234567
	passRange = [2]int{1234567, 1234568}
	actual = PassMeetsConditions(pass, passRange)
	if actual != expected {
		t.Error("Pass must be a 6-digit number. Expected invalid pass.")
	}

	expected = false
	pass = 123450
	passRange = [2]int{123456, 123456}
	actual = PassMeetsConditions(pass, passRange)
	if actual != expected {
		t.Errorf("Pass %v should not be in range %v. Expected invalid pass.", pass, passRange)
	}

	// Digits never decrease or stay the same
	expected = false
	pass = 122433
	passRange = [2]int{100000, 999999}
	actual = PassMeetsConditions(pass, passRange)
	if actual != expected {
		t.Errorf("Pass %v digits are decreasing. Expected invalid pass.", pass)
	}

	expected = true
	pass = 133679
	passRange = [2]int{100000, 999999}
	actual = PassMeetsConditions(pass, passRange)
	if actual != expected {
		t.Errorf("Pass %v digits are increasing. Expected valid pass.", pass)
	}

	// double 11 never decrease
	expected = true
	pass = 111111
	passRange = [2]int{100000, 999999}
	actual = PassMeetsConditions(pass, passRange)
	if actual != expected {
		t.Errorf("Pass %v digits never decrease. Expected valid pass.", pass)
	}

	expected = false
	pass = 223450
	passRange = [2]int{100000, 999999}
	actual = PassMeetsConditions(pass, passRange)
	if actual != expected {
		t.Errorf("Pass %v digits are decreasing. Expected invalid pass.", pass)
	}

	expected = false
	pass = 123789
	passRange = [2]int{100000, 999999}
	actual = PassMeetsConditions(pass, passRange)
	if actual != expected {
		t.Errorf("Pass %v does not have two adjacent digits being the same. Expected invalid pass.", pass)
	}
}
