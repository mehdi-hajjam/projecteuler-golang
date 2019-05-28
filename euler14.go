//Euler 14 could be solved while not assuming Collatz suite does get to 1, but it was easier to assume they do.
//Pretty sure the code of the Collatz function could be optimized, or that of the main function. For instance,
//not trying the number that are included in the previous collatz suits could help, cause clearly they lead to shorter
//sequences... Right now it runs in 54s, under the 60s target

package main

import "github.com/user/eulerlib"
import "fmt"

func main(){
var s []int
for i:=2;i<1000001;i++{
s=append(s,eulerlib.Collatz(i,0))
}
r, _:=eulerlib.MaxSlice2(s)
fmt.Println(r+2)
}
