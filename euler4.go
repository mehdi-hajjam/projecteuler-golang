//Euler 4 looks for the maximum palindromic number that is the product of two three digit numbers

package main

import "fmt"
import "github.com/user/eulerlib"
//import "math"

func main() {

var results []int
var a []int
for i:=999;i>99;i=i-1{
for j:=999;j>99;j=j-1{
a=eulerlib.Basey(i*j,10)
if eulerlib.IsEqualTo(a,eulerlib.Reverse(a)) {
results=append(results,i*j)
}
}
}
fmt.Println("done")
fmt.Println(eulerlib.MaxSlice(results))


//a:=eulerlib.Basey(10100,10)
//fmt.Println(a)
//b:=eulerlib.Reverse(a)
//fmt.Println(b)
//fmt.Println(eulerlib.IsEqualTo(a,b))


}