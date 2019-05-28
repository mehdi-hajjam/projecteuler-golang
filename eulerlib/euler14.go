package  eulerlib

//This functions gives back the length of the Collatz suite for a number a if j=1 (includes the final 1 number)
func Collatz(a,j int) int {

if a>1{
switch {
case  IsaMultipleof(float64(a),float64(2)):
j=Collatz(a/2,j+1)

case  !IsaMultipleof(float64(a),float64(2)):
j=Collatz(3*a+1,j+1)
}
}
return j
}

//This functions returns the index in the slice of the largest element as well as its value
func MaxSlice2 (a []int) (int, int) {
j:=0
max:=a[0]
for i:=1;i<len(a);i++{
if a[i]>max {
max=a[i]
j=i
}
}

return j, max
}