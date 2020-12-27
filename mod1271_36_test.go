package iso7064

import "testing"

func TestNewMod127136Calculator(t *testing.T) {
	calc := NewMod127136Calculator()
	if calc.Modulus != 1271 {
		t.Fatalf("Modulus is equal to %d instead of 1271\n", calc.Modulus)
	}
	if calc.Radix != 36 {
		t.Fatalf("Radix is equal to %d instead of 36\n", calc.Radix)
	}
	if !calc.IsDouble {
		t.Fatal("IsDouble is false, but should be true\n")
	}
	if calc.Charset != "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		t.Fatalf("Charset is equal to %s instead of 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ\n", calc.Charset)
	}
}
