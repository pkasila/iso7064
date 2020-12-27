package iso7064

import "testing"

func TestNewMod112Calculator(t *testing.T) {
	calc := NewMod112Calculator()
	if calc.Modulus != 11 {
		t.Fatalf("Modulus is equal to %d instead of 11\n", calc.Modulus)
	}
	if calc.Radix != 2 {
		t.Fatalf("Radix is equal to %d instead of 2\n", calc.Radix)
	}
	if calc.IsDouble {
		t.Fatal("IsDouble is true, but should be false\n")
	}
	if calc.Charset != "0123456789X" {
		t.Fatalf("Charset is equal to %s instead of 0123456789X\n", calc.Charset)
	}
}
