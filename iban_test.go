package iso7064

import "testing"

type IBANTestable struct {
	IBAN        string // IBAN can be valid or invalid, determined by Valid
	Valid       bool
	CheckDigits string // CheckDigits are always valid
}

func getIBANs() []IBANTestable {
	return []IBANTestable{
		{
			IBAN:        "SE3550000000054910000003",
			Valid:       true,
			CheckDigits: "35",
		},
		{
			IBAN:        "CH9300762011623852957",
			Valid:       true,
			CheckDigits: "93",
		},
		{
			IBAN:        "DE89370400440532013000",
			Valid:       true,
			CheckDigits: "89",
		},
		{
			IBAN:        "IE29AIBK93115212345678",
			Valid:       true,
			CheckDigits: "29",
		},
		{
			IBAN:        "IE29AIBK9311521234567-",
			Valid:       false,
			CheckDigits: "00",
		},
	}
}

func TestIBANCalculator_Verify(t *testing.T) {
	calc := NewIBANCalculator()
	for _, iban := range getIBANs() {
		valid, err := calc.Verify(iban.IBAN)

		if !iban.Valid {
			if valid != iban.Valid {
				t.Fatalf("IBAN %s: %s\n", iban.IBAN, err)
			}
			continue
		}

		if err != nil || valid != iban.Valid {
			t.Fatalf("IBAN %s: %s\n", iban.IBAN, err)
		}
	}
}

func TestIBANCalculator_Compute(t *testing.T) {
	calc := NewIBANCalculator()
	for _, iban := range getIBANs() {
		computed, err := calc.Compute(iban.IBAN)
		if !iban.Valid {
			continue
		}
		if err != nil || computed != iban.IBAN {
			t.Fatalf("passed IBAN %s - Received IBAN %s, err: %s\n", iban.IBAN, computed, err)
		}
	}
}

func TestIBANCalculator_ComputeChars(t *testing.T) {
	calc := NewIBANCalculator()
	for _, iban := range getIBANs() {
		computed, err := calc.ComputeChars(iban.IBAN)

		if !iban.Valid {
			continue
		}

		if err != nil || computed != iban.CheckDigits {
			t.Fatalf("real digits %s - received digits %s, err: %s\n", iban.CheckDigits, computed, err)
		}
	}
}

func TestNewIBANCalculator(t *testing.T) {
	calc := NewIBANCalculator()
	if calc.baseCalculator.Modulus != 97 {
		t.Fatalf("Modulus is equal to %d instead of 97\n", calc.baseCalculator.Modulus)
	}
	if calc.baseCalculator.Radix != 10 {
		t.Fatalf("Radix is equal to %d instead of 10\n", calc.baseCalculator.Radix)
	}
	if !calc.baseCalculator.IsDouble {
		t.Fatal("IsDouble is false, but should be true\n")
	}
	if calc.baseCalculator.Charset != "0123456789" {
		t.Fatalf("Charset is equal to %s instead of 0123456789\n", calc.baseCalculator.Charset)
	}
}