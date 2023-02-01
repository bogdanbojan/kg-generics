package dupes

// Your Dupes function goes here!
func Dupes[T comparable](s []T) bool {
	seen := map[T]bool{}

	if len(s) < 2 {
		return false
	}

	for _, x := range s {
		if seen[x] {
			return true
		}
		seen[x] = true
	}

	return false
}
