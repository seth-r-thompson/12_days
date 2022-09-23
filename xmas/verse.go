package xmas

import (
	"fmt"
	"twelve_days/helpers/convert"
)

const giftStart = 1
const dayStart = 1

func PrintVerse(day int, gifts []string) {
	list := ""
	for gift := giftStart; gift <= day; gift++ {
		list += convert.NumToWord(1+day-gift) + " " + gifts[day-gift]
		if gift != day {
			list += ", "
		}
		if gift == day-1 {
			list += "and "
		}
	}
	fmt.Println("On the", convert.NumToOrdinal(day), "Day of Christmas, my true love gave to me:", list)
}

func PrintVerses(days int, gifts []string) {
	for day := dayStart; day <= days; day++ {
		if day > len(gifts) {
			fmt.Println("out of gifts!") // TODO: panic instead?
			break
		}
		PrintVerse(day, gifts)
	}
}
