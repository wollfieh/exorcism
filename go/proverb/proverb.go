package proverb

func Proverb(rhyme []string) []string {
	length := len(rhyme)
	res := make([]string, length)

	if length == 0 {
		return res
	}

	for i := 0; i < length-1; i++ {
		res[i] = "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
	}

	res[length-1] = "And all for the want of a " + rhyme[0] + "."
	return res
}
