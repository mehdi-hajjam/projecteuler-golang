package eulerlib

//This function calculates the sum of the natural divisors of a number (numbers less than the number which divide evenly into the number).
func SumOfDivisors(n int) int {

//utiliser IsaPrime pour éviter des calculs car un nombre amicable n'est pas premier!
//1 ne rentre pas dans la définition du sujet car a!=b

//peut-être accélérable en faisant que n diminue au fur et à mesure qu'on trouve des diviseurs!
var result []int

for i:=2;i<n;i++{

	if IsaMultipleof(float64(n),float64(i)){
		result=append(result,i)
	}
}

//1 is in the divisors < number
sum:=1

for j:=0;j<len(result);j++{
	sum=sum+result[j]
}

return sum

}