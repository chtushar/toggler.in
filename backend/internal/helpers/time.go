package helpers

import "time"

func UNIXTimestampFromNow(minutes int) int64 {
	now := time.Now().Local()
	return now.Add(time.Minute * time.Duration(minutes)).Unix()
}