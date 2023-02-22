// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

var frase string

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			{
				frase = "Pin Pan"
			}
		case i%3 == 0:
			{
				frase = "Pin"
			}
		case i%5 == 0:
			{
				frase = "Pan"
			}
		default:
			{
				frase = strconv.Itoa(i)
			}
		}
		imprime()
	}

}

func imprime() {
	fmt.Println(frase)
}
