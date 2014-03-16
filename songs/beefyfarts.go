package songs

// import(
//   "github.com/carsonbaker/picopop/sequencer"
//   "github.com/carsonbaker/picopop/mixer"
//   "github.com/carsonbaker/picopop/global"
//   "github.com/carsonbaker/picopop/util"
//   "github.com/carsonbaker/picopop/instruments"
// )

// func RenderBeefyFartSong() []byte {

//   type UserNote struct {
//     note_symbol string
//     duration int // in hemidemisemiquavers! (1/64 notes)
//   }

//   user_melody := []UserNote{{"c1", 2}, {"c2", 2}, {"c1", 2}, {"d2s", 2}, {"f2", 2}, {"c1", 2}, {"a2s", 2}, {"g2", 2}, {"a0s", 2}, {"a1s", 2}, {"a0s", 2}, {"d2s", 2}, {"f2", 2}, {"a0s", 2}, {"a2s", 2}, {"g2", 2}, {"f0", 2}, {"f1", 2}, {"f0", 2}, {"d2s", 2}, {"f2", 2}, {"f0", 2}, {"a2s", 2}, {"g2", 2}, {"f0", 2}, {"f1", 2}, {"f0", 2}, {"d2s", 2}, {"f2", 2}, {"f0", 2}, {"a2s", 2}, {"g2", 2}, {"c1", 2}, {"c2", 2}, {"c1", 2}, {"c2", 2}, {"f2", 6}, {"g2", 8}, {"g2", 2}, {"c2", 2}, {"a0s", 2}, {"f2", 2}, {"a0s", 2}, {"a1s", 2}, {"d2", 4}, {"d2", 2}, {"f2", 2}, {"f2", 2}, {"f1", 2}, {"f0", 2}, {"f1", 2}, {"d2", 4}, {"d2", 2}, {"f2", 2}, {"f2", 2}, {"f1", 2}, {"c2", 2}, {"f0", 4}, {"d2s", 4}, {"d2", 2}, {"c2", 2}, {"c2", 2}, {"c1", 2}, {"c0", 2}, {"c0", 2}}

//   matrix_melody := []sequencer.Note{}

//   // convert string notes to float amounts
//   for _, user_note := range user_melody {
//     matrix_note := sequencer.Note{ util.Pitches[user_note.note_symbol] * 4 , user_note.duration }
//     matrix_melody = append(matrix_melody, matrix_note)
//   }

//   sequencer1 := sequencer.MatrixSequencer{}
//   sequencer1.Melody = matrix_melody
//   sequencer1.Instrument = instruments.WavetableGenerator{}

//   pcm := sequencer1.RenderAll()
//   bytes  := mixer.MixToBytes(pcm, global.BytesPerSample)

//   util.DebugAudioToFile(pcm)

//   return bytes

// }
