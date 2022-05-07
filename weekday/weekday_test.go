package weekday

import (
	"testing"
	"time"
)

func TestCountWeekdayInMonth(t *testing.T) {
	t.Run("第1土曜日の場合、1が返ること", func(t *testing.T) {
		date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
		actual := CountWeekdayInMonth(date)
		expected := 1
		if actual != expected {
			t.Error("failed test")
		}
	})

	t.Run("第2土曜日の場合、2が返ること", func(t *testing.T) {
		date := time.Date(2022, 1, 8, 0, 0, 0, 0, time.Local)
		actual := CountWeekdayInMonth(date)
		expected := 2
		if actual != expected {
			t.Error("failed test")
		}
	})

	t.Run("第3土曜日の場合、3が返ること", func(t *testing.T) {
		date := time.Date(2022, 1, 15, 0, 0, 0, 0, time.Local)
		actual := CountWeekdayInMonth(date)
		expected := 3
		if actual != expected {
			t.Error("failed test")
		}
	})

	t.Run("第4土曜日の場合、4が返ること", func(t *testing.T) {
		date := time.Date(2022, 1, 22, 0, 0, 0, 0, time.Local)
		actual := CountWeekdayInMonth(date)
		expected := 4
		if actual != expected {
			t.Error("failed test")
		}
	})

	t.Run("第5土曜日の場合、5が返ること", func(t *testing.T) {
		date := time.Date(2022, 1, 29, 0, 0, 0, 0, time.Local)
		actual := CountWeekdayInMonth(date)
		expected := 5
		if actual != expected {
			t.Error("failed test")
		}
	})
}
