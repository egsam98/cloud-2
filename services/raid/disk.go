package raid

import (
	"fmt"
	"strings"
)

var _ fmt.Stringer = (*Disk)(nil)

type Disk struct {
	ID        int
	Fragments []DiskFragment
}

func (d Disk) String() string {
	var ss []string
	for _, df := range d.Fragments {
		ss = append(ss, fmt.Sprintf("%s (%.3f GB)", df.Label, df.SizeGB))
	}
	return fmt.Sprintf("Disk %d: [%s]", d.ID, strings.Join(ss, ", "))
}

type DiskFragment struct {
	Label  string
	SizeGB float64
}
