package module01

import "fmt"

// GCD func
func GCD(a, b int) int {
	// find factors of a & b i.e. array n1,n2
	// find intersection of factors n1,n2 i.e. final
	// multiply all elements in final
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	n1, n2 := Factor(primes, a), Factor(primes, b)
	final := intersection(n1, n2)
	fmt.Println(final)

	mul := 1
	for _, f := range final {
		mul = mul * f
	}
	return mul
	// solution euclidean formula gcd
	// for {
	// 	if b == 0 {
	// 		return a
	// 	}

	// 	a, b = b, a%b
	// }
	// return 0
}

func intersection(n1, n2 []int) (final []int) {
	for _, k := range n1 {
		present := false
		for _, l := range n2 {
			if k == l {
				present = true
				break
			}
		}
		if present == true {
			fl := false
			for _, f := range final {
				if f == k {
					fl = true
					break
				}
			}
			if fl == false {
				final = append(final, k)
			}
		}
	}
	return
}
