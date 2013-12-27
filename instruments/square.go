package instruments

import (
  "math"
  "github.com/carsonbaker/copenhagen/util"
  "github.com/carsonbaker/copenhagen/global"
)

type SquareGenerator struct {
  EffectedInstrument
}

func (s SquareGenerator) Play(freq float64, duration int) []int16 {
  generator := func(pos int) int16 {
    phase_angle := global.FreqRad * freq * float64(pos)
    sample := math.Sin(phase_angle)
    if sample > 0 { return util.AudioFloatToInt16(1) } else { return util.AudioFloatToInt16(0) }
  }

  buffer := RunBuffer(generator, duration)
  return s.EffectBuffer(buffer)
}
