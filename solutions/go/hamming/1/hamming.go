package hamming



import "fmt"

// Distance calculates the Hamming distance between two DNA strands.
// Returns an error if the strands have different lengths.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("strands must be of equal length")
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}