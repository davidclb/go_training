package exos

// SafeAppend ajoute v à une COPIE de s, sans jamais modifier le backing array de s.
func SafeAppend(s []int, v int) []int {
	c := make([]int, len(s))
	copy(c, s)
	return append(c, v)

}
