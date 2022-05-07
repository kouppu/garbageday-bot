package schedule

import "fmt"

func CreateGarbagedayMessage(garbagedays []Garbageday) string {
	message := "今日のごみ収集は"
	if len(garbagedays) == 0 {
		message += "\nありません。"
	}
	for _, gagarbageday := range garbagedays {
		message += fmt.Sprintf("「%s」", gagarbageday.Name)
	}
	message += "です。"

	return message
}
