package sequencer

import (
  "github.com/carsonbaker/copenhagen/instruments"
  "github.com/carsonbaker/copenhagen/global"
)

type Note struct {
    Freq float64
    Duration int // in hemidemisemiquavers! (1/64 notes)
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

func (ms MatrixSequencer) RenderAll() []int16 {

  audio := []int16{}

  for _, note := range ms.Melody {
    // render the instrument noise
    note_buffer := ms.Instrument.Play(note.Freq, note.Duration * hemidemisemiquaver_duration_samples)
    audio = append(audio, note_buffer...)
  }

  return audio

}
