// This file is subject to a BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"math"
)

// SampleSet defines a set of 16-bit PCM samples for morse code sound generation.
type SampleSet struct {
	Dot            []uint16 // tone of 1 unit length.
	Dash           []uint16 // tone of 3 unit lengths.
	CharacterPause []uint16 // Pause between parts of a single character of 1 unit length
	LetterPause    []uint16 // Pause between characters of 3 unit lengths
	WordPause      []uint16 // Pause between words of 7 unit lengths
}

// MakeSamples creates some raw, 16-bit PCM audio samples for morse code sounds.
// These are the building blocks for all morse code characters.
func MakeSamples(rate, channels int, unit, volume uint, frequency float64) *SampleSet {
	fu := float64(unit) * 0.001
	fv := float64(volume) * 0.01

	dot := make([]float64, int(math.Ceil(float64(rate)*fu)))
	dash := make([]float64, int(math.Ceil(float64(rate)*fu*3)))
	cp := make([]float64, int(math.Ceil(float64(rate)*fu)))
	lp := make([]float64, int(math.Ceil(float64(rate)*fu*3)))
	wp := make([]float64, int(math.Ceil(float64(rate)*fu*7)))

	// Fill the dot and dash samples with a sine wave of
	// adequate length and frequency.
	makeSample(dot, rate, fv, frequency)
	makeSample(dash, rate, fv, frequency)

	// Convert amd return the float64 samples as 16-bit PCM audio.
	return &SampleSet{
		Dot:            toPCM(dot, channels),
		Dash:           toPCM(dash, channels),
		CharacterPause: toPCM(cp, channels),
		LetterPause:    toPCM(lp, channels),
		WordPause:      toPCM(wp, channels),
	}
}

// toPCM converts the given single-channel float64 sample to a multi-channel,
// 16-bit PCM sample.
func toPCM(sample []float64, channels int) []uint16 {
	var index int

	buf := make([]uint16, len(sample)*channels)

	for _, f := range sample {
		value := uint16(f * 32767)

		// Copy signal to all required channels.
		for c := 0; c < channels; c++ {
			buf[index+c] = value
		}

		index += channels
	}

	return buf
}

// makeSample creates a constant tone for the given sample rate and frequency.
func makeSample(samples []float64, rate int, volume, frequency float64) {
	for i := range samples {
		samples[i] = math.Sin(
			2*math.Pi*float64(i)/(float64(rate)/frequency),
		) * volume
	}
}

// Scale returns one of a 12-note scale of frequencies where a is the
// base frequency for note A and the rest is derived as:
//
//    fn = a × (b)^n
//
// where:
//
//    a = the given frequency for note A; e.g.: 440 Hz
//    b = 2^(1/12) ~= 1.059463094359
//
// For example:
//
//   Scale(X, Y) =>
//       X=440, Y="C"  => 261.63
//       X=440, Y="C#" => 277.18
//       X=440, Y="D"  => 293.66
//       X=440, Y="D#" => 311.13
//       X=440, Y="E"  => 329.63
//       X=440, Y="F"  => 349.23
//       X=440, Y="F#" => 369.99
//       X=440, Y="G"  => 392.00
//       X=440, Y="G#" => 415.30
//       X=440, Y="A"  => 440.00
//       X=440, Y="A#" => 466.16
//       X=440, Y="B"  => 493.88
//
// Ref: http://www.phy.mtu.edu/~suits/NoteFreqCalcs.html
func Scale(a float64, n string) float64 {
	const b = 1.059463094359
	switch n {
	case "C":
		return a * math.Pow(b, -9)
	case "C#":
		return a * math.Pow(b, -8)
	case "D":
		return a * math.Pow(b, -7)
	case "D#":
		return a * math.Pow(b, -6)
	case "E":
		return a * math.Pow(b, -5)
	case "F":
		return a * math.Pow(b, -4)
	case "F#":
		return a * math.Pow(b, -3)
	case "G":
		return a * math.Pow(b, -2)
	case "G#":
		return a * math.Pow(b, -1)
	case "A":
		return a
	case "A#":
		return a * math.Pow(b, 1)
	case "B":
		return a * math.Pow(b, 2)
	}
	return 0
}
