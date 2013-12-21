package global

import ("math")

const Samplerate      = 8000
const Channels        = 2
const TwoPi           = math.Pi * 2
const BitDepth        = 16
const BytesPerSample  = Channels * BitDepth / 8
const FreqRad         = TwoPi / Samplerate
