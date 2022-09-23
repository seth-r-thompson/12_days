package main

import (
	"twelve_days/helpers/io"
	"twelve_days/xmas"
)

func main() {
	var err error

	gifts := io.Flag[[]string]{Name: "gifts", Value: nil}
	today := io.Flag[bool]{Name: "today", Description: "Start the countdown from today", Value: false}
	if err = io.Parse(&gifts, &today); err != nil {
		panic(err.Error())
	}

	days := xmas.Days
	if today.Value == true {
		days, err = xmas.DaysUntilChristmas()
		if err != nil {
			panic(err.Error())
		}
	}

	// elicit gifts if not given
	if len(gifts.Value) == 0 {
		gifts.Value, err = io.ScansWithPrompt("enter gift:", days)
		if err != nil {
			panic(err.Error())
		}
	}

	xmas.PrintVerses(days, gifts.Value)
}
