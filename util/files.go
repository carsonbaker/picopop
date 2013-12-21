package util

import(
  "encoding/binary"
  "math"
  "os"
  "fmt"
)

const SIZEOF_INT16 = 2 // in bytes

func BytesToInt16(bytes []byte) []int16 {
  data := make([]int16, len(bytes)/SIZEOF_INT16)
  for i := range data {
    // assuming little endian
    data[i] = int16(binary.LittleEndian.Uint16(bytes[i*SIZEOF_INT16:(i+1)*SIZEOF_INT16]))
  }
  return data
}

func AudioFloatToInt16(val float64) int16 {
  return int16(val * (math.MaxUint16 / 2))
}

func DebugAudioToFile(audio []byte) {
  fo, err := os.Create("output.pcm")
  if err != nil { panic(err) }
    // close fo on exit and check for its returned error
    defer func() {
      if err := fo.Close(); err != nil {
          panic(err)
      }
  }()

  if _, err := fo.Write(audio); err != nil {
    panic(err)
  }
  fmt.Println("Audio written to output.pcm!")
}
