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
    note_buffer := ms.Instrument.Play(note.Freq, note.Duration * hemidemisemiquaver_duration_samples)

    // TODO
    // this should go into some kind of envelope patch class
    // this code does not belong here..

    // linearly advance the first n samples
    attack_length := 800 // in samples
    attack_start := 0
    for i := attack_start; i < attack_length; i++ {
      note_buffer[i] = int16(float64(note_buffer[i]) * (float64(i) / float64(attack_length)))
    }

    // linearly decay the last n samples
    decay_length := 800 // in samples
    decay_start := len(note_buffer) - decay_length
    for i := decay_start; i < len(note_buffer); i++ {
      distance := len(note_buffer) - i 
      note_buffer[i] = int16(float64(note_buffer[i]) * (float64(distance) / float64(decay_length)))
    }

    // end this code does not belong here


    audio = append(audio, note_buffer...)
  }

  return audio

}
