package raid

import (
	"errors"
	"fmt"
)

// build10 returns disks built with RAID-10 architecture and their redundancy
func build10(disksCount int, diskSize float64, dataCount int, dataSize float64) ([]Disk, float64, error) {
	if disksCount < 4 || disksCount%2 != 0 {
		return nil, 0, errors.New("disks count must be even (minimum 4)")
	}

	stripeDisksCount := disksCount / 2
	dataSizePerDisk := dataSize / float64(stripeDisksCount)
	if dataSizePerDisk*float64(dataCount) > diskSize {
		return nil, 0, fmt.Errorf("insufficient disk size, increase it by %f GB",
			dataSizePerDisk*float64(dataCount)-diskSize)
	}

	var disks []Disk
	for i := 0; i < stripeDisksCount; i++ {
		disk := Disk{ID: len(disks) + 1}
		for j := 0; j < dataCount; j++ {
			disk.Fragments = append(disk.Fragments, DiskFragment{
				Label:  fmt.Sprintf("%c%d", 'A'+j, i+1),
				SizeGB: dataSizePerDisk,
			})
		}
		disks = append(disks, disk, Disk{
			ID:        disk.ID + 1,
			Fragments: disk.Fragments,
		})
	}

	used := dataSizePerDisk * float64(dataCount) * float64(disksCount)
	return disks, Redundancy(disksCount, diskSize, used), nil
}
