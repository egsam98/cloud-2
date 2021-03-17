package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/egsam98/cloud-2/services/raid"
)

const DatasCount = 2

var Raid = &cobra.Command{
	Use:     "[type]",
	Example: "{binary_path} 60 --disks-count=4 --disk-size=60 --data-size=60",
	Short:   fmt.Sprintf("Build RAID-controller (available types: %v)", raid.Types),
	Args:    cobra.ExactArgs(1),
	RunE:    raidRun,
}

// Flags for Raid cmd
var (
	diskSize   float64
	disksCount int
	dataSize   float64
)

func init() {
	flags := Raid.Flags()
	flags.Float64Var(&diskSize, "disk-size", 60, "disk size (GB)")
	flags.IntVar(&disksCount, "disks-count", 4, "disks count")
	flags.Float64Var(&dataSize, "data-size", 60, "data size (GB)")
}

// raidRun is run function for Raid cmd
func raidRun(cmd *cobra.Command, args []string) (err error) {
	raidType := raid.Type(args[0])
	cmd.Printf("RAID %s\nDisks count: %d, disk size: %.3f, datas count: %d, data size: %.3f\n\n",
		raidType, disksCount, diskSize, DatasCount, dataSize)

	disks, redundancy, err := raid.Build(raidType, raid.BuildArgs{
		DisksCount: disksCount,
		DiskSize:   diskSize,
		DatasCount: DatasCount,
		DataSize:   dataSize,
	})
	if err != nil {
		return err
	}

	cmd.Println("Result:")
	for _, disk := range disks {
		cmd.Println(disk)
	}
	cmd.Printf("Redundancy: %.3f\n", redundancy)
	return nil
}
