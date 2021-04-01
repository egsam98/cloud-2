package raid

import (
	"errors"
	"fmt"
	"math"
)

// build1E returns disks built with RAID-1E architecture and their redundancy
func build1E(disksCount int, diskSize float64, dataCount int, dataSize float64) ([]Disk, float64, error) {
	if disksCount < 3 {
		return nil, 0, errors.New("minimum disks count must be 3")
	}

	partsCount := int(math.Trunc(float64(disksCount) / 2))
	dataSizePerDisk := dataSize / float64(partsCount)

	disks := make([]Disk, disksCount)
	fragments := make([]DiskFragment, 0, dataCount*partsCount*2)
	for i := 0; i < disksCount; i++ {
		disks[i] = Disk{ID: i + 1}
	}

	for k := 0; k < dataCount; k++ {
		for i := 0; i < partsCount; i++ {
			frag := DiskFragment{
				Label:  fmt.Sprintf("%c%d", 'A'+k, i+1),
				SizeGB: dataSizePerDisk,
			}
			fragments = append(fragments, frag, frag)
		}
	}

	for i, frag := range fragments {
		i %= disksCount
		disks[i].Fragments = append(disks[i].Fragments, frag)
		fragsLen := float64(len(disks[i].Fragments))
		if dataSizePerDisk*fragsLen > diskSize {
			return nil, 0, fmt.Errorf("insufficient disk size, increase it by %f GB",
				dataSizePerDisk*fragsLen-diskSize)
		}
	}

	used := dataSizePerDisk * float64(len(fragments))
	return disks, Redundancy(disksCount, diskSize, used), nil
}
