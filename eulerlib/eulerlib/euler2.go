package eulerlib

//This function returns the xth Fibonnaci number
func Fibonacci(x int) int {
var a int
switch{
       case x==0:
       a=0
       case x==1:
       a=1
       case x>1:
       a=Fibonacci(x-1)+Fibonacci(x-2)
}
return a
}