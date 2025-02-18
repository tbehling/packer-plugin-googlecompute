// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecompute

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type StepCreateDisks struct {
	DiskConfiguration []BlockDevice
}

func (s *StepCreateDisks) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)

	if !s.needToCreateDisks() {
		ui.Say("no persistent disk to create")
		return multistep.ActionContinue
	}

	driver := state.Get("driver").(Driver)
	config := state.Get("config").(*Config)

	for i, disk := range s.DiskConfiguration {
		if disk.VolumeType == LocalScratch {
			continue
		}

		if disk.SourceVolume != "" {
			continue
		}

		ui.Say(fmt.Sprintf("Creating persistent disk %s", disk.DiskName))

		_, errCh := driver.CreateDisk(disk)

		var err error
		select {
		case err = <-errCh:
		case <-time.After(config.StateTimeout):
			err = errors.New("time out while waiting for disk to create")
		}
		if err != nil {
			err := fmt.Errorf("failed to create disk: %s", err)
			ui.Say(err.Error())
			state.Put("error", err)
			return multistep.ActionHalt
		}

		if len(disk.ReplicaZones) != 0 {
			region, _ := getRegionFromZone(config.Zone)
			// Generate the source URI for attachment later
			s.DiskConfiguration[i].SourceVolume = fmt.Sprintf("projects/%s/regions/%s/disks/%s",
				config.ProjectId,
				region,
				disk.DiskName)
		} else {
			// Generate the source URI for attachment later
			s.DiskConfiguration[i].SourceVolume = fmt.Sprintf("projects/%s/zones/%s/disks/%s",
				config.ProjectId,
				config.Zone,
				disk.DiskName)
		}
	}

	return multistep.ActionContinue
}

func (s *StepCreateDisks) needToCreateDisks() bool {
	for _, cfg := range s.DiskConfiguration {
		if cfg.VolumeType == LocalScratch {
			continue
		}

		if cfg.SourceVolume != "" {
			continue
		}

		return true
	}

	return false
}

func (s *StepCreateDisks) Cleanup(state multistep.StateBag) {
	ui := state.Get("ui").(packersdk.Ui)
	config := state.Get("config").(*Config)
	driver := state.Get("driver").(Driver)

	for _, gceDisk := range s.DiskConfiguration {
		if gceDisk.KeepDevice {
			ui.Say(fmt.Sprintf("Keeping disk %q", gceDisk.DiskName))
			continue
		}

		// Scratch volumes are not to be deleted since they are
		// linked to the instance and are always automatically deleted.
		if gceDisk.VolumeType == LocalScratch {
			continue
		}

		ui.Say(fmt.Sprintf("Deleting disk %q", gceDisk.DiskName))

		zone := config.Zone
		if len(gceDisk.ReplicaZones) != 0 {
			zone, _ = getRegionFromZone(zone)
		}

		var err error

		errCh := driver.DeleteDisk(zone, gceDisk.DiskName)
		select {
		case err = <-errCh:
		case <-time.After(config.StateTimeout):
			err = errors.New("time out while waiting for disk to delete")
		}

		if err != nil {
			ui.Error(fmt.Sprintf(
				"Error deleting disk. Please delete it manually.\n\n"+
					"Name: %s\n"+
					"Error: %s", gceDisk.DiskName, err))
		} else {
			ui.Say(fmt.Sprintf("Disk %q successfully deleted", gceDisk.DiskName))
		}
	}
}
