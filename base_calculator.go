package iso7064

import (
	"errors"
	"strings"
)

type BaseCalculator struct {
	Calculator
	Modulus		int
	Radix		int
	Charset		string
	IsDouble	bool
}

// Verify checks check characters are valid or not
func (c *BaseCalculator) Verify(input string) (bool, error) {
	// Uppercase input
	upperCased := strings.ToUpper(input)

	// Get digit number
	var digitNumber int
	if c.IsDouble {
		digitNumber = 2
	} else {
		digitNumber = 1
	}

	// If digitNumber is equal or bigger than input, then send false and error
	if len(upperCased) <= digitNumber {
		return false, errors.New("InputTooShort")
	}

	// Only computable data
	dataOnly := upperCased[0:len(upperCased)-digitNumber]

	computed, err := c.Compute(dataOnly)

	if err != nil {
		return false, err
	}

	return upperCased == computed, nil
}

// Compute computes check characters and returns input string with them
func (c *BaseCalculator) Compute(input string) (string, error) {
	// Uppercase input
	upperCased := strings.ToUpper(input)

	cChars, err := c.ComputeChars(upperCased)
	if err != nil {
		return upperCased, err
	}

	return upperCased + cChars, nil
}

// ComputeChars computes check characters and returns input string with them
func (c *BaseCalculator) ComputeChars(input string) (string, error) {
	// Process the string
	p := 0
	for _, char := range input {
		index := strings.IndexRune(c.Charset, char)
		if index < 0 {
			return input, errors.New("InvalidRune")
		}
		p = ((p + index) * c.Radix)%c.Modulus
	}

	if c.IsDouble {
		p = (p*c.Radix)%c.Modulus
	}

	checksum := (c.Modulus - p + 1)%c.Modulus

	if c.IsDouble {
		second := checksum % c.Radix
		first := (checksum - second) / c.Radix
		return c.Charset[first:first+1] + c.Charset[second:second+1], nil
	} else {
		return c.Charset[checksum:checksum+1], nil
	}
}