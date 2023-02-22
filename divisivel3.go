package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("O número: %d é divisivel por 3", i)
			fmt.Println()

		}

	}

}
