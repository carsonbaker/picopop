package effects



type Enveloper struct {

}

func (e Enveloper) Filter(audio []int16) []int16 {

  audio_len := len(audio)

  // linearly advance the first n samples
  attack_length := 500 // in samples
  attack_start := 0
  for i := attack_start; i < attack_length; i++ {
    audio[i] = int16(float64(audio[i]) * (float64(i) / float64(attack_length)))
  }

  // linearly decay the next n samples
  decay_length := 500 // in samples
  sustain_level := 0.5 // 0-100%
  j := 0
  for i := attack_length; i < attack_length + decay_length; i++ {
    distance := decay_length - j
    level := sustain_level * (1.0 / float64(distance)) * float64(i) + 1.0
    // fmt.Println(int16(level))
    audio[i] = int16(float64(audio[i]) * level)
    j++
  }

  // sustain the next n samples
  // decay_length := 500 // in samples
  // sustain_level := .5 // 0-100%
  // for i := attack_length; i < attack_length + decay_length; i++ {
  //   audio[i] = int16(float64(audio[i]) * (float64(i) / float64(attack_length)))
  // }

  // linearly release the last n samples
  release_length := 2000 // in samples
  release_start := audio_len - release_length
  for i := release_start; i < len(audio); i++ {
    distance := audio_len - i 
    audio[i] = int16(float64(audio[i]) * (float64(distance) / float64(release_length)))
  }

  return audio

}
