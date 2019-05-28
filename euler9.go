//Euler 9 finds a special Pythagorean triplet  !!

package main

import "fmt"
import "math"

func main() {

a:=1
b:=1
c:=1

for i:=1;i<=1000;i++ {
for j:=i+1;j<=1000;j++ {

k:=math.Sqrt(float64(i*i+j*j))

if i+j+int(k)==1000 && k==math.Floor(k){
a=i
b=j
c=int(math.Sqrt(float64(a*a+b*b)))
}

}


}
product:=a*b*c
fmt.Println(a, b, c, product)
}