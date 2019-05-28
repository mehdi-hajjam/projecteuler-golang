//euler12 finds the first triangular number divisible by at least 500 divisors.
package main

import "fmt"
import "github.com/user/eulerlib"

func main() {

var a=1
var sum=0
var nbdivisor=1

for nbdivisor <= 500 {

var primepowerslice [][]int = eulerlib.PrimePower(sum)
nbdivisor=1
for i:=0;i<len(primepowerslice);i++{
	nbdivisor=nbdivisor*(primepowerslice[i][1]+1)
}
//fmt.Println(sum,nbdivisor)
sum=sum+a
a=a+1
}
fmt.Println(eulerlib.PrimePower(20))
fmt.Println(sum-a+1, nbdivisor)
}