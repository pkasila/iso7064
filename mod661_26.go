package iso7064

// NewMod66126Calculator creates ISO 7064 calculator with modulus equal to 661, radix equal to 26,
// charset "ABCDEFGHIJKLMNOPQRSTUVWXYZ" and double digit configuration
func NewMod66126Calculator() *BaseCalculator {
	return &BaseCalculator{
		Modulus:  661,
		Radix:    26,
		Charset:  "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		IsDouble: true,
	}
}