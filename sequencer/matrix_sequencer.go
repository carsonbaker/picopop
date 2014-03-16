package sequencer

import (
  "github.com/carsonbaker/picopop/instruments"
  "github.com/carsonbaker/picopop/global"

)

type Note struct {
  Tick int // in hemidemisemiquavers! (1/64 notes)
  Freq float64
  Duration int // in hemidemisemiquavers
}

type MatrixSequencer struct {
  Melody []Note
  Instrument instruments.Instrument
}

const tempo = 80.0
const hemidemisemiquaver_duration = 7.5 / tempo
const hemidemisemiquaver_duration_samples = int(hemidemisemiquaver_duration * global.Samplerate)

// returns the end time of the last note in the sequence
// in units of hemidemisemiquavers
func (ms MatrixSequencer) Length() int {
  length := 0
  for _, note := range ms.Melody {
    length += note.Duration
  }
  return length
}

// returns the end time of the last note in the sequence
// in units of samples
func (ms MatrixSequencer) LengthSamples() int {
  return ms.Length() * hemidemisemiquaver_duration_samples
}

func(ms MatrixSequencer) TicksAtEndOfMelody() int {
  longest := 0
  for _, note := range ms.Melody {
    candidate := note.Duration + note.Tick
    if candidate > longest {
      longest = candidate
    }
  }
  return longest
}


func (ms MatrixSequencer) RenderAll() []int16 {

  audio_length := ms.TicksAtEndOfMelody()
  audio := make([]int16, audio_length * hemidemisemiquaver_duration_samples)

  for _, note := range ms.Melody {
    // render the instrument noise
    note_buffer := ms.Instrument.Play(note.Freq, note.Duration * hemidemisemiquaver_duration_samples)
    // audio = append(audio, note_buffer...)

    base_i := hemidemisemiquaver_duration_samples * note.Tick
    for i := 0; i < len(note_buffer); i++ {
      audio[base_i + i] += note_buffer[i] / 3
    }

  }

  return audio

}
