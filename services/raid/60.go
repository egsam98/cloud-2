package raid

import (
	"errors"
	"fmt"
	"strconv"
)

// build60 returns disks built with RAID-60 architecture and their redundancy
func build60(disksCount int, diskSize float64, dataCount int, dataSize float64) ([]Disk, float64, error) {
	if disksCount < 8 || disksCount%2 != 0 {
		return nil, 0, errors.New("disks count must be even (minimum 8)")
	}

	dataSizePerDisk := dataSize / float64(disksCount-2*2) // two Hamming codes used twice per one data => 2 * 2
	if dataSizePerDisk*float64(dataCount) > diskSize {
		return nil, 0, fmt.Errorf("insufficient disk size, increase it by %f GB",
			dataSizePerDisk*float64(dataCount)-diskSize)
	}

	build6 := func(disksCount, startID, startDataID int) []Disk {
		disks := make([]Disk, disksCount)
		for i := 0; i < disksCount; i++ {
			disks[i] = Disk{ID: startID + i}
		}

		for j := 0; j < dataCount; j++ {
			lastID := startDataID
			for i := 0; i < disksCount; i++ {
				pPos := disksCount - j - 2
				for pPos < 0 {
					pPos = disksCount + pPos
				}
				qPos := disksCount - j - 1
				for qPos < 0 {
					qPos = disksCount + qPos
				}

				var num string
				var sizeGB float64
				switch i {
				case pPos:
					num = "_p"
					sizeGB = 0
				case qPos:
					num = "_q"
					sizeGB = 0
				default:
					num = strconv.Itoa(lastID)
					sizeGB = dataSizePerDisk
					lastID++
				}
				disks[i].Fragments = append(disks[i].Fragments, DiskFragment{
					Label:  fmt.Sprintf("%c%s", 'A'+j, num),
					SizeGB: sizeGB,
				})
			}
		}
		return disks
	}

	var disks []Disk
	count := disksCount / 2
	disks = append(disks, build6(count, 1, 1)...)
	disks = append(disks, build6(count, 1+count, 1+count-2)...)

	var used float64
	for _, disk := range disks {
		for _, fragment := range disk.Fragments {
			used += fragment.SizeGB
		}
	}
	return disks, Redundancy(disksCount, diskSize, used), nil
}
