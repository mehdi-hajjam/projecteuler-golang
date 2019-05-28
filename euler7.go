//Euler 7 lets you find the 10001st prime number. Runs in 76 secs though. To be improved.
package main

import "github.com/user/eulerlib"
import "fmt"

func main() {
var notyet bool=true
var compteur int=0
var i int = 2

//fmt.Println(eulerlib.IsaPrime(4))

for notyet==true {
if eulerlib.IsaPrime(i){
compteur=compteur+1
fmt.Println(compteur,i)
}
if compteur==10001{
notyet=false
}
i=i+1
}
fmt.Println("Voici le 10001e nombre premier:", i-1)
}