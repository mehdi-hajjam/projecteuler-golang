//Euler 5 goal is to find the smallest multiple evenly divisible by numbers from 1 to 20

package main

import "fmt"
import "math"
import "github.com/user/eulerlib"


func main() {
var a [][]int
var result [][]int
var temp []int
//stocke dans a la decomposition en nombre premier des nombres de 1 a 20
for i:=1;i<=20;i++{
a=append(a,eulerlib.PrimePower(i)...)
}
for j:=0;j<len(a);j++{
temp=a[j]
for i:=j+1;i<len(a);i++{
switch {
case temp[0]==a[i][0]:
if temp[1]<a[i][1]{
temp=a[i]
}
//on supprime les puissances de nombre premiers non unique et non maximales pour ne pas que result contienne des doublons
a=append(a[:i],a[i+1:]...)
i--
}
}
//on stocke dans result uniquement les puissances maximales des nombres premiers
result=append(result,temp)
}
var res int = 1
for l:=0;l<len(result);l++{
res=res * int(math.Pow(float64(result[l][0]),float64(result[l][1])))
}
//fmt.Println(result)
fmt.Println(res)
fmt.Println(eulerlib.IsaPrime(10))
}