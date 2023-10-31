package weightconv

func KgToLb(k Kilo) Pound  { return Pound(k / 0.453) }
func LbToKg(lb Pound) Kilo { return Kilo(lb * 0.453) }
