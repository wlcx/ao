// This file is subject to a BSD license.
// Its contents can be found in the enclosed LICENSE file.

package ao

// #include <ao/ao.h>
import "C"
import (
	"errors"
	"unsafe"
)

// Device holds an opaque type defining output device data.
type Device struct {
	ptr *C.ao_device
}

// PlayU16 is the same as Play() but accepts a slice of 16 bit PCM sample data.
// This function assumes the sample format byte order is set to EndianNative.
func (d *Device) PlayU16(data []uint16) error {
	sz := len(data)
	if sz == 0 {
		return nil
	}

	return d.Play((*(*[1<<31 - 1]byte)(unsafe.Pointer(&data[0])))[:sz*2])
}

// Play plays a block of audio data to an open device. Samples are interleaved
// by channels (Time 1, Channel 1; Time 1, Channel 2; Time 2, Channel 1; etc.)
// in the memory buffer.
//
// Returns an error if playback failed. In which case, the device should
// be closed.
func (d *Device) Play(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	if d.ptr == nil {
		return errors.New("device is closed")
	}

	if C.ao_play(
		d.ptr,
		(*C.char)(unsafe.Pointer(&data[0])),
		C.uint_32(len(data)),
	) == 0 {
		return errors.New("playback failed; device should be closed")
	}

	return nil
}

// Close closes the audio device and frees the memory allocated
// by the device structure.
//
// An error is returned if closing of the device failed.
// If this device was writing to a file, the file may be corrupted.
func (d *Device) Close() error {
	var err error

	if d.ptr != nil {
		if C.ao_close(d.ptr) <= 0 {
			err = errors.New("failed to close device correctly")
		}
		d.ptr = nil
	}

	return err
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
