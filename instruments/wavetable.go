package instruments

import (
  "math"
  "io/ioutil"
  "github.com/carsonbaker/copenhagen/util"
  "fmt"
)

type WavetableGenerator struct {
  EffectedInstrument
}

func InterpolateHermite4pt3oX(x []int16, ineg1 int, i0 int, i1 int, i2 int, t float64) int16 {
  x0 := float64(x[ineg1])
  x1 := float64(x[i0])
  x2 := float64(x[i1])
  x3 := float64(x[i2])

  c0 := x1
  c1 := 0.5 * (x2 - x0)
  c2 := x0 - (2.5 * x1) + (2 * x2) - (0.5 * x3)
  c3 := (0.5 * (x3 - x0)) + (1.5 * (x1 - x2))
  
  res := (((((c3 * t) + c2) * t) + c1) * t) + c0
  return int16(res)
}

func InterpolateLinear(x []int16, off0 int, off1 int, t float64) int16 {
  x0 := float64(x[off0])
  x1 := float64(x[off1])
  return int16(x0 + ((x1 - x0) * t))
}

func ChangePitch(input []int16, stretch float64) []int16 {
  index      := float64(0.0)
  input_len  := len(input)
  output_len := int(float64(input_len) / stretch)
  output     := make([]int16, output_len)
  o_count    := 0

  for o_count < output_len {
    i0 := int(index) % input_len
    i1 := (int(index) + 1) % input_len

    frac := index - math.Floor(index) // will be between 1 and 0

    if true {
      output[o_count] += InterpolateLinear(input, i0, i1, frac)
    } else {
      ineg1 := (int(index) - 1) % input_len
      i2 := (int(index) + 2) % input_len
      output[o_count] += InterpolateHermite4pt3oX(input, ineg1, i0, i1, i2, frac)
    }

    index += stretch
    o_count++
  }
  return output
}

func waveFromFile(path string) []int16 {
  raw, err := ioutil.ReadFile(path)
  if err != nil { panic(err) }
  return util.BytesToInt16(raw)
}

func waveFromFileForPitch(freq float64, path string) []int16 {
  return ChangePitch(waveFromFile(path), freq) // 100 hz is the root frequency of the file audio
}

func (s WavetableGenerator) Play(freq float64, duration int) []int16 {

  waveform       := waveFromFile("samples/wavetable/vfs48-16/1.raw")
  waveform_freq  := 100.0
  stretch        := freq / waveform_freq
  index          := 0.0
  waveform_count := 1
  repeat         := -1
  repeat_count   := 0
  wave_max       := 38

  generator := func(pos int) int16 {

    pos_in_cycle := pos % len(waveform)

    if pos_in_cycle == 0 {
      if(waveform_count != wave_max) {
        if repeat_count > repeat {

          waveform_path := fmt.Sprintf("samples/wavetable/vfs48-16/%d.raw",waveform_count)
          waveform = waveFromFile(waveform_path)
          waveform_count = (waveform_count + 1) % 39
          if waveform_count == 0 {
            waveform_count = 1 
          } 
          repeat_count = 0
        }
        repeat_count++
      }
    }

    cursor := int(index) % len(waveform)
    index += stretch
    return waveform[cursor]

  }

  return RunBuffer(generator, duration)

}
