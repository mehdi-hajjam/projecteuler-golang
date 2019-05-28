//Euler19 goal is to calcule how many 1st of month were a Sunday from 1 jan 1901 to 31 dec 2000


package main

import "fmt"
import "github.com/user/eulerlib"


func main(){

result:=0
for i:=1901;i<=2000;i++{
	for j:=1;j<=12;j++{
//01 JAN 1900 is a Monday, so modulo 7 = 0 means it was a Sunday
		if eulerlib.IsaMultipleof(float64(HowManyDays(j,i)),float64(7)){
			result=result+1
		}
		}
	}

fmt.Println(result)
}


//This function returns how many 29FEB happened between 01 JAN 1900 (cause 1900 does not have one, hence the i=1901) and before month month of year year
//Simplified version because 2000 is divisible by 4 and by 400
func HowMany29Feb(month, year int) int {
febint:=0
num:=0
if month >= 3 {
	num=1
}
for i:=1901;i<year+num;i++{
	if eulerlib.IsaMultipleof(float64(i),float64(4)){
		febint=febint+1
	}
}

return febint
}

//This function gives the normal of non 29FEB days between 01 JAN 1900 and month month of year year
func HowManyDays(month, year int) int {

total:=(year-1900)*365+HowMany29Feb(month,year)+1

for j:=1;j<=month;j++{
		switch{
		case j==1:
			total=total+0
		case j==2:
			total=total+31
		case j==3:
			total=total+28
		case j==4:
			total=total+31
		case j==5:
			total=total+30
		case j==6:
			total=total+31
		case j==7:
			total=total+30
		case j==8:
			total=total+31
		case j==9:
			total=total+31
		case j==10:
			total=total+30
		case j==11:
			total=total+31
		case j==12:
			total=total+30
		}
}
return total
}