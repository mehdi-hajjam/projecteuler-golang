//Euler 2 could maybe improved in terms of execution time if the lib were to use concurrency

package main

import "github.com/user/eulerlib"
import "fmt"

func main(){
var sum int=0
var b int=0
for i:=0;b<4000000;i++ {
a:=eulerlib.Fibonacci(i)
if a<4000000 && eulerlib.IsaMultipleof(float64(a),2.0) {
sum=sum+a
}
b=a
}
fmt.Println(sum)
}