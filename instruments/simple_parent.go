package instruments

import (
  "math"
  "github.com/carsonbaker/copenhagen/global"
)

type SimpleParent struct { }

func (SimpleParent) RunBuffer(generator func(int) int16, duration float64) []int16 {
  var sample_size = int(math.Floor(duration * global.Samplerate))
  buffer := make([]int16, sample_size)
  for i := 0; i < sample_size; i += 1 {
    buffer[i] = generator(i)
  }
  return buffer
}
