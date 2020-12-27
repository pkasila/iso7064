package iso7064

import "testing"

func TestNewMod66126Calculator(t *testing.T) {
	calc := NewMod66126Calculator()
	if calc.Modulus != 661 {
		t.Fatalf("Modulus is equal to %d instead of 661\n", calc.Modulus)
	}
	if calc.Radix != 26 {
		t.Fatalf("Radix is equal to %d instead of 26\n", calc.Radix)
	}
	if !calc.IsDouble {
		t.Fatal("IsDouble is false, but should be true\n")
	}
	if calc.Charset != "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		t.Fatalf("Charset is equal to %s instead of ABCDEFGHIJKLMNOPQRSTUVWXYZ\n", calc.Charset)
	}
}
