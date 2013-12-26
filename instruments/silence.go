package instruments

type SilenceGenerator struct {

}

func (s SilenceGenerator) Play(duration int) []int16 {
  generator := func(pos int) int16 { return 0 }
  return RunBuffer(generator, duration)
}
