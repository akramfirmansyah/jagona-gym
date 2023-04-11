package utils

import "time"

func ParseTime(data string) (time.Time, error) {
	layout, err := time.Parse("2006-01-02", data)

	return layout, err
}
