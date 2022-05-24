package alphanum

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	ErrorInputMixed     = "character input has both letters and numbers"
	ErrorInputBad       = "character input was neither all letter(s) or all number(s)"
	ErrorModeUndeclared = "character mode undeclared in logic"
)

var alphanum = " ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Parse(s string) (result string, err error) {
	var (
		runes []rune
		mode  string
	)

	// Check input
	for _, r := range s {
		if unicode.IsLetter(r) {
			if mode == "" {
				runes = []rune{r}
				mode = "L"
			} else if mode == "L" {
				runes = append(runes, r)
			} else {
				err = fmt.Errorf(ErrorInputMixed)
				return "", err
			}
		} else if unicode.IsNumber(r) {
			if mode == "" {
				runes = []rune{r}
				mode = "N"
			} else if mode == "N" {
				runes = append(runes, r)
			} else {
				err = fmt.Errorf(ErrorInputMixed)
				return "", err
			}
		} else {
			err = fmt.Errorf(ErrorInputBad)
			return "", err
		}
	}

	// Handle Conversion
	S := strings.ToUpper(s)
	switch mode {
	case "L":
		// result = "All Letters"
		result = fmt.Sprintf("%d", ColumnLettersToNumbers(S))
	case "N":
		// result = "All Numbers"
		i, _ := strconv.Atoi(s)
		result = NumbersToColumnLetters(i)
	default:
		result = ""
		err = fmt.Errorf(ErrorModeUndeclared)
	}

	return result, err
}

// ColumnLettersToNumbers accepts lettered column names and returns their 1s index number value
// @see [Convert spreadsheet column letters (A,B, …, Z, AA, AB, …) to index numbers, and vice versa](https://gist.github.com/robinhouston/99746cac543e6b3ea61f1d245e9b19cc)
func ColumnLettersToNumbers(col string) int {
	x := 0
	for _, v := range col {
		i := int(v - 64)
		y := (26 * x) + i
		x = y
	}
	return x
}

// NumbersToColumnLetters accepts an 1s based index and returns a lettered column name
// @see [Convert spreadsheet column letters (A,B, …, Z, AA, AB, …) to index numbers, and vice versa](https://gist.github.com/robinhouston/99746cac543e6b3ea61f1d245e9b19cc)
func NumbersToColumnLetters(index int) (letters string) {
	var (
		alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i     int
		r     int
	)
	for index > 0 {
		i = index - 1
		r = i % 26
		index = i / 26
		letters = string(alpha[r]) + letters
	}
	return letters
}

func ParseSlice(s []string) (result []string) {
	for _, v := range s {
		v, err := Parse(v)
		if err != nil {
			result = append(result, "")
		} else {
			result = append(result, v)
		}
	}

	return result
}
