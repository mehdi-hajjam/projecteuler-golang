//Euler 20

package main

import "github.com/user/eulerlib"
import "fmt"

func main(){
var amicable []int	
var result []int

for i:=0;i<10001;i++{
result=append(result,eulerlib.SumOfDivisors(i+1))
}
for j:=0;j<len(result);j++{
	if result[j]<len(result){
	if result[result[j]-1]==j+1 && result[j]!=j+1{
		amicable=append(amicable,j+1)
		amicable=append(amicable,result[j])
	}
}
}
fmt.Println(amicable)
sum:=0
for k:=0;k<len(amicable);k++{
sum=sum+amicable[k]
}
fmt.Println(sum/2)
}