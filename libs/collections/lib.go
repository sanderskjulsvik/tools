package collections

func Contains[T comparable](paths []T, path T) bool {
	for _, p := range paths {
		if p == path {
			return true
		}
	}
	return false
}
