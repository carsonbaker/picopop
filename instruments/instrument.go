package instruments

import (
  "../effects"
)

type Instrument interface {
  Play(freq float64, duration int) []int16
}

type EffectedInstrument struct {
  enveloper effects.Enveloper
}

func RunBuffer(generator func(int) int16, duration int) []int16 {
  buffer := make([]int16, duration)
  for i := 0; i < duration; i += 1 {
    buffer[i] = generator(i)
  }
  return buffer
}

func (ei EffectedInstrument) EffectBuffer(buffer []int16) []int16 {
  return ei.enveloper.Filter(buffer)
}
