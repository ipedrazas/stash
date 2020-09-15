package utils

// IntIndex returns the first index of the provided int value in the provided
// int slice. It returns -1 if it is not found.
func IntIndex(vs []int, t int) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// IntInclude returns true if the provided int value exists in the provided int
// slice.
func IntInclude(vs []int, t int) bool {
	return IntIndex(vs, t) >= 0
}

// IntAppendUnique appends toAdd to the vs int slice if toAdd does not already
// exist in the slice. It returns the new or unchanged int slice.
func IntAppendUnique(vs []int, toAdd int) []int {
	if IntInclude(vs, toAdd) {
		return vs
	}

	return append(vs, toAdd)
}

// IntAppendUniques appends a slice of int values to the vs int slice. It only
// appends values that do not already exist in the slice. It returns the new or
// unchanged int slice.
func IntAppendUniques(vs []int, toAdd []int) []int {
	for _, v := range toAdd {
		vs = IntAppendUnique(vs, v)
	}

	return vs
}
