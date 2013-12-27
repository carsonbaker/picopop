package effects

type Enveloper struct {

}

func (e Enveloper) Filter(audio []int16) []int16 {

  // linearly advance the first n samples
  attack_length := 2000 // in samples
  attack_start := 0
  for i := attack_start; i < attack_length; i++ {
    audio[i] = int16(float64(audio[i]) * (float64(i) / float64(attack_length)))
  }

  // linearly decay the last n samples
  decay_length := 5000 // in samples
  decay_start := len(audio) - decay_length
  for i := decay_start; i < len(audio); i++ {
    distance := len(audio) - i 
    audio[i] = int16(float64(audio[i]) * (float64(distance) / float64(decay_length)))
  }

  return audio

}
