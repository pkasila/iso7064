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
	}
}

func TestIBANCalculator_Verify(t *testing.T) {
	calc := NewIBANCalculator()
	for _, iban := range getIBANs() {
		valid, err := calc.Verify(iban.IBAN)
		if err != nil || valid != iban.Valid {
			t.Fatalf("IBAN %s: %s\n", iban.IBAN, err)
		}
	}
}

func TestIBANCalculator_Compute(t *testing.T) {
	calc := NewIBANCalculator()
	for _, iban := range getIBANs() {
		if !iban.Valid {
			// if invalid IBAN, then continue
			continue
		}
		computed, err := calc.Compute(iban.IBAN)
		if err != nil || computed != iban.IBAN {
			t.Fatalf("passed IBAN %s - Received IBAN %s, err: %s\n", iban.IBAN, computed, err)
		}
	}
}

func TestIBANCalculator_ComputeChars(t *testing.T) {
	calc := NewIBANCalculator()
	for _, iban := range getIBANs() {
		computed, err := calc.ComputeChars(iban.IBAN)
		if err != nil || computed != iban.CheckDigits {
			t.Fatalf("real digits %s - received digits %s, err: %s\n", iban.CheckDigits, computed, err)
		}
	}
}