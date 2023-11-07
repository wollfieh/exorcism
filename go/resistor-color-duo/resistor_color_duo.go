package resistorcolorduo

var (
	ring map[string]int = map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}
)

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	return 10*ring[colors[0]] + ring[colors[1]]
}
