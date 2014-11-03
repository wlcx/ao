## Morse

This program demonstrates how to use the libao bindings to play morse code
signals.

The application supports [international Morse code symbols][wmc].
Any unknown character is reported through stderr and is otherwise ignored.

[wmv]: http://en.wikipedia.org/wiki/Morse_code


### Install

	go get github.com/jteeuwen/ao/morse


### Usage

	$ morse [options] <sentence>


