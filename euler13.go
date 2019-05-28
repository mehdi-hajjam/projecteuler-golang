//Euler13 reads a file and sums the number it finds on each line
//I had to use the Math/big package to deal with the full 50 digit figures. Maybe could have been solved otherwise since only 10 first digit asked for.
package main

import (
  "bufio"
  "fmt"
  "os"
  //"strconv"
  "math/big"
)

//import "log"

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func main() {

pwd, _:=os.Getwd()	
var path string = pwd+"/goworkspace/euler13_data.txt"
r, _ :=readLines(path)
sum:=new(big.Int)
for i:=0;i<len(r);i++{
bigint:=new(big.Int)
bigint.SetString(r[i],10)
sum.Add(sum,bigint)
}
fmt.Println(sum)
}