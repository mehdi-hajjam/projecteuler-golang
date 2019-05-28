package eulerlib

import "math"

func IsaMultipleof(x,y float64) bool{
     var a bool = false
     if math.Remainder(x,y)==0.0 {
        a=true
     }
     return a
}