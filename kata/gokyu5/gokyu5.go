package gokyu5

import (
	"fmt"
	"strconv"
)

/*
Product of consecutive Fib numbers
5 kyu
https://www.codewars.com/kata/5541f58a944b85ce6d00006a/train/go
*/
func ProductFib(prod uint64) [3]uint64 {
	fibs := fib(prod)
	product := uint64(0)
	n := 1
	for product < prod {
		product = fibs[n-1] * fibs[n]
		n += 1
	}
	if product == prod {
		return [3]uint64{fibs[n-2], fibs[n-1], uint64(1)}
	} else {
		return [3]uint64{fibs[n-2], fibs[n-1], uint64(0)}
	}
}

func fib(n uint64) (result []uint64) {
	result = make([]uint64, 0)
	result = append(result, 0)
	if n == 1 {
		return
	}
	result = append(result, 1)
	if n == 2 {
		return
	}
	for i := uint64(3); result[len(result)-1] <= n; i++ {
		result = append(result, result[i-3]+result[i-2])
	}
	return
}

/*
Prime in numbers
5 kyu
https://www.codewars.com/kata/54d512e62a5e54c96200019e/train/go
*/
func PrimeFactors(n int) (result string) {
	result = ""
	primeNumbers := primeFactors(n)
	lastNumber := 0
	for _, p := range primeNumbers {
		if p == lastNumber {
			continue
		}
		lastNumber = p
		occurences := occurences(primeNumbers, p)
		if occurences == 1 {
			result += "(" + strconv.Itoa(p) + ")"
		} else {
			result += "(" + strconv.Itoa(p) + "**" + strconv.Itoa(occurences) + ")"
		}
	}
	return
}

func primeFactors(number int) (result []int) {
	result = make([]int, 0)
	for i := 2; i <= number/i; i++ {
		for number%i == 0 {
			result = append(result, i)
			number /= i
		}
	}
	if number > 1 {
		result = append(result, number)
	}
	return
}

func occurences(numbers []int, number int) (result int) {
	result = 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == number {
			result += 1
		}
	}
	return
}

/*
RGB to hex conversion
5 kyu
https://www.codewars.com/kata/513e08acc600c94f01000001/train/go
*/
func RGB(r, g, b int) string {
	return rgb(r) + rgb(g) + rgb(b)
}

func rgb(value int) string {
	if value <= 0 {
		return "00"
	} else if value >= 255 {
		return "FF"
	} else {
		hex := fmt.Sprintf("%X", value)
		if len(hex) == 1 {
			hex = "0" + hex
		}
		return hex
	}
}

/*
King in check
5 kyu
https://www.codewars.com/kata/5e28ae347036fa001a504bbe/train/go
*/
func KingIsInCheck(board [8][8]rune) bool {
	//  Do your magic here
	return false
}
