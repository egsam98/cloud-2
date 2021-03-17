package raid

// Redundancy of disks as free space / total space
func Redundancy(disksCount int, diskSize float64, used float64) float64 {
	total := diskSize * float64(disksCount)
	return (total - used) / total
}
