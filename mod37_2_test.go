package iso7064

import "testing"

func TestNewMod372Calculator(t *testing.T) {
	calc := NewMod372Calculator()
	if calc.Modulus != 37 {
		t.Fatalf("Modulus is equal to %d instead of 37\n", calc.Modulus)
	}
	if calc.Radix != 2 {
		t.Fatalf("Radix is equal to %d instead of 2\n", calc.Radix)
	}
	if calc.IsDouble {
		t.Fatal("IsDouble is true, but should be false\n")
	}
	if calc.Charset != "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ*" {
		t.Fatalf("Charset is equal to %s instead of 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ*\n", calc.Charset)
	}
}
