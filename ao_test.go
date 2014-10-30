// This file is subject to a BSD license.
// Its contents can be found in the enclosed LICENSE file.

package ao

import (
	"crypto/rand"
	"testing"
)

// null is a special driver for testing.
// It does not generate actual sound.
const driverName = "null"

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

	var buf [16 * 1024]byte
	for i := 0; i < 32; i++ {
		n, err := rand.Read(buf[:])
		if err != nil {
			t.Error(err)
			return
		}

		err = dev.Play(buf[:n])
		if err != nil {
			t.Error(err)
			return
		}
	}
}
