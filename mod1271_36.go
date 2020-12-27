package iso7064

// NewMod127136Calculator creates ISO 7064 calculator with modulus equal to 1271, radix equal to 36,
// charset "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ" and double digit configuration
func NewMod127136Calculator() *BaseCalculator {
	return &BaseCalculator{
		Modulus:  1271,
		Radix:    36,
		Charset:  "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		IsDouble: true,
	}
}