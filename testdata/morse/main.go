// This file is subject to a BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jteeuwen/ao"
)

// Config defines application properties.
type Config struct {
	Driver string
	Note   string
	Base   float64
	Unit   uint
	Volume uint
}

func main() {
	var sf ao.SampleFormat
	var cfg Config

	// Process command line arguments.
	text := parseArgs(&sf, &cfg)

	// Create morse code tones for the specified frequency.
	frequency := Scale(cfg.Base, cfg.Note)
	samples := MakeSamples(
		sf.Rate,
		sf.Channels,
		cfg.Unit,
		cfg.Volume,
		frequency,
	)

	// Convert the input text into a sequence of morse code audio samples.
	morsecode := Translate(text, samples)

	// Initialize libao.
	ao.Init()

	// Get the requested- or system's default audio driver.
	var id int
	if len(cfg.Driver) == 0 {
		id = ao.DefaultDriver()
	} else {
		id = ao.DriverByName(cfg.Driver)
	}

	if id == -1 {
		fmt.Fprintln(os.Stderr, "no valid audio driver found")
		ao.Shutdown()
		os.Exit(1)
	}

	// Open an audio device.
	device, err := ao.OpenLive(id, &sf, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "open audio device:", err)
		ao.Shutdown()
		os.Exit(1)
	}

	// Play the samples.
	device.PlayU16(morsecode)
	device.Close()
	ao.Shutdown()
}

// parseArgs processes all command line arguments and ensures there
// are sane values. The given parameters are filled out accordingly.
// This function writes the appropriate errors to stderr and exits the
// program when poop hits the fan.
func parseArgs(sf *ao.SampleFormat, cfg *Config) string {
	cfg.Unit = 50
	cfg.Note = "F#"
	cfg.Base = 440
	cfg.Volume = 100

	sf.ByteOrder = ao.EndianNative
	sf.Matrix = ao.MatrixDefault
	sf.Bits = 16
	sf.Rate = 8000
	sf.Channels = 2

	flag.IntVar(&sf.Rate, "r", sf.Rate, "Number of samples per second per channel.")
	flag.IntVar(&sf.Channels, "c", sf.Channels, "Number of channels.")
	flag.StringVar(&sf.Matrix, "m", sf.Matrix, "Channel matrix for audio driver.")
	flag.StringVar(&cfg.Driver, "d", cfg.Driver, "Name of audio driver to use. Empty value implies default system driver.")
	flag.StringVar(&cfg.Note, "n", cfg.Note, "Note to play at: C, C#, D, D#, E, F, F#, G, G#, A, A#, B")
	flag.Float64Var(&cfg.Base, "f", cfg.Base, "Base frequency in herz for tone scale (value of note A).")
	flag.UintVar(&cfg.Unit, "u", cfg.Unit, "Duration of 1 unit (dot) in milliseconds.")
	flag.UintVar(&cfg.Volume, "v", cfg.Volume, "Volume of output in range 0..100")

	flag.Usage = func() {
		fmt.Println("usage:", os.Args[0], "[options] <sentence>")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() == 0 || len(flag.Arg(0)) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if cfg.Base == 0 {
		cfg.Base = 440
	}

	if cfg.Unit == 0 {
		cfg.Unit = 10
	}

	if cfg.Volume > 100 {
		cfg.Volume = 100
	}

	cfg.Note = strings.ToUpper(cfg.Note)
	switch cfg.Note {
	case "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B":
	default:
		fmt.Fprintln(os.Stderr, "unknown note:", cfg.Note)
		flag.Usage()
		os.Exit(1)
	}

	return flag.Arg(0)
}
