package genconv

func MToF(m Meter) Feet { return Feet(m*3.28084)}
func FToM(f Feet) Meter{ return Meter(f*0.3048)}
func PToK(p Pound) Kilogram{ return Kilogram(p*0.453592)}
func KToP(k Kilogram) Pound{return Pound(k*2.20462)}
func LToO(l Litre) Ounce{return Ounce(l*33.814)}
func OToL(o Ounce) Litre{return Litre(o*0.0295735)}
