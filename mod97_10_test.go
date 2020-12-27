package iso7064

import "testing"

func TestNewMod9710Calculator(t *testing.T) {
	calc := NewMod9710Calculator()
	if calc.Modulus != 97 {
		t.Fatalf("Modulus is equal to %d instead of 97\n", calc.Modulus)
	}
	if calc.Radix != 10 {
		t.Fatalf("Radix is equal to %d instead of 10\n", calc.Radix)
	}
	if !calc.IsDouble {
		t.Fatal("IsDouble is false, but should be true\n")
	}
	if calc.Charset != "0123456789" {
		t.Fatalf("Charset is equal to %s instead of 0123456789\n", calc.Charset)
	}
}
