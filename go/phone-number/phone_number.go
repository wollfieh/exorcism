package phonenumber

import (
	"fmt"
)

func Number(phoneNumber string) (string, error) {
	out := make([]rune, len(phoneNumber))
	pos := 0

	for _, v := range phoneNumber {
		if v <= '9' && v >= '0' {
			out[pos] = v
			pos++
		}
	}

	if pos == 11 && out[0] == '1' {
		// valid international
		out = out[1:]
		pos--
	}
	if pos == 10 && out[0] > '1' && out[3] > '1' {
		return string(out[:pos]), nil
	}
	return "", fmt.Errorf("not a phone number: %s", string(out[:pos]))
}

func AreaCode(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return phoneNumber, err
	}
	return num[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	// (613) 995-0253
	num, err := Number(phoneNumber)
	if err != nil {
		return phoneNumber, err
	}

	return fmt.Sprintf("(%s) %s-%s", num[0:3], num[3:6], num[6:]), nil
}
