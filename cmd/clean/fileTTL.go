package clean

import (
	"io/fs"
	"time"
)

type timeOffset struct {
	days   int
	months int
	years  int
}

var fileTTL = timeOffset{
	days:   0,
	months: 6,
	years:  0,
}

func IsFileOlderThanTTL(info fs.FileInfo) bool {
	modTimeThresholdDate := info.ModTime().AddDate(fileTTL.years, fileTTL.months, fileTTL.days)
	return modTimeThresholdDate.Before(time.Now())
}
