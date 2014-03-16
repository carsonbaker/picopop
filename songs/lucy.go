package songs

// import(
//   "github.com/carsonbaker/picopop/sequencer"
//   "github.com/carsonbaker/picopop/mixer"
//   "github.com/carsonbaker/picopop/global"
//   "github.com/carsonbaker/picopop/util"
//   "github.com/carsonbaker/picopop/instruments"
// )

// func RenderLucyInTheSkySong() []byte {

//   type UserNote struct {
//     note_symbol string
//     duration int // in hemidemisemiquavers! (1/64 notes)
//   }

//   // user_melody := []UserNote{{"a2", 6}, {"c3", 6}, {"e3", 6}, {"a2", 6}, {"e3", 6}, {"a2", 6}, {"f2s", 6}, {"a2", 6}, {"e3", 6}, {"a2", 6}, {"d3", 3}, {"c3s", 3}, {"a2", 3}}

//   user_melody := []UserNote{{"a2", 6}, {"c3", 6}}

//   matrix_melody := []sequencer.Note{}

//   // convert string notes to float amounts
//   for _, user_note := range user_melody {
//     matrix_note := sequencer.Note{ util.Pitches[user_note.note_symbol]*2, user_note.duration }
//     matrix_melody = append(matrix_melody, matrix_note)
//   }

//   sequencer1 := sequencer.MatrixSequencer{}
//   sequencer1.Melody = matrix_melody
//   sequencer1.Instrument = instruments.SquareGenerator{}

//   pcm := sequencer1.RenderAll()
//   bytes  := mixer.MixToBytes(pcm, global.BytesPerSample)

//   util.DebugAudioToFile(pcm)

//   return bytes

// }
