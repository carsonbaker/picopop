package instruments

import (
  "math"
  "github.com/carsonbaker/copenhagen/util"
  "github.com/carsonbaker/copenhagen/global"
)

type SineGenerator struct {
  SimpleParent
}

func (s SineGenerator) Play(freq float64, duration float64) []int16 {
  generator := func(pos int) int16 {
    phase_angle := global.FreqRad * freq * float64(pos)
    return util.AudioFloatToInt16(math.Sin(phase_angle))
  }
  return s.RunBuffer(generator, duration)
}
