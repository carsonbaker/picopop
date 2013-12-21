package instruments

import (
  "math"
)

const Samplerate      = 8000
const Channels        = 2
const TwoPi           = math.Pi * 2
const BitDepth        = 16
const BytesPerSample  = Channels * BitDepth / 8
const FreqRad         = TwoPi / Samplerate

type SimpleParent struct { }

func (SimpleParent) RunBuffer(generator func(int) int16, duration float64) []byte {

  var sample_size = int(math.Floor(duration * Samplerate))
  buffer := make([]byte, sample_size * BytesPerSample)

  cursor := 0

  for i := 0; i < sample_size; i += 1 {

    val := generator(i)

    // TODO -- the following lines DO NOT WORK!
    // volume := 0.5
    // sample := val * int16(1.0 / volume)

    sample := val
    var h, l uint8 = uint8(sample >> 8), uint8(sample & 0xff)

    buffer[cursor] = h
    buffer[cursor+1] = l

    // for the moment just mirroring the left channel with the right
    buffer[cursor+2] = h
    buffer[cursor+3] = l

    cursor += BytesPerSample

  }
  return buffer
}
