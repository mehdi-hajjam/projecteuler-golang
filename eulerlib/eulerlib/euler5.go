package eulerlib

import "math"
//import "fmt"

//This function gives the prime decomposition under the following form : Primepower (14)=[[7 1][2 1]] ([prime power_of_prime])
func PrimePower (x int) [][]int {

var a [][]int
var b int = x
//for i:=2;i<=int(math.Sqrt(float64(x)));i++{
for i:=2;i<=b;i++{
var p int = 0
var j int= i
for IsaMultipleof(float64(b),float64(j)){
p=p+1
j=j*i
}
if p>0 {
a=append(a,[]int{i,p})
b=b/int(math.Pow(float64(i),float64(p)))
//fmt.Println("b:",b)
}
}
return a


}

//This function tells whether x is a prime number or not
//It has been improved to run faster to meet the 60sec constraint in Project Euler for numbers until 2 million.
func IsaPrime (x int) bool {
var rep bool=true
if IsaMultipleof(float64(x),float64(2))==true && x>2 {
rep=false
}
if IsaMultipleof(float64(x),float64(3))==true && x>3 {
rep=false
}
for i:=1;i<=int(math.Floor(math.Sqrt(float64(x))+6))/6 && rep==true;i++{
if IsaMultipleof(float64(x),float64(6*i+1))==true && x>7{
rep=false
}
if IsaMultipleof(float64(x),float64(6*i-1))==true && x>5 {
rep=false
}
}
return rep
}