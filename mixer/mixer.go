package mixer

// TODO -- the following lines DO NOT WORK!
// volume := 0.5
// sample := val * int16(1.0 / volume)

func MixToBytes(pcm []int16, bytesPerSample int) []byte {

  buffer := make([]byte, len(pcm) * bytesPerSample)
  cursor := 0

  for i := 0; i < len(pcm); i++ {

    sample := pcm[i]
    var h, l uint8 = uint8(sample >> 8), uint8(sample & 0xff)

    buffer[cursor] = h
    buffer[cursor+1] = l

    // for the moment just mirroring the left channel with the right
    buffer[cursor+2] = h
    buffer[cursor+3] = l

    cursor += bytesPerSample

  }

  return buffer

}
