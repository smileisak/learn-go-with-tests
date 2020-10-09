package numerals

import "strings"

// RomanNumeral a new type to hold Roman numeral
type RomanNumeral struct {
	Value  int
	Symbol string
}

// RomanNumerals is a type to hold a list of RomanNumeral Type
type RomanNumerals []RomanNumeral

// ValueOf return a value of a given symbol
func (r RomanNumerals) ValueOf(symbol string) int {
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

var allRomanNumerals = []RomanNumeral{
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

// ConvertToRoman converts an arabic number to a roman number
func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

//ConvertToArabic converts a Roman number to an Arabic number
func ConvertToArabic(roman string) int {
	total := 0
	return total
}
