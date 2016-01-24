package genconv
import "fmt"

type Meter float64
type Feet float64
type Pound float64
type Kilogram float64
type Litre float64
type Ounce float64

func (m Meter) String() string { return fmt.Sprintf("%g[m]", m) }
func (f Feet) String() string { return fmt.Sprintf("%g[ft]", f) }
func (p Pound) String() string { return fmt.Sprintf("%g[lb]", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g[kg]", k) }
func (l Litre) String() string { return fmt.Sprintf("%g[l]", l) }
func (o Ounce) String() string { return fmt.Sprintf("%g[oz]", o) }
