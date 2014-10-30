// This file is subject to a BSD license.
// Its contents can be found in the enclosed LICENSE file.

package ao

// #include <ao/ao.h>
import "C"

// Device holds an opaque type defining output device data.
type Device struct {
	ptr *C.ao_device
}

// Close closes the device after use.
func (d *Device) Close() error {
	if d.ptr != nil {
		C.ao_close(d.ptr)
		d.ptr = nil
	}
	return nil
}

// OpenLive opens a live playback audio device for output.
//
// The driver id can be retrieved by either DriverId() or DefaultDriver().
// The sample format defines the format of the output stream.
//
// The options map is optional and defines device configuration settings.
// Refer to https://xiph.org/ao/doc/drivers.html for a list of options
// supported by the current driver.
//
// Returns an error if the device could not be opened.
// Be sure to call Device.Close() once you are done with it.
func OpenLive(driver int, fmt *SampleFormat, options map[string]string) (*Device, error) {
	coptions := makeOptions(options)
	defer freeOptions(coptions)

	dev, err := C.ao_open_live(C.int(driver), fmt.toC(), coptions)
	if err != nil {
		return nil, err
	}

	return &Device{dev}, nil
}
