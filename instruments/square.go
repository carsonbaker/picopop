package instruments

import (
  "math"
  "github.com/carsonbaker/copenhagen/util"
)

type SquareGenerator struct {
  SimpleParent
}

func (s SquareGenerator) Play(freq float64, duration float64) []byte {
  generator := func(pos int) int16 {
    phase_angle := FreqRad * freq * float64(pos)
    sample := math.Sin(phase_angle)
    if sample > 0 { return util.AudioFloatToInt16(1) } else { return util.AudioFloatToInt16(0) }
  }
  return s.RunBuffer(generator, duration)
}
