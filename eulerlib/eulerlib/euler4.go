package eulerlib

import "math"

//This function stores in reverse in a slice the base y decomposition of number x
func Basey(x,y int) []int {
var a []int
var b int = x
var rest float64
for math.Floor(float64(b)) > 0.0 {
rest=math.Remainder(float64(b),float64(y))
switch {
case rest<0:
a=append(a,y+int(rest))
case rest>=0:
a=append(a,int(rest))
}
b=b/y
}
return a
}

//This function operates a permutation leading to the first being the last, the second being one before last etc : it's a mirror
//See https://gobyexample.com/slices to understand that if you do not use make then your slice length is going to be 0 and everything goes panic
func Reverse(x []int) []int {

var a =make([]int,len(x))
var length int = len(x)
for i:=1;i<=length;i++{
a[i-1]=x[length-i]
}
return a
}

//This function compares two slices and returns true if they are equal
func IsEqualTo(a, b []int) bool {

var equal bool=true
var i int=0
for equal==true && i<len(a) {
equal=(a[i]==b[i])
i=i+1
}
return equal
}


