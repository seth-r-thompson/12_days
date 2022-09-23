package convert

import "strconv"

func NumToOrdinal(num int) string {
	ordinal := strconv.Itoa(num)
	switch num {
	case 1:
		ordinal += "st"
	case 2:
		ordinal += "nd"
	case 3:
		ordinal += "rd"
	default:
		ordinal += "th"
	}
	return ordinal
}

func NumToWord(num int) string {
	switch num {
	case 1:
		return "a" // special case
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	default:
		panic("num < 1 or > 12")
	}
}
