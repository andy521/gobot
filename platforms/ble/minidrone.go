package ble

import (
	"fmt"

	"github.com/hybridgroup/gobot"
)

var _ gobot.Driver = (*BLEMinidroneDriver)(nil)

type BLEMinidroneDriver struct {
	name       string
	connection gobot.Connection
	stepsfa0a uint16
	stepsfa0b uint16
	stepsfa0c uint16
	gobot.Eventer
}

// NewBLEMinidroneDriver creates a BLEMinidroneDriver by name
func NewBLEMinidroneDriver(a *BLEAdaptor, name string) *BLEMinidroneDriver {
	n := &BLEMinidroneDriver{
		name:       name,
		connection: a,
		Eventer:    gobot.NewEventer(),
	}

	return n
}
func (b *BLEMinidroneDriver) Connection() gobot.Connection { return b.connection }
func (b *BLEMinidroneDriver) Name() string                 { return b.name }

// adaptor returns BLE adaptor
func (b *BLEMinidroneDriver) adaptor() *BLEAdaptor {
	return b.Connection().(*BLEAdaptor)
}

// Start tells driver to get ready to do work
func (b *BLEMinidroneDriver) Start() (errs []error) {
	return
}

// Halt stops minidrone driver (void)
func (b *BLEMinidroneDriver) Halt() (errs []error) { return }

func (b *BLEMinidroneDriver) Init() (err error) {
	b.stepsfa0b++
	buf := []byte{0x04, byte(b.stepsfa0b), 0x00, 0x04, 0x01, 0x00, 0x32, 0x30, 0x31, 0x34, 0x2D, 0x31, 0x30, 0x2D, 0x32, 0x38, 0x00}
	err = b.adaptor().WriteCharacteristic("9a66fa000800919111e4012d1540cb8e", "9a66fa0b0800919111e4012d1540cb8e", buf)

	f := func(b []byte, e error) {
			fmt.Printf("battery: %d\n", b[len(b)-1])
	}

	b.adaptor().SubscribeNotify("9a66fb000800919111e4012d1540cb8e", "9a66fb0f0800919111e4012d1540cb8e", f)
	return err
}
