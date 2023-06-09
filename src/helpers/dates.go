package helpers

import (
	"time"
)

func FormatterDateToGo(dateStr string) (string, error) {
	t, err := time.Parse(time.RFC3339Nano, dateStr)
	if err != nil {
		return "", err
	}

	formatted := t.Format("2006-01-02")

	return formatted, nil
}

func AddOneMonth(dateStr string) (string, error) {
	t, err := time.Parse(time.RFC3339Nano, dateStr)
	if err != nil {
		panic(err)
	}

	newTime := t.AddDate(0, 1, 0)

	// Format the time.Time value using a Go time format string.
	formatted := newTime.Format("2006-01-02")

	return formatted, nil
}
