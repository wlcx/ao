// This file is subject to a BSD license.
// Its contents can be found in the enclosed LICENSE file.

package ao

// #include <ao/ao.h>
import "C"

// SampleFormat defines the format of audio samples.
//
// The Matrix field specifies the mapping of input channels to intended
// speaker/ouput location (or empty to specify no mapping). The matrix is
// specified as a comma seperated list of channel locations equal to the
// number and in the order of the input channels.
//
// The channel mnemonics are as follows:
//
//     L    : Left speaker, located forward and to the left of the listener.
//     R    : Right speaker, located forward and to the right of the listener.
//     C    : Center speaker, located directly forward of the listener between
//            the Left and Right speakers.
//     M    : Monophonic, a virtual speaker for single-channel output.
//     CL   : Left of Center speaker (used in some Widescreen formats),
//            located forward of the listener between the Center and Left
//            speakers. Alternatively referred to as 'Left Center'.
//     CR   : Right of Center speaker (used in some Widescreen formats),
//            located forward of the listener between the Center and Right
//            speakers. Alternatively referred to as 'Right Center'.
//     BL   : Back Left speaker, located behind and to the left of the
//            listener. Alternatively called 'Left Surround' (primarily by
//            Apple) or 'Surround Rear Left' (primarily by Dolby).
//     BR   : Back Right speaker, located behind and to the right of the
//            listener. Alternatively called 'Right Surround' (primarily by
//            Apple) or 'Surround Rear Right' (primarily by Dolby).
//     BC   : Back Center speaker, located directly behind the listener.
//            Alternatively called 'Center Surround' (primarily by Apple) or
//            'Surround Rear Center' (primarily by Dolby).
//     SL   : Side Left speaker, located directly to the listener's left side.
//            The Side Left speaker is also referred to as 'Left Surround
//            Direct' (primarily by Apple) or 'Surround Left' (primarily by Dolby)
//     SR   : Side Right speaker, located directly to the listener's right
//            side. The Side Right speaker is also referred to as 'Right
//            Surround Direct' (primarily by Apple) or 'Surround Right'
//            (primarily by Dolby)
//     LFE  : Low Frequency Effect (subwoofer) channel. This is channel is
//            usually lowpassed and meant only for bass, though in some recent
//            formats it is a discrete, full-range channel. Microsoft calls
//            this the 'Low Frequency' channel.
//     X    : Unused/Invalid channel, to be dropped in the driver and not
//            output to any speaker.
//     A1   : 'auxiliary' channels, not mapped to a location. Intended for
//     A2      driver-specific use.
//     A3
//     A4
//
// Note: the 'surround' speakers referred to in other systems can be either
// the side or back speakers depending on vendor. For example, Apple calls the
// back speakers 'surround' and the side speakers 'direct surround'.
//
// Dolby calls the back speakers 'surround rear' and the side speakers
// 'surround', resulting in a direct naming conflict. For this reason,
// libao explicitly refers to speakers as 'back' and 'side' rather than
// 'surround'.
//
// Refer to the Matrix*** constants for examples of common matrix
// configurations.
type SampleFormat struct {
	Matrix    string    // String defining the channel input matrix.
	Bits      int       // Bits per sample.
	Rate      int       // Samples per second per channel.
	Channels  int       // Number of audio channels.
	ByteOrder ByteOrder // Byte ordering of the sample data. Defaults to EndianNative
}

// Bitrate computes the bitrate for a given sample in bits per second.
func (sf *SampleFormat) Bitrate() int {
	return sf.Rate * sf.Bits * sf.Channels
}

// toC converts the sample format to its C equivalent.
func (sf *SampleFormat) toC() *C.ao_sample_format {
	csf := &C.ao_sample_format{
		bits:        C.int(sf.Bits),
		rate:        C.int(sf.Rate),
		channels:    C.int(sf.Channels),
		byte_format: C.int(sf.ByteOrder),
		matrix:      C.CString(sf.Matrix),
	}

	// Matrix should be explicitely set to NULL if the string is empty.
	// A zero-length string is not considered valid.
	if len(sf.Matrix) == 0 {
		csf.matrix = nil
	}

	if csf.byte_format == 0 {
		csf.byte_format = C.AO_FMT_NATIVE
	}

	return csf
}

// Common examples of channel orderings.
// These can be assigned as-is to the SampleFormat.Matrix field.
//
// Channel mappings for most formats are usually not tied to a single
// channel matrix (there are a few exceptions like Vorbis I, where the number
// of channels always maps to a specific order); these examples cannot
// be blindly applied to a given file type and number of channels.
//
// The mapping must usually be read or intuited from the input.
const (
	MatrixDefault      = "L,R"                   // Stereo ordering in virtually all file formats
	MatrixQuadraphonic = "L,R,BL,BR"             // Quadraphonic ordering for most file formats
	Matrix51           = "L,R,C,LFE,BR,BL"       // Channel order of a 5.1 WAV or FLAC file
	Matrix71           = "L,R,C,LFE,BR,BL,SL,SR" // Channel order of a 7.1 WAV or FLAC file
	Matrix51Vorbis     = "L,C,R,BR,BL,LFE"       // Channel order of a six channel (5.1) Vorbis I file
	Matrix71Vorbis     = "L,C,R,BR,BL,SL,SR,LFE" // Channel order of an eight channel (7.1) Vorbis file
	MatrixAIFF         = "L,CL,C,R,RC,BC"        // Channel order of a six channel AIFF[-C] file
)

// ByteOrder defines endianess for sample data.
type ByteOrder int

// Known byte orders.
const (
	EndianLittle ByteOrder = C.AO_FMT_LITTLE // Samples are in little-endian order.
	EndianBig    ByteOrder = C.AO_FMT_BIG    // Samples are in big-endian order
	EndianNative ByteOrder = C.AO_FMT_NATIVE // Samples are in the native ordering of the computer.
)
