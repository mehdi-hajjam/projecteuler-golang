package eulerlib

func MaxSlice (a []int) int {
max:=a[0]
for i:=1;i<len(a);i++{
if a[i]>max {
max=a[i]
}
}

return max
}