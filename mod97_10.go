package iso7064

// NewMod9710Calculator creates ISO 7064 calculator with modulus equal to 97, radix equal to 10, charset "0123456789"
// and double digit configuration
func NewMod9710Calculator() *BaseCalculator {
	return &BaseCalculator{
		Modulus:  97,
		Radix:    10,
		Charset:  "0123456789",
		IsDouble: true,
	}
}