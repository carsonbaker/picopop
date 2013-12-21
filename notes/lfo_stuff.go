
func squareWave(freq float64, pos float64) float64 {
  if sineWave(freq, pos) > 0 {
    return 1
  } else {
    return -1
  }
}

func sineWave(freq float64, pos float64) float64 {
  phase_angle := FreqRad * freq * pos
  return math.Sin(phase_angle)
}

func lfo_mod_freq(pos float64, root_freq float64, f func(float64, float64) float64, lfo_freq float64, lfo_amp float64) float64 {
  delta_mod := (sineWave(lfo_freq, pos) * lfo_amp)
  root_freq = root_freq + delta_mod
  return f(root_freq, pos)
}

func lfo_mod_amp(pos float64, root_freq float64, f func(float64, float64) float64, lfo_freq float64, lfo_amp float64) float64 {
  delta_mod := (sineWave(lfo_freq, pos) * lfo_amp)
  return f(root_freq, pos) + (1-delta_mod)
}
