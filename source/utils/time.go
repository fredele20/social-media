package utils

import "time"

func CurrentTime() time.Time {
	return time.Now().UTC().Truncate(time.Second)
}

func ConvertUtcToNigerianTime(timeStamp time.Time) (time.Time, error) {
	loc, e := time.LoadLocation("Africa/Algiers")
	if e != nil {
		return timeStamp, e
	}

	return timeStamp.In(loc), nil
}
