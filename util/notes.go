package util

const C0  = 16.35
const C0s = 17.32
const D0  = 18.35
const D0s = 19.45
const E0  = 20.60
const F0  = 21.83
const F0s = 23.12
const G0  = 24.50
const G0s = 25.96
const A0  = 27.50
const A0s = 29.14
const B0  = 30.87
const C1  = 32.70
const C1s = 34.65
const D1  = 36.71
const D1s = 38.89
const E1  = 41.20
const F1  = 43.65
const F1s = 46.25
const G1  = 49.00
const G1s = 51.91
const A1  = 55.00
const A1s = 58.27
const B1  = 61.74
const C2  = 65.41
const C2s = 69.30
const D2  = 73.42
const D2s = 77.79
const E2  = 82.41
const F2  = 87.31
const F2s = 92.50
const G2  = 98.00
const G2s = 103.83
const A2  = 110.00
const A2s = 116.54
const B2  = 123.47
const C3  = 130.81
const C3s = 138.59
const D3  = 146.83
const D3s = 155.56
const E3  = 164.81
const F3  = 174.61
const F3s = 185.00
const G3  = 196.00
const G3s = 207.65
const A3  = 220.00
const A3s = 233.08
const B3  = 246.94
const C4  = 261.63
const C4s = 277.18
const D4  = 293.66
const D4s = 311.13
const E4  = 329.63
const F4  = 349.23
const F4s = 369.99
const G4  = 392.00
const G4s = 415.30
const A4  = 440.00
const A4s = 466.16
const B4  = 493.88
const C5  = 523.25
const C5s = 554.37
const D5  = 587.33
const D5s = 622.25
const E5  = 659.26
const F5  = 698.46
const F5s = 739.99
const G5  = 783.99
const G5s = 830.61
const A5  = 880.00
const A5s = 932.33
const B5  = 987.77
const C6  = 1046.50
const C6s = 1108.73
const D6  = 1174.66
const D6s = 1244.51
const E6  = 1318.51
const F6  = 1396.91
const F6s = 1479.98

var Pitches map[string]float64
 
func InitPitches() {
  Pitches = map[string]float64 {
    "c0"  : C0,
    "c0s" : C0s,
    "d0"  : D0,
    "d0s" : D0s,
    "e0"  : E0,
    "f0"  : F0,
    "f0s" : F0s,
    "g0"  : G0,
    "g0s" : G0s,
    "a0"  : A0,
    "a0s" : A0s,
    "b0"  : B0,
    "c1"  : C1,
    "c1s" : C1s,
    "d1"  : D1,
    "d1s" : D1s,
    "e1"  : E1,
    "f1"  : F1,
    "f1s" : F1s,
    "g1"  : G1,
    "g1s" : G1s,
    "a1"  : A1,
    "a1s" : A1s,
    "b1"  : B1,
    "c2"  : C2,
    "c2s" : C2s,
    "d2"  : D2,
    "d2s" : D2s,
    "e2"  : E2,
    "f2"  : F2,
    "f2s" : F2s,
    "g2"  : G2,
    "g2s" : G2s,
    "a2"  : A2,
    "a2s" : A2s,
    "b2"  : B2,
    "c3"  : C3,
    "c3s" : C3s,
    "d3"  : D3,
    "d3s" : D3s,
    "e3"  : E3,
    "f3"  : F3,
    "f3s" : F3s,
    "g3"  : G3,
    "g3s" : G3s,
    "a3"  : A3,
    "a3s" : A3s,
    "b3"  : B3,
    "c4"  : C4,
    "c4s" : C4s,
    "d4"  : D4,
    "d4s" : D4s,
    "e4"  : E4,
    "f4"  : F4,
    "f4s" : F4s,
    "g4"  : G4,
    "g4s" : G4s,
    "a4"  : A4,
    "a4s" : A4s,
    "b4"  : B4,
    "c5"  : C5,
    "c5s" : C5s,
    "d5"  : D5,
    "d5s" : D5s,
    "e5"  : E5,
    "f5"  : F5,
    "f5s" : F5s,
    "g5"  : G5,
    "g5s" : G5s,
    "a5"  : A5,
    "a5s" : A5s,
    "b5"  : B5,
    "c6"  : C6,
    "c6s" : C6s,
    "d6"  : D6,
    "d6s" : D6s,
    "e6"  : E6,
    "f6"  : F6,
    "f6s" : F6s,
  }
}

func Initialize() {
  InitPitches()  
}
