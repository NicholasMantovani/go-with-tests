package propertybased

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{90, "XC"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q want %q", got, test.Roman)
			}

		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d want %d", got, test.Arabic)
			}

		})
	}
}

func TestConvertingToArabicRecursive(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabicRecursive(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d want %d", got, test.Arabic)
			}

		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic) //This will not be printed if you only run go test it need the -v arg
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return arabic == fromRoman
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}
}

func BenchmarkConvertToArabicRecursive(b *testing.B) {
	testSet := []string{}
	for i := 0; i < b.N; i++ {
		n := i
		if n > 3999 {
			n = n % 4000
		}
		testSet = append(testSet, ConvertToRoman(uint16(n)))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConvertToArabicRecursive(testSet[i])
	}
}

func BenchmarkConvertToArabic(b *testing.B) {
	testSet := []string{}
	for i := 0; i < b.N; i++ {
		n := i
		if n > 3999 {
			n = n % 4000
		}
		testSet = append(testSet, ConvertToRoman(uint16(n)))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConvertToArabic(testSet[i])
	}
}
