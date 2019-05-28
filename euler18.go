//Euler18 is searching for the  best route in a tree. Treated so as to be able to deal with Euler67 easily (no brute force used)

package main

import "github.com/user/eulerlib"
import "fmt"
import "io/ioutil"
import "strings"
import "math"

func main() {

//Error cathching is missing. b is []bytes
b, _ := ioutil.ReadFile("goworkspace/euler18_data.txt")
//string() converts into []string, and strings.NewReader creates an ioreader needed for ReadInts to work
s,_ := eulerlib.ReadInts(strings.NewReader(string(b)))
fmt.Println(s)
fmt.Println(len(s))
offset:=14

for i:=14;i>=0;i--{
	offset=offset+i
	for j:=0;j<i;j++{
		s[len(s)-1-offset+j]=int(math.Max(float64(s[len(s)-1-offset+j]+s[len(s)-1-offset+j+i]),float64(s[len(s)-1-offset+j]+s[len(s)-1-offset+j+i+1])))
		
	}
}
fmt.Println(s)
}