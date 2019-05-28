//Euler15 finds the number of roads in a 20x20 square grid. You start in the top left corner and need to reach the bottom right corner.
//The only direction you can take is down or right.
//The trick lies in understanding that for every intersection (apart from the arrival point, which has a number of roads equal to the
//beginning one), the number of roads is the sum of the number of roads in the intersections down and to the right of it
//Also noticeable is the obvious symmetry of the problem
//Beware : a 20x20 grid means a 21x21 intersection board!
package main

import "fmt"

func main() {

//creation of the matrix of intersection

grid := make([][]int, 21)
for i:=0;i<21;i++{
	s:=make([]int, 21)
	grid[i]=s
}

//initialize the obvious values to 1 (where no other road but one is possible)

for i:=0;i<20;i++{
		grid[20][i]=1
		grid[i][20]=1
	}

for i:=0;i<21;i++{
fmt.Println(grid[i])
}

//calculate the values for the upper triangle (column number >= line number), deduce the value for the symmetric value
for j:=19;j>=0;j--{
	for i:=j;i>=0;i--{
		grid[i][j]=grid[i+1][j]+grid[i][j+1]
		if i!=j {
		grid[j][i]=grid[i][j]
		}

	}
}
for i:=0;i<21;i++{
fmt.Println(grid[i])
}
}