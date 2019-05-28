//Euler 3 made me use a slice. Also, I had to implement the maximum of a slice in eulerlib. Answer is quite small compared to the number!

package main

import "fmt"
import "github.com/user/eulerlib"

var a []int
var b int = 600851475143

func main(){
for i:=1;i<=b;i++{
if eulerlib.IsaMultipleof(float64(b),float64(i)) {
a=append(a,i)
b=b/i
}
}
fmt.Println(eulerlib.MaxSlice(a))
}