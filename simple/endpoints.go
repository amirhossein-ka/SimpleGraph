package simple

type Endpoint [][2]byte

func (e *Endpoint) UniquePairs() {
	// for this, first we order the bytes from smaller to larger (or normalize it),
	// then we use a map to check if that [2]byte is seen before or no.

	unique := map[[2]byte]struct{}{}
	normalize := [2]byte{}
	var result = Endpoint{}
	for _, b := range *e {
		if b[0] < b[1] {
			normalize = b
		} else {
			normalize[0], normalize[1] = b[1], b[0]
		}

		if _, exists := unique[normalize]; !exists {
			unique[normalize] = struct{}{}
			result = append(result, normalize)
		}
	}
	*e = result
}
