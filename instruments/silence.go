package instruments

type SilenceGenerator struct {
  SimpleParent
}

func (s SilenceGenerator) Play(duration float64) []int16 {
  generator := func(pos int) int16 { return 0 }
  return s.RunBuffer(generator, duration)
}
