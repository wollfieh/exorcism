package strand

func ToRNA(dna string) string {
	result := make([]rune, len(dna))
	complement := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}

	for i, r := range dna {
		result[i] = complement[r]
	}
	return string(result)

}
