// This file is subject to a BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"fmt"
	"os"
	"strings"
)

// Translate translates the given text into a sequence of morse code signals.
func Translate(text string, samples *SampleSet) []uint16 {
	var out []uint16

	text = strings.ToLower(text)
	words := strings.Fields(text)

	for i, word := range words {
		out = append(out, translateWord(word, samples)...)

		if i < len(words)-1 {
			out = append(out, samples.WordPause...)
		}
	}

	return out
}

// translateWord translates a single word to morse code samples.
//
// Ref: http://en.wikipedia.org/wiki/Morse_code
func translateWord(word string, s *SampleSet) []uint16 {
	var out []uint16

	for i, c := range word {
		switch c {
		case 'a': // ._
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)

		case 'b': // _...
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case 'c': // _._.
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case 'd': // _..
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case 'e': // .
			out = append(out, s.Dot...)

		case 'f': // .._.
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case 'g': // __.
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case 'h': // ....
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case 'i': // ..
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case 'j': // .___
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)

		case 'k': // _._
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)

		case 'l': // ._..
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case 'm': // __
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)

		case 'n': // _.
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case 'o': // ___
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)

		case 'p': // .__.
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case 'q': // __._
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)

		case 'r': // ._.
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case 's': // ...
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case 't': // _
			out = append(out, s.Dash...)

		case 'u': // .._
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)

		case 'v': // ..._
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)

		case 'w': // .__
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)

		case 'x': // _.._
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)

		case 'y': // _.__
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)

		case 'z': // __..
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case '1': // .____
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)

		case '2': // ..___
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)

		case '3': // ...__
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)

		case '4': // ...._
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)

		case '5': // .....
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case '6': // _....
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case '7': // __...
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case '8': // ___..
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case '9': // ____.
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dot...)

		case '0': // _____
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)
			out = append(out, s.CharacterPause...)
			out = append(out, s.Dash...)

		default:
			fmt.Fprintf(os.Stderr, "unknown character: %c\n", c)
		}

		if i < len(word)-1 {
			out = append(out, s.LetterPause...)
		}
	}

	return out
}
