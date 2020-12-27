package iso7064

// NewMod112Calculator creates ISO 7064 calculator with modulus equal to 11, radix equal to 2, charset "0123456789X"
// and single digit configuration
func NewMod112Calculator() *BaseCalculator {
	return &BaseCalculator{
		Modulus:  11,
		Radix:    2,
		Charset:  "0123456789X",
		IsDouble: false,
	}
}