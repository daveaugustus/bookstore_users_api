package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02 3:4:5 pm"
)

func GetNow() time.Time {
	return time.Now()
}

func GetNowString() string {

	return GetNow().Format(apiDateLayout)
}
