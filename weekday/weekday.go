package weekday

import (
	"time"
)

// 引数の日付が、月の何回目の曜日かを返す
func CountWeekdayInMonth(target time.Time) int {
	// 月末日を取得
	eom := time.Date(target.Year(), target.Month(), 1, 0, 0, 0, 0, target.Location()).AddDate(0, 1, -1)

	weekdayInMonth := 1
	for i := 1; i <= eom.Day(); i++ {
		t := time.Date(target.Year(), target.Month(), i, 0, 0, 0, 0, target.Location())
		if target.Day() == t.Day() {
			break
		}
		if t.Day()%7 != 0 {
			continue
		}

		weekdayInMonth++
	}

	return weekdayInMonth
}
