package lenghtconv

func MToF(m Metre) Feet { return Feet(m / 0.308) }
func FToM(f Feet) Metre { return Metre(f * 0.308) }
