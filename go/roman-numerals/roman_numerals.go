package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	var result string
	if input < 1 || input > 3999 {
		return "", errors.New("out of range")
	}
	result += getThousands(input)
	input = input % 1000
	result += getHundreds(input)
	input = input % 100
	result += getTens(input)
	input = input % 10
	result += getOnes(input)

	return result, nil
}
func getOnes(input int) string {
	switch input {
	case 9:
		return "IX"
	case 8:
		return "VIII"
	case 7:
		return "VII"
	case 6:
		return "VI"
	case 5:
		return "V"
	case 4:
		return "IV"
	case 3:
		return "III"
	case 2:
		return "II"
	case 1:
		return "I"
	default:
		return ""

	}
}
func getTens(input int) string {
	switch int(input / 10) {
	case 9:
		return "XC"
	case 8:
		return "XXC"
	case 7:
		return "LXX"
	case 6:
		return "LX"
	case 5:
		return "L"
	case 4:
		return "XL"
	case 3:
		return "XXX"
	case 2:
		return "XX"
	case 1:
		return "X"
	default:
		return ""
	}
}
func getHundreds(input int) string {
	switch int(input / 100) {
	case 9:
		return "CM"
	case 8:
		return "CCM"
	case 7:
		return "DCC"
	case 6:
		return "DC"
	case 5:
		return "D"
	case 4:
		return "CD"
	case 3:
		return "CCD"
	case 2:
		return "CC"
	case 1:
		return "C"
	default:
		return ""
	}

}

func getThousands(input int) string {
	switch input / 1000 {
	case 3:
		return "MMM"
	case 2:
		return "MM"
	case 1:
		return "M"
	default:
		return ""
	}
}
