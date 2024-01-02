package filesize

import (
	"fmt"
	"math"
)

type FileSize struct {
	Size   int64
	Unit   Units
	SIUnit bool
}

type Units int

const (
	B = iota
	KB
	MB
	GB
	TB
)
const bytesMultiplier int = 1024
const bytesMultiplierInSIUnit int = 1000

func (fz *FileSize) ToBytes() string {
	var value int64
	var multiplier = bytesMultiplier
	if fz.SIUnit {
		multiplier = bytesMultiplierInSIUnit
	}

	if fz.Unit == B {
		value = fz.Size
	} else {
		value = fz.Size * int64(math.Pow(float64(multiplier), float64(fz.Unit)))
	}

	return fmt.Sprintf("%d Bytes", value)
}
