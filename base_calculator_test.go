package iso7064

import "testing"

type BCTestable struct {
	Value       string
	Valid       bool
	CheckDigits string
	Calc		string
}

func getCalc() map[string]*BaseCalculator {
	return map[string]*BaseCalculator {
		"MOD112": {
			Modulus:  11,
			Radix:    2,
			Charset:  "0123456789X",
			IsDouble: false,
		},
		"MOD9710": {
			Modulus:  97,
			Radix:    10,
			Charset:  "0123456789",
			IsDouble: true,
		},
	}
}

func getTestables() []BCTestable {
	return []BCTestable{
		{
			Value:       "12341234",
			Valid:       true,
			CheckDigits: "35",
			Calc:		 "MOD9710",
		},
		{
			Value:       "12341234",
			Valid:       false,
			CheckDigits: "00",
			Calc:		 "MOD9710",
		},
		{
			Value:       "12341234",
			Valid:       true,
			CheckDigits: "8",
			Calc:		 "MOD112",
		},
		{
			Value:       "12341234",
			Valid:       false,
			CheckDigits: "0",
			Calc:		 "MOD112",
		},
		{
			Value:       "ABCDEF",
			Valid:       false,
			CheckDigits: "0",
			Calc:		 "MOD112",
		},
		{
			Value:       "",
			Valid:       false,
			CheckDigits: "00",
			Calc:		 "MOD9710",
		},
	}
}

func TestBaseCalculator_Verify(t *testing.T) {
	calc := getCalc()
	for _, testable := range getTestables() {
		valid, err := calc[testable.Calc].Verify(testable.Value+testable.CheckDigits)

		if !testable.Valid {
			if valid != testable.Valid {
				t.Fatalf("%s: %s\n", testable.Value, err)
			}

			continue
		}

		if err != nil || valid != testable.Valid {
			t.Fatalf("%s: %s\n", testable.Value, err)
		}
	}
}

func TestBaseCalculator_Compute(t *testing.T) {
	calc := getCalc()
	for _, testable := range getTestables() {
		computed, err := calc[testable.Calc].Compute(testable.Value)
		if !testable.Valid {
			if computed == testable.Value+testable.CheckDigits {
				t.Fatalf("SHOULD BE INVALID: passed %s - received %s, err: %s\n", testable.Value, computed, err)
			}

			continue
		}
		if err != nil || computed != testable.Value+testable.CheckDigits {
			t.Fatalf("passed %s - received %s, err: %s\n", testable.Value, computed, err)
		}
	}
}

func TestBaseCalculator_ComputeChars(t *testing.T) {
	calc := getCalc()
	for _, testable := range getTestables() {
		computed, err := calc[testable.Calc].ComputeChars(testable.Value)

		if !testable.Valid {
			if computed == testable.CheckDigits {
				t.Fatalf("SHOULD BE INVALID: passed %s - received %s, err: %s\n", testable.Value, computed, err)
			}

			continue
		}

		if err != nil || computed != testable.CheckDigits {
			t.Fatalf("real digits %s - received digits %s, err: %s\n", testable.CheckDigits, computed, err)
		}
	}
}