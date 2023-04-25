// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package googlecompute

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatBlockDevice is an auto-generated flat version of BlockDevice.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatBlockDevice struct {
	AttachmentMode    *string                    `mapstructure:"attachment_mode" cty:"attachment_mode" hcl:"attachment_mode"`
	DeviceName        *string                    `mapstructure:"device_name" cty:"device_name" hcl:"device_name"`
	DiskEncryptionKey *FlatCustomerEncryptionKey `mapstructure:"disk_encryption_key" cty:"disk_encryption_key" hcl:"disk_encryption_key"`
	DiskName          *string                    `mapstructure:"disk_name" cty:"disk_name" hcl:"disk_name"`
	InterfaceType     *string                    `mapstructure:"interface_type" cty:"interface_type" hcl:"interface_type"`
	IOPS              *int                       `mapstructure:"iops" cty:"iops" hcl:"iops"`
	KeepDevice        *bool                      `mapstructure:"keep_device" cty:"keep_device" hcl:"keep_device"`
	ReplicaZones      []string                   `mapstructure:"replica_zones" required:"false" cty:"replica_zones" hcl:"replica_zones"`
	SourceVolume      *string                    `mapstructure:"source_volume" cty:"source_volume" hcl:"source_volume"`
	VolumeSize        *int                       `mapstructure:"volume_size" required:"true" cty:"volume_size" hcl:"volume_size"`
	VolumeType        *BlockDeviceType           `mapstructure:"volume_type" required:"true" cty:"volume_type" hcl:"volume_type"`
}

// FlatMapstructure returns a new FlatBlockDevice.
// FlatBlockDevice is an auto-generated flat version of BlockDevice.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*BlockDevice) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatBlockDevice)
}

// HCL2Spec returns the hcl spec of a BlockDevice.
// This spec is used by HCL to read the fields of BlockDevice.
// The decoded values from this spec will then be applied to a FlatBlockDevice.
func (*FlatBlockDevice) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"attachment_mode":     &hcldec.AttrSpec{Name: "attachment_mode", Type: cty.String, Required: false},
		"device_name":         &hcldec.AttrSpec{Name: "device_name", Type: cty.String, Required: false},
		"disk_encryption_key": &hcldec.BlockSpec{TypeName: "disk_encryption_key", Nested: hcldec.ObjectSpec((*FlatCustomerEncryptionKey)(nil).HCL2Spec())},
		"disk_name":           &hcldec.AttrSpec{Name: "disk_name", Type: cty.String, Required: false},
		"interface_type":      &hcldec.AttrSpec{Name: "interface_type", Type: cty.String, Required: false},
		"iops":                &hcldec.AttrSpec{Name: "iops", Type: cty.Number, Required: false},
		"keep_device":         &hcldec.AttrSpec{Name: "keep_device", Type: cty.Bool, Required: false},
		"replica_zones":       &hcldec.AttrSpec{Name: "replica_zones", Type: cty.List(cty.String), Required: false},
		"source_volume":       &hcldec.AttrSpec{Name: "source_volume", Type: cty.String, Required: false},
		"volume_size":         &hcldec.AttrSpec{Name: "volume_size", Type: cty.Number, Required: false},
		"volume_type":         &hcldec.AttrSpec{Name: "volume_type", Type: cty.String, Required: false},
	}
	return s
}