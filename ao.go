// This file is subject to a BSD license.
// Its contents can be found in the enclosed LICENSE file.

package ao

// #cgo pkg-config: ao
//
// #include <stdlib.h>
// #include <ao/ao.h>
import "C"
import (
	"errors"
	"sync/atomic"
	"unsafe"
)

// libInitialized is an atomically updated flag which determines if the
// ao subsystems have been initialized. It is set by Init() and
// unset in Shutdown()
var libInitialized uint32

// Init must be called before anything else in this package and
// be balanced by a call to Shutdown().
//
// It initializes the internal libao data structures and loads all of the
// available plugins. The system and user configuration files are also read
// at this time if available.
//
// Init() must be called in the main thread because several sound system
// interfaces used by libao must be initialized in the main thread.
// One example is the system aRts interface, which stores some global state
// in thread-specific keys that it fails to delete on shutdown. If aRts
// is initialized in a non-main thread that later exits, these undeleted
// keys will cause a segmentation fault.
//
// If you want to reload the configuration files without restarting your
// program, first call Shutdown(), then call Init() again. Multiple successive
// calls to either Init() or Shutdown() will be silently ignored.
func Init() {
	if atomic.CompareAndSwapUint32(&libInitialized, 0, 1) {
		C.ao_initialize()
	}
}

// Shutdown unloads all of the plugins and deallocates any internal data
// structures the library has created. It should be called prior to program exit.
//
// If you want to reload the configuration files without restarting your
// program, first call Shutdown(), then call Init() again. Multiple successive
// calls to either Init() or Shutdown() will be silently ignored.
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
func DefaultDriver() (id int, err error) {
	id = int(C.ao_default_driver_id())
	if id == -1 {
		err = errors.New("No default driver found")
	}
	return
}

// DriverByName attempts to fetch the id for the given driver.
// Refer to https://xiph.org/ao/doc/drivers.html for a list of supported
// driver names.
//
// Returns -1 if no matching driver was found.
func DriverByName(name string) (id int, err error) {
	cname := C.CString(name)
	v := C.ao_driver_id(cname)
	C.free(unsafe.Pointer(cname))
	id = int(v)
	if id == -1 {
		err = errors.New("No matching driver found")
	}
	return
}
