// This file is subject to a BSD license.
// Its contents can be found in the enclosed LICENSE file.

package ao

// #include <ao/ao.h>
import "C"

// makeOptions turns the given map into a linked list of ao_option structs.
func makeOptions(m map[string]string) *C.ao_option {
	if len(m) == 0 {
		return nil
	}

	var opt *C.ao_option

	for k, v := range m {
		C.ao_append_option(
			&opt,
			C.CString(k),
			C.CString(v),
		)
	}

	return opt
}

// freeOptions frees previously allocated ao_option sets.
func freeOptions(opt *C.ao_option) {
	C.ao_free_options(opt)
}
