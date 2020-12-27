package iso7064

// NewMod372Calculator creates ISO 7064 calculator with modulus equal to 37, radix equal to 2,
// charset "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ*" and single digit configuration
func NewMod372Calculator() *BaseCalculator {
	return &BaseCalculator{
		Modulus:  37,
		Radix:    2,
		Charset:  "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ*",
		IsDouble: false,
	}
}