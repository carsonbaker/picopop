package instruments

import (
  "math"
  "../util"
  "../global"
)

type SineGenerator struct {
  EffectedInstrument
}

func (s SineGenerator) Play(freq float64, duration int) []int16 {
  generator := func(pos int) int16 {
    phase_angle := global.FreqRad * freq * float64(pos)
    return util.AudioFloatToInt16(math.Sin(phase_angle))
  }

  buffer := RunBuffer(generator, duration)
  return s.EffectBuffer(buffer)
}
