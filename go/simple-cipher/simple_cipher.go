package cipher

import "strings"

type shift struct {
	Key int
}

func (s shift) Encode(input string) string {
	var output strings.Builder
	output.Grow(len(input))

	for _, r := range input {

		switch {
		case r >= 'A' && r <= 'Z':
			r += 32 // make lowercase
			fallthrough
		case r >= 'a' && r <= 'z':
			offset := r - 'a'
			result := 'a' + (rune(s.Key)+offset+26)%26 // +26 because of negative key+offset
			output.WriteRune(result)
		}
	}

	return output.String()
}
func (s shift) Decode(input string) string {
	decoder := NewShift(s.Key * -1)
	return decoder.Encode(input)
}
func NewShift(distance int) Cipher {
	if (distance >= 1 && distance <= 25) ||
		(distance >= -25 && distance <= -1) {
		return shift{Key: distance}
	}
	return nil
}

// Caesar is a special case of shift with key=3
func NewCaesar() Cipher {
	return shift{Key: 3}
}

// Cipher noop returns as is. Used for the 'a' in vigenere keys
type noop struct{}

func (n noop) Encode(input string) string {
	return input
}
func (n noop) Decode(input string) string {
	return input
}
func NewNoop() Cipher {
	return noop{}
}

// Vigenere uses an array of different shift ciphers with
// shift key according to the vigenere key

type vigenere struct {
	Key     string
	ciphers []Cipher
}

func NewVigenere(key string) Cipher {
	var c vigenere
	var strength, shift int
	for _, v := range key {
		switch {
		case v == 'a':
			strength += 0
			c.ciphers = append(c.ciphers, NewNoop())
		case v >= 'b' && v <= 'z':
			shift = int(v - 'a')
			strength += shift
			c.ciphers = append(c.ciphers, NewShift(shift))
		default:
			// got not allowed character in key
			return nil
		}
	}
	if strength == 0 {
		// key is all 'a' s
		return nil
	}
	return c
}

func (v vigenere) Encode(input string) string {
	var result strings.Builder
	result.Grow(len(input))

	var o string
	var pos int

	for _, r := range input {
		switch {
		case r >= 'A' && r <= 'Z':
			r += 32 // make lowercase
			fallthrough
		case r >= 'a' && r <= 'z':
			o = v.ciphers[pos%len(v.ciphers)].Encode(string(r))
			result.WriteString(o)
			pos++
		}
	}
	return result.String()
}

func (v vigenere) Decode(input string) string {
	var result, o string
	var pos int

	for _, r := range input {
		o = v.ciphers[pos%len(v.ciphers)].Decode(string(r))
		result += o
		pos++
	}

	return result
}
