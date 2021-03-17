package raid

import (
	"fmt"
)

// RAID type
type Type string

const (
	Type10 Type = "10"
	Type60 Type = "60"
)

var Types = [...]Type{Type10, Type60}

// BuildArgs is arguments for Build
type BuildArgs struct {
	DisksCount int
	DiskSize   float64
	DatasCount int
	DataSize   float64
}

// Build disks with specific RAID architecture and calculate their redundancy
func Build(t Type, args BuildArgs) (disks []Disk, red float64, err error) {
	switch t {
	case Type10:
		disks, red, err = build10(args.DisksCount, args.DiskSize, args.DatasCount, args.DataSize)
		if err != nil {
			return
		}
	case Type60:
		disks, red, err = build60(args.DisksCount, args.DiskSize, args.DatasCount, args.DataSize)
		if err != nil {
			return
		}
	default:
		return nil, 0, fmt.Errorf("raid type must be one of: %v", Types)
	}
	return
}
