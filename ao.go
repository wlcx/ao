// This file is subject to a BSD license.
// Its contents can be found in the enclosed LICENSE file.

package ao

// #cgo pkg-config: ao
//
// #include <stdlib.h>
// #include <ao/ao.h>
import "C"
import (
	"sync/atomic"
	"unsafe"
)

// libInitialized is an atomically updated flag which determines if the
// ao subsystems have been initialized. It is set by Init() and
// unset in Shutdown()
var libInitialized uint32

// Init must be called before anything else in this package.
//
// This loads the plugins from disk, reads the libao configuration files,
// and identifies an appropriate default output driver if none is specified
// in the configuration files.
func Init() {
	if atomic.CompareAndSwapUint32(&libInitialized, 0, 1) {
		C.ao_initialize()
	}
}

// Shutdown cleans up all ffmpeg subsystems.
func Shutdown() {
	if atomic.CompareAndSwapUint32(&libInitialized, 1, 0) {
		C.ao_shutdown()
	}
}

// DefaultDriver returns the ID number of the default live output driver.
// If the configuration files specify a default driver, its ID is returned.
// Otherwise the library tries to pick a live output driver that will work
// on the host platform.
//
// If no default device is available, you may still use the NULL device
// for testing porpuses.
//
// If no audio hardware is available, it is in use, or is not in the "standard"
// configuration, this returns -1.
func DefaultDriver() int { return int(C.ao_default_driver_id()) }

// DriverByName attempts to fetch the id for the given driver.
// Refer to https://xiph.org/ao/doc/drivers.html for a list of supported
// driver names.
//
// Returns -1 if no matching driver was found.
func DriverByName(name string) int {
	cname := C.CString(name)
	v := C.ao_driver_id(cname)
	C.free(unsafe.Pointer(cname))
	return int(v)
}
