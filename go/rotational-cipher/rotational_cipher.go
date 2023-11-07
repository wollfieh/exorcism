package rotationalcipher

// func RotationalCipher implements rot13 style substitution cypher
// by computing the offset from first letter
func RotationalCipher(plain string, shiftKey int) string {
	res := make([]byte, len(plain))

	for i, v := range []byte(plain) {
		switch {
		case v >= 'a' && v <= 'z':
			res[i] = 'a' + ((v-'a')+byte(shiftKey))%26
		case v >= 'A' && v <= 'Z':
			res[i] = 'A' + ((v-'A')+byte(shiftKey))%26
		default:
			res[i] = v

		}
	}
	return string(res)
}
