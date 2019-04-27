package entity

import "time"

type Birth struct {
	Date     time.Time
	Location Location
}
