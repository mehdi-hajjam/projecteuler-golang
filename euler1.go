//Euler 1 is only tricky because "below 1000" means strictly below!

package main

import "fmt"
import "github.com/user/eulerlib"

func main() {

var sum int = 0
for i:=0;i<1000;i++{
switch{
case eulerlib.IsaMultipleof(float64(i),3):
sum=sum+i
case eulerlib.IsaMultipleof(float64(i),5):
sum=sum+i
}
}
fmt.Println(sum)
}