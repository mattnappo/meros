package core

// SplitBytes splits a []byte n times
func SplitBytes(bytes []byte, n int) ([][]byte, error) {
	parts := make([][]byte, len(bytes))
	// var parts [][]byte
	var i int
	for _, b := range bytes {
		for i = 0; i < n; i++ {
			var split []byte
			append(split)
		}
	}
	return nil, nil // Placeholder
}
