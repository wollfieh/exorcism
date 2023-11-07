// Package twofer, exercism exercise
// https://exercism.org/tracks/go/exercises/two-fer
package twofer

// ShareWith returns a string constructed with $name or "you"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
