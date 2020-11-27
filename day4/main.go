package main

import (
	"fmt"
	"github.com/elkoiko/advent2019/day4/password"
)

func main() {
	nbEligiblePasswords := password.GetDifferentPasswordsInRange(240298, 784956)
	fmt.Printf("%v password possibilities.\n", nbEligiblePasswords)
}
