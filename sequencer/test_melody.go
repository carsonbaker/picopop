package sequencer

import (
  "github.com/carsonbaker/copenhagen/instruments"
  "github.com/carsonbaker/copenhagen/util"
)

const c0  = 16.35
const c0s = 17.32
const d0  = 18.35
const d0s = 19.45
const e0  = 20.60
const f0  = 21.83
const f0s = 23.12
const g0  = 24.50
const g0s = 25.96
const a0  = 27.50
const a0s = 29.14
const b0  = 30.87
const c1  = 32.70
const c1s = 34.65
const d1  = 36.71
const d1s = 38.89
const e1  = 41.20
const f1  = 43.65
const f1s = 46.25
const g1  = 49.00
const g1s = 51.91
const a1  = 55.00
const a1s = 58.27
const b1  = 61.74
const c2  = 65.41
const c2s = 69.30
const d2  = 73.42
const d2s = 77.79
const e2  = 82.41
const f2  = 87.31
const f2s = 92.50
const g2  = 98.00
const g2s = 103.83
const a2  = 110.00
const a2s = 116.54
const b2  = 123.47
const c3  = 130.81
const c3s = 138.59
const d3  = 146.83
const d3s = 155.56
const e3  = 164.81
const f3  = 174.61
const f3s = 185.00
const g3  = 196.00
const g3s = 207.65
const a3  = 220.00
const a3s = 233.08
const b3  = 246.94
const c4  = 261.63
const c4s = 277.18
const d4  = 293.66
const d4s = 311.13
const e4  = 329.63
const f4  = 349.23
const f4s = 369.99
const g4  = 392.00
const g4s = 415.30
const a4  = 440.00
const a4s = 466.16
const b4  = 493.88
const c5  = 523.25
const c5s = 554.37
const d5  = 587.33
const d5s = 622.25
const e5  = 659.26
const f5  = 698.46
const f5s = 739.99
const g5  = 783.99
const g5s = 830.61
const a5  = 880.00
const a5s = 932.33
const b5  = 987.77
const c6  = 1046.50
const c6s = 1108.73
const d6  = 1174.66
const d6s = 1244.51
const e6  = 1318.51
const f6  = 1396.91
const f6s = 1479.98

type note struct {
    freq float64
    duration float64
}

type AxelFMelody struct { }

func (AxelFMelody) Play() []byte {

  instr := instruments.SquareGenerator{}

  audio := []byte{}

  // melody := []note{{f2, 0.5}, {g2s, 0.775}, {f2, 0.25}, {f2, 0.125}, {a2s, 0.25}, {f2, 0.25}, {d2s, 0.25}, {f2, 0.5}, {c3, 0.775}, {f2, 0.25}, {f2, 0.125}, {c3s, 0.25}, {c3, 0.25}, {g2s, 0.25}, {f2, 0.25}, {c3, 0.775}, {f3, 0.25}, {f2, 0.125}, {d2s, 0.25}, {d2s, 0.125}, {c2, 0.25}, {g2, 0.25}, {f2, 1.25}}

  melody := []note{{c1, 0.2}, {c2, 0.2}, {c1, 0.2}, {d2s, 0.2}, {f2, 0.2}, {c1, 0.2}, {a2s, 0.2}, {g2, 0.2}, {a0s, 0.2}, {a1s, 0.2}, {a0s, 0.2}, {d2s, 0.2}, {f2, 0.2}, {a0s, 0.2}, {a2s, 0.2}, {g2, 0.2}, {f0, 0.2}, {f1, 0.2}, {f0, 0.2}, {d2s, 0.2}, {f2, 0.2}, {f0, 0.2}, {a2s, 0.2}, {g2, 0.2}, {f0, 0.2}, {f1, 0.2}, {f0, 0.2}, {d2s, 0.2}, {f2, 0.2}, {f0, 0.2}, {a2s, 0.2}, {g2, 0.2}, {c1, 0.2}, {c2, 0.2}, {c1, 0.2}, {c2, 0.2}, {f2, 0.1}, {g2, 0.7}, {g2, 0.2}, {c2, 0.2}, {a0s, 0.2}, {f2, 0.2}, {a0s, 0.2}, {a1s, 0.2}, {d2, 0.4}, {d2, 0.2}, {f2, 0.2}, {f2, 0.2}, {f1, 0.2}, {f0, 0.2}, {f1, 0.2}, {d2, 0.4}, {d2, 0.2}, {f2, 0.2}, {f2, 0.2}, {f1, 0.2}, {c2, 0.2}, {f0, 0.4}, {d2s, 0.4}, {d2, 0.2}, {c2, 0.2}, {c2, 0.2}, {c1, 0.2}, {c0, 0.2}, {c0, 0.2}}

  // melody := []note{{a0s, 0.1}, {a0s, 0.1}, {c1, 0.2}, {c3, 0.05}, {d3s, 0.05}, {g3, 0.05}, {d3s, 0.05}, {c1, 0.2}, {d1s, 0.1}, {d1s, 0.1}, {f1, 0.2}, {a3, 0.05}, {f3, 0.05}, {g3, 0.05}, {c3, 0.05}, {f1, 0.2}, {a0s, 0.1}, {a0s, 0.1}, {c1, 0.2}, {c3, 0.05}, {d3s, 0.05}, {g3, 0.05}, {d3s, 0.05}, {c1, 0.2}, {a0s, 0.1}, {a0s, 0.1}, {c1, 0.2}, {g2, 0.05}, {c3, 0.05}, {d3s, 0.05}, {c3, 0.05}, {c1, 0.2}, {g0s, 0.1}, {g0s, 0.1}, {g1s, 0.2}, {c3, 0.05}, {d3s, 0.05}, {g3s, 0.05}, {d3s, 0.05}, {d2s, 0.2}, {a0s, 0.1}, {a0s, 0.1}, {a1s, 0.2}, {d3, 0.05}, {f3, 0.05}, {g3s, 0.05}, {f3, 0.05}, {d2, 0.2}, {a0s, 0.1}, {a0s, 0.1}, {c1, 0.2}, {c3, 0.05}, {e3, 0.05}, {g3, 0.05}, {e3, 0.05}, {g3, 0.05}, {c4, 0.05}, {g3, 0.05}, {c4, 0.05}, {e5, 0.05}, {c5, 0.05}, {e5, 0.05}, {c1, 0.05}}

  // melody := []note{{c0, 0.7}, {d0, 0.7}, {e0, 0.7}, {f0, 0.7}, {g0, 0.7}, {a0, 0.7}, {b0, 0.7}, {c1, 0.7}, {d1, 0.7}, {e1, 0.7}, {f1, 0.7}, {g1, 0.7}, {a1, 0.7}, {b1, 0.7}, {c2, 0.7}, {d2, 0.7}, {e2, 0.7}, {f2, 0.7}, {g2, 0.7}, {a2, 0.7}, {b2, 0.7}, {c3, 0.7}, {d3, 0.7}, {e3, 0.7}, {f3, 0.7}, {g3, 0.7}, {a3, 0.7}, {b3, 0.7}, {c4, 0.7}, {d4, 0.7}, {e4, 0.7}, {f4, 0.7}, {g4, 0.7}, {a4, 0.7}, {b4, 0.7}, {c5, 0.7}, {d5, 0.7}, {e5, 0.7}, {f5, 0.7}, {g5, 0.7}, {a5, 0.7}, {b5, 0.7}, {c6, 0.7}}

  for _, note := range melody {
    note_buffer := instr.Play(note.freq*3, note.duration)
    audio = append(audio, note_buffer...)
  }

  util.DebugAudioToFile(audio)

  return audio

}
