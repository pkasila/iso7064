package iso7064

import (
	"strconv"
	"strings"
)

type IBANCalculator struct {
	Calculator
	baseCalculator *BaseCalculator
}

func NewIBANCalculator() *IBANCalculator {
	return &IBANCalculator{
		baseCalculator: NewMod9710Calculator(),
	}
}

func (c *IBANCalculator) convertIBAN(input string) string {
	countryAndCheckChars := input[0:4]
	etc := input[4:]

	rearranged := etc + countryAndCheckChars

	converted := ""

	for _, char := range rearranged {
		i := int(char)

		if i > 64 && i < 91 {
			// A=10, B=11 etc...
			i -= 55
			// Add int as string to mod string
			converted += strconv.Itoa(i)
		} else {
			converted += string(char)
		}
	}

	return converted
}

func (c *IBANCalculator) Verify(input string) (bool, error) {
	return c.baseCalculator.Verify(c.convertIBAN(strings.ToUpper(input)))
}

func (c *IBANCalculator) Compute(input string) (string, error) {
	// Uppercase input
	upperCased := strings.ToUpper(input)

	cChars, err := c.ComputeChars(upperCased) // Without check digits
	if err != nil {
		return upperCased, err
	}

	countryCode := upperCased[0:2]
	etc := upperCased[4:]

	return countryCode + cChars + etc, nil
}

func (c *IBANCalculator) ComputeChars(input string) (string, error) {
	conv := c.convertIBAN(input)
	return c.baseCalculator.ComputeChars(conv[:len(conv)-2]) // Without place holder
}