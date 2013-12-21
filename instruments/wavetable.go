package instruments

import (
  "math"
  "io/ioutil"
  "github.com/carsonbaker/copenhagen/util"
  "fmt"
)

type WavetableGenerator struct {
  SimpleParent
}


func changePitch(input []int16, stretch float64) []int16 {
  index := float64(0.0)
  input_len := len(input)
  output_len := int(float64(input_len) / stretch)
  output := make([]int16, output_len)
  o_count := 0

  for o_count < output_len {
    i1 := int(index) % input_len
    i2 := (int(index) + 1) % input_len

    frac := index - math.Floor(index) // will be between 1 and 0
    s1 := input[i1]
    s2 := input[i2]

    output[o_count] += int16(float64(s1) + (float64(s2 - s1) * frac))

    index += stretch

    o_count++
  }
  return output
}

func waveForPitch(freq float64, path string) []int16 {
  raw, err := ioutil.ReadFile(path)
  if err != nil { panic(err) }
  shorts := util.BytesToInt16(raw)
  return changePitch(shorts, freq / 100.0) // 100 hz is the root frequency of the file audio
}

func (s WavetableGenerator) Play(freq float64, duration float64) []byte {

  waveform_count := 0
  first := true
  waveform := waveForPitch(freq, "samples/wavetable/vocal/v1.raw")
  
  generator := func(pos int) int16 {

    pos_in_cycle := pos % len(waveform)

    if pos_in_cycle == 0 && !first {
      waveform_count = waveform_count % 104 + 1
      fmt.Println("Using waveform", (waveform_count/8)+1)
      waveform_path := fmt.Sprintf("samples/wavetable/vocal/v%d.raw",(waveform_count/8)+1)
      waveform = waveForPitch(freq, waveform_path)
    }

    first = false

    // fmt.Println(floats[pos_in_cycle]))

    return waveform[pos_in_cycle]
    // return b[int(pos)] b[int(pos)+1]

  }
  return s.RunBuffer(generator, duration)

}
