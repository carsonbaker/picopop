package sequencer
// 
// import (
//   "github.com/carsonbaker/copenhagen/instruments"
//   "github.com/carsonbaker/copenhagen/util"
//   "github.com/carsonbaker/copenhagen/global"
// )
// 
// type note struct {
//     start_time int
//     freq float64
//     duration float64
// }
// 
// type MatrixSequencer struct { }
// 
// func (MatrixSequencer) Play() []byte {
// 
//   melody := []note{{0, 200, 0.5}, {4000, 250, 0.7}, {8000, 300, 0.7}}
// 
//   instr := instruments.SquareGenerator{}
// 
//   song_length_secs = 2
//   song_length := song_length_secs * Samplerate
//   sequence_csr := 0
// 
//   audio := make([]int16, song_length)
// 
//   for time_step := 0; time_step < song_length; time_step++ {
// 
//     for {
//       note = melody[sequence_csr]
//       if note.start_time <= time_step {
//         // composite to buffer
//         instr.Play(note.freq, note.duration)
//         audio[time_step] += 
//         sequence_csr++ // perhaps next note is at the same time?
//       } else {
//         break
//       }
//     }
// 
//   }
// 
//   for _, note := range melody {
//     note_buffer := instr.Play(note.freq*3, note.duration)
//     audio = append(audio, note_buffer...)
//   }
// 
//   util.DebugAudioToFile(audio)
// 
//   return audio
// 
// }
