//Euler 10 lets you find the sum of the primes below 2 million

package main

import "fmt"
import "github.com/user/eulerlib"

func main() {

sum:=0
for i:=2;i<2000000;i++ {

if eulerlib.IsaPrime(i){
sum=sum+i
}
}
fmt.Println(sum)
}