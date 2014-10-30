// This file is subject to a BSD license.
// Its contents can be found in the enclosed LICENSE file.

package ao

import (
	"fmt"
	"testing"
)

const (
	//file = "testdata/test.mp3"
	file = "testdata/test.opus"
	//file = "testdata/test.wav"
)

const driverName = "null" // null = special driver for testing.

func Test(t *testing.T) {
	Init()
	defer Shutdown()

	driver := DriverByName(driverName)
	if driver < 0 {
		t.Errorf("No driver named %q", driverName)
		return
	}

	options := map[string]string{
		"debug": "",
		"dev":   "hw:0",
	}

	sf := &SampleFormat{
		Matrix:    MatrixDefault,
		ByteOrder: EndianNative,
		Bits:      16,
		Rate:      44100,
		Channels:  2,
	}

	dev, err := OpenLive(driver, sf, options)
	if err != nil {
		t.Error(err)
		return
	}

	defer dev.Close()

	fmt.Println(dev)
}
