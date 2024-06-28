package math

func Abs[I ~int | ~int8 | ~int64](x I) I {
	if x < 0 {
		return -x
	}
	return x
}
