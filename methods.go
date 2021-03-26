package main

type cel float64
type kel float64
type fah float64

func (c cel) kel() kel {
	return kel(c + 273.15)
}
func (c cel) fah() fah {
	return fah((c * 9.0 / 5.0) + 32.0)
}
func (k kel) cel() cel {
	return cel(k - 273.15)
}
func (k kel) fah() fah {
	return k.cel().fah()
}
func (f fah) cel() cel {
	return cel((f - 32.0) * 5.0 / 9.0)
}
func (f fah) kel() kel {
	return f.cel().kel()
}
