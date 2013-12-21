package songs

import(
  "github.com/carsonbaker/copenhagen/sequencer"
  "github.com/carsonbaker/copenhagen/mixer"
  "github.com/carsonbaker/copenhagen/global"
)

func RenderBeefyFartSong() []byte {
  pcm := sequencer.PlaySweepTest()
  bytes  := mixer.MixToBytes(pcm, global.BytesPerSample)
  return bytes
}
