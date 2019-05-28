//Euler20 is strongly based on Euler16

package main

import "fmt"
import "math/big"

func main() {

a:=big.NewInt(1)
b:=big.NewInt(1)
c:=big.NewInt(9)
d:=big.NewInt(10)
e:=big.NewInt(1)
r:=big.NewInt(0)
sum:=big.NewInt(0)

for i:=1;i<100;i++{
b.Add(b,e)
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
