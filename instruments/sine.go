package instruments

import (
  "math"
  "github.com/carsonbaker/copenhagen/util"
)

type SineGenerator struct {
  SimpleParent
}

func (s SineGenerator) Play(freq float64, duration float64) []byte {
  generator := func(pos int) int16 {
    phase_angle := FreqRad * freq * float64(pos)
    return util.AudioFloatToInt16(math.Sin(phase_angle))
  }
  return s.RunBuffer(generator, duration)
}
