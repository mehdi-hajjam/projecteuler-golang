//Euler16 made me realize that big.Int type is a struct. Interesting to use methods on this type. No use for slices to solve this.

package main

import "fmt"
import "math/big"

func main() {

a:=big.NewInt(1)
b:=big.NewInt(2)
c:=big.NewInt(9)
d:=big.NewInt(10)
r:=big.NewInt(0)
sum:=big.NewInt(0)

for i:=1;i<1001;i++{
a.Mul(a,b)
}


for a.Cmp(c)>0 {
r.Rem(a,d)	
a.Quo(a,d)
fmt.Println(r)
sum.Add(sum,r)
}
sum.Add(sum,a)
fmt.Println(sum)
}
