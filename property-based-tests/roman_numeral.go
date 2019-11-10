package roman

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var RomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ArabicToRoman(arabic uint16) string {
	sb := strings.Builder{}

	for _, numeral := range RomanNumerals {
		for arabic >= numeral.Value {
			sb.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return sb.String()
}

func RomanToArabic(roman string) uint16 {
	var arabic uint16 = 0

	for _, numeral := range RomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			roman = roman[len(numeral.Symbol):]
			arabic += numeral.Value
		}
	}

	return arabic
}
