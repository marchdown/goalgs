package main

import "fmt"


var idx = 0

func main() {
	for idx < 100 {
		idx += 1
		case
		if (idx%15) == 0 {
			 fmt.Printf("Fizz-Buzz")
		} else if (idx%3) == 0 {
			 fmt.Printf("Fizz")
		} else if (idx%5)== 0 {
			 fmt.Printf("Buzz")
		} else {
			fmt.Printf("%d", idx)
		}
		fmt.Printf("\n")
	}
}
