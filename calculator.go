package iso7064

type Calculator interface {
	Verify(input string)		(bool, error)
	Compute(input string)		(string, error)
	ComputeChars(input string)	(string, error)
}