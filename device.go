// This file is subject to a BSD license.
// Its contents can be found in the enclosed LICENSE file.

package ao

// #include <ao/ao.h>
import "C"
import (
	"unsafe"
)

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

// OpenFile open a file for audio output. The file format is determined by
// the audio driver used.
//
// The driver id can be retrieved by either DriverId() or DefaultDriver().
// Live output drivers cannot be used with this function. Use OpenLive instead.
// Some file formats (notably .WAV) cannot be correctly written to non-seekable
// files (like stdout).
//
// The sample format defines the format of the output stream.
//
// If overwrite is true, the file is automatically overwritten.
// Otherwise a preexisting file will cause the function to report a failure.
//
// The optional map defines device configuration settings.
// Refer to https://xiph.org/ao/doc/drivers.html for a list of options
// supported by the current driver.
//
// Returns an error if the device could not be opened.
// Be sure to call Device.Close() once you are done with it.
func OpenFile(driver int, filename string, overwrite bool, fmt *SampleFormat, options map[string]string) (*Device, error) {
	coptions := makeOptions(options)
	cfilename := C.CString(filename)

	defer func() {
		freeOptions(coptions)
		C.free(unsafe.Pointer(cfilename))
	}()

	var coverwrite C.int
	if overwrite {
		coverwrite = 1
	}

	dev, err := C.ao_open_file(
		C.int(driver),
		cfilename,
		coverwrite,
		fmt.toC(),
		coptions,
	)

	if err != nil {
		return nil, err
	}

	return &Device{dev}, nil
}

// OpenLive opens a live playback audio device for output.
//
// The driver id can be retrieved by either DriverId() or DefaultDriver().
// File output drivers cannot be used with this function. Use OpenFile() instead.
//
// The sample format defines the format of the output stream.
//
// The optional map defines device configuration settings.
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
