package schedule

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type Schedule struct {
	Garbagedays []Garbageday `json:"garbageDays"`
}

type Garbageday struct {
	Name           string `json:"name"`
	Weekday        int    `json:"weekday"`
	WeekdayInMonth int    `json:"weekdayInMonth"`
}

func ReadJSON(filename string) (*Schedule, error) {
	file, err := ioutil.ReadFile(filepath.Join(filename))
	if err != nil {
		return nil, err
	}

	var schedule Schedule
	json.Unmarshal(file, &schedule)

	return &schedule, nil
}
