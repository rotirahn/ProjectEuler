package main

import "fmt"

//Palindrome checking function

func palindrome() {

	x := 998001
	z := 0

	var a, b, c, d, e, f int

	for x >= 100000 {

		a = x/1e5
		b = x/1e4 - a*10
		c = x/1e3 - a*1e2 - b*10
		d = x/1e2 - a*1e3 - b*1e2 - c*10
		e = x/10 - a*1e4 - b*1e3 - c*1e2 - d*10
		f = x - a*1e5 - b*1e4 - c*1e3 - d*1e2 - e*10

		if z == 0 {
			if a*1e5+b*1e4+c*1e3+d*1e2+e*10+f == f*1e5+e*1e4+d*1e3+c*1e2+b*10+a {
				z = factor(x)
			}
			x -= 1
		} else {
			break
		}
	}
}

//Factor finding function

func factor(x int) int {

	for i := 999; i >= 100; i-- {
		if x%i == 0 && x/i <= 999 {
			fmt.Println(x)
			return 1
		}
	}

	return 0
}

//Main function
func main() {
	palindrome()
}