package schedule

import (
	"time"

	"github.com/suhrr/garbageday-bot/weekday"
)

func ExtractTodayGarbagedays(garbagedays []Garbageday) []Garbageday {
	var todayList []Garbageday
	now := time.Now()

	for _, gagarbageday := range garbagedays {
		if gagarbageday.Weekday != int(now.Weekday()) {
			continue
		}
		// WeekdayInMonthの値が0以外なら、N回目の曜日かを判定
		if gagarbageday.WeekdayInMonth != 0 && weekday.CountWeekdayInMonth(now) != gagarbageday.WeekdayInMonth {
			continue
		}

		todayList = append(todayList, gagarbageday)
	}

	return todayList
}
