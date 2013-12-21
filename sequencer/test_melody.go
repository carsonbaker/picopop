package sequencer

import (
  "github.com/carsonbaker/copenhagen/instruments"
  "github.com/carsonbaker/copenhagen/util"
)

type note struct {
    freq float64
    duration float64
}

func PlaySweepTest() []int16 {
  instr := instruments.WavetableGenerator{}
  audio := []int16{}
  root := instr.Play2(10, 0.1) // 800 samples long
  
  shifted := instruments.ChangePitch(root, 80)
  for duration := 0; duration < 800; {
    duration += len(audio)
    audio = append(audio, shifted...)
  }
  
  // silence := instruments.SilenceGenerator{}.Play(100)

  // for i := 1.0; i < 32.0; i+= 0.025 {
  //   audio = append(audio, instruments.ChangePitch(root, i)...)
  //   // audio = append(audio, sil...)
  // }

  // audio = root

  util.DebugAudioToFile(audio)
  return audio
}

func PlayTestMelody() []int16 {

  instr := instruments.WavetableGenerator{}

  audio := []int16{}

  // melody := []note{{f2, 0.5}, {g2s, 0.775}, {f2, 0.25}, {f2, 0.125}, {a2s, 0.25}, {f2, 0.25}, {d2s, 0.25}, {f2, 0.5}, {c3, 0.775}, {f2, 0.25}, {f2, 0.125}, {c3s, 0.25}, {c3, 0.25}, {g2s, 0.25}, {f2, 0.25}, {c3, 0.775}, {f3, 0.25}, {f2, 0.125}, {d2s, 0.25}, {d2s, 0.125}, {c2, 0.25}, {g2, 0.25}, {f2, 1.25}}

  melody := []note{{c1, 0.2}, {c2, 0.2}, {c1, 0.2}, {d2s, 0.2}, {f2, 0.2}, {c1, 0.2}, {a2s, 0.2}, {g2, 0.2}, {a0s, 0.2}, {a1s, 0.2}, {a0s, 0.2}, {d2s, 0.2}, {f2, 0.2}, {a0s, 0.2}, {a2s, 0.2}, {g2, 0.2}, {f0, 0.2}, {f1, 0.2}, {f0, 0.2}, {d2s, 0.2}, {f2, 0.2}, {f0, 0.2}, {a2s, 0.2}, {g2, 0.2}, {f0, 0.2}, {f1, 0.2}, {f0, 0.2}, {d2s, 0.2}, {f2, 0.2}, {f0, 0.2}, {a2s, 0.2}, {g2, 0.2}, {c1, 0.2}, {c2, 0.2}, {c1, 0.2}, {c2, 0.2}, {f2, 0.1}, {g2, 0.7}, {g2, 0.2}, {c2, 0.2}, {a0s, 0.2}, {f2, 0.2}, {a0s, 0.2}, {a1s, 0.2}, {d2, 0.4}, {d2, 0.2}, {f2, 0.2}, {f2, 0.2}, {f1, 0.2}, {f0, 0.2}, {f1, 0.2}, {d2, 0.4}, {d2, 0.2}, {f2, 0.2}, {f2, 0.2}, {f1, 0.2}, {c2, 0.2}, {f0, 0.4}, {d2s, 0.4}, {d2, 0.2}, {c2, 0.2}, {c2, 0.2}, {c1, 0.2}, {c0, 0.2}, {c0, 0.2}}

  // melody := []note{{a0s, 0.1}, {a0s, 0.1}, {c1, 0.2}, {c3, 0.05}, {d3s, 0.05}, {g3, 0.05}, {d3s, 0.05}, {c1, 0.2}, {d1s, 0.1}, {d1s, 0.1}, {f1, 0.2}, {a3, 0.05}, {f3, 0.05}, {g3, 0.05}, {c3, 0.05}, {f1, 0.2}, {a0s, 0.1}, {a0s, 0.1}, {c1, 0.2}, {c3, 0.05}, {d3s, 0.05}, {g3, 0.05}, {d3s, 0.05}, {c1, 0.2}, {a0s, 0.1}, {a0s, 0.1}, {c1, 0.2}, {g2, 0.05}, {c3, 0.05}, {d3s, 0.05}, {c3, 0.05}, {c1, 0.2}, {g0s, 0.1}, {g0s, 0.1}, {g1s, 0.2}, {c3, 0.05}, {d3s, 0.05}, {g3s, 0.05}, {d3s, 0.05}, {d2s, 0.2}, {a0s, 0.1}, {a0s, 0.1}, {a1s, 0.2}, {d3, 0.05}, {f3, 0.05}, {g3s, 0.05}, {f3, 0.05}, {d2, 0.2}, {a0s, 0.1}, {a0s, 0.1}, {c1, 0.2}, {c3, 0.05}, {e3, 0.05}, {g3, 0.05}, {e3, 0.05}, {g3, 0.05}, {c4, 0.05}, {g3, 0.05}, {c4, 0.05}, {e5, 0.05}, {c5, 0.05}, {e5, 0.05}, {c1, 0.05}}

  // melody := []note{{c0, 0.7}, {d0, 0.7}, {e0, 0.7}, {f0, 0.7}, {g0, 0.7}, {a0, 0.7}, {b0, 0.7}, {c1, 0.7}, {d1, 0.7}, {e1, 0.7}, {f1, 0.7}, {g1, 0.7}, {a1, 0.7}, {b1, 0.7}, {c2, 0.7}, {d2, 0.7}, {e2, 0.7}, {f2, 0.7}, {g2, 0.7}, {a2, 0.7}, {b2, 0.7}, {c3, 0.7}, {d3, 0.7}, {e3, 0.7}, {f3, 0.7}, {g3, 0.7}, {a3, 0.7}, {b3, 0.7}, {c4, 0.7}, {d4, 0.7}, {e4, 0.7}, {f4, 0.7}, {g4, 0.7}, {a4, 0.7}, {b4, 0.7}, {c5, 0.7}, {d5, 0.7}, {e5, 0.7}, {f5, 0.7}, {g5, 0.7}, {a5, 0.7}, {b5, 0.7}, {c6, 0.7}}


  for _, note := range melody {
    note_buffer := instr.Play2(note.freq*3, note.duration)
    audio = append(audio, note_buffer...)
  }

  util.DebugAudioToFile(audio)

  return audio

}
