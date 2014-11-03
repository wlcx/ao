## Morse

This program demonstrates how to use the libao bindings to play morse code
signals.

The application supports [international Morse code symbols][wmc].
Any unknown character is reported through stderr and is otherwise ignored.

[wmc]: http://en.wikipedia.org/wiki/Morse_code


### Install

	go get github.com/jteeuwen/ao/testdata/morse


### Usage

	$ morse [options] <sentence>

To hear the morse code for a sentence using default sample format:

	$ morse "some test string"

To change the frequency of the tones:

	$ morse -f 880 "some test string"

To change the note of the signal tones:

	$ morse -n "C#" "some test string"

Run the program with the `-h` flag for more options.
