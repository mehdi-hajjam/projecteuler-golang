package eulerlib

//This function calculates the sum of the squares up to n included

func SumSquares(n int) int {

var result int = 0

for i:=1;i<=n;i++{
result=result+i*i
}
return result
}






//This function calculates the square of the sum up to n included

func SquareSum(n int) int {

var result int = 0

for i:=1;i<=n;i++{
result=result+i
}
return result*result
}
