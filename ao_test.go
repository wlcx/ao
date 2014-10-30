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

const driverName = "null" // special driver for testing.

var (
	options = map[string]string{
		"debug":   "",
		"verbose": "",
	}

	format = &SampleFormat{
		ByteOrder: EndianNative,
		Bits:      16,
		Rate:      44100,
		Channels:  2,
	}
)

func Test(t *testing.T) {
	Init()
	defer Shutdown()

	driver := DriverByName(driverName)
	if driver < 0 {
		t.Errorf("No driver named %q", driverName)
		return
	}

	dev, err := OpenLive(driver, format, options)
	if err != nil {
		t.Error(err)
		return
	}

	defer dev.Close()

	fmt.Println(dev)
}
