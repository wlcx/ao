## ao

Package ao defines bindings for [libao](https://xiph.org/ao).
Libao is a cross-platform audio library that allows programs to output audio
using a simple API on a wide variety of platforms.


### Supported platforms

* Null output (handy for testing without a sound device)
* WAV files
* AU files
* RAW files
* OSS (Open Sound System, used on Linux and FreeBSD)
* ALSA (Advanced Linux Sound Architecture)
* aRts (Analog RealTime Synth, used by KDE)
* PulseAudio (next generation GNOME sound server)
* esd (EsounD or Enlightened Sound Daemon)
* Mac OS X
* Windows (98 and later)
* AIX
* Sun/NetBSD/OpenBSD
* IRIX
* NAS (Network Audio Server)
* RoarAudio (Modern, multi-OS, networked Sound System)
* OpenBSD's sndio


### Usage

    go get github.com/jteeuwen/ao


### Dependencies

	https://xiph.org/ao


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

