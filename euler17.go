//Euler17 is just "dénombrement"...

package main

import "fmt"

func main(){

//one+two+three+four+five+six+seven+eight+nine
units:=3+3+5+4+4+3+5+5+4
//ten+eleven+twelve+thirteen+fourteen+fifteen+sixteen+seventeen+eighteen+nineteen
teens:=3+6+6+8+8+7+7+9+8+8
//twenty+thirty+forty+fifty+sixty+seventy+eighty+ninety
dozens:=6+6+5+5+5+7+6+6
hundred:=7
thousand:=8
and:=3
//beware : there is not always and with hundred, hence the -9*3
sum:=10*teens+100*dozens+10*9*units+100*units+900*(hundred+and)-9*3+(3+thousand)

fmt.Println(sum)
}
