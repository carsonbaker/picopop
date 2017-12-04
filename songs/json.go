package songs

import(
  "encoding/json"
  "fmt"
  "../sequencer"
  "../mixer"
  "../global"
  "../util"
  "../instruments"
)

type NoteSequence struct {
  Licks []Lick
}

type Lick struct {
  Tick int
  Note string
  Dura int
  Velo int
  Modd int
}

func RenderJSONSong(message []byte) []byte {

  var sequence NoteSequence
  err := json.Unmarshal(message, &sequence.Licks)

  if err != nil {
    fmt.Println("Oops. Error in parsing JSON.")
    fmt.Println(err)
    fmt.Println("JSON is: ", string(message))
    return nil
  }

  fmt.Println(sequence.Licks)

  matrix_melody := []sequencer.Note{}

  // convert string notes to float amounts
  for _, user_lick := range sequence.Licks {
    fmt.Println(user_lick)
    fmt.Println(user_lick.Note)
    matrix_note := sequencer.Note{ user_lick.Tick, util.Pitches[user_lick.Note], user_lick.Dura }
    matrix_melody = append(matrix_melody, matrix_note)
  }

  sequencer1 := sequencer.MatrixSequencer{}
  sequencer1.Melody = matrix_melody
  sequencer1.Instrument = instruments.WavetableGenerator{}

  pcm := sequencer1.RenderAll()
  bytes  := mixer.MixToBytes(pcm, global.BytesPerSample)

  util.DebugAudioToFile(pcm)

  return bytes

}
