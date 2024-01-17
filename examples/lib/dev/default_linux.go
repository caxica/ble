package dev

import (
	"github.com/caxica/ble"
	"github.com/caxica/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
