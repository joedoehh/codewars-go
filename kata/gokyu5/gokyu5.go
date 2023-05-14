package gokyu5

import (
	"fmt"
	"math"
	"strconv"
)

/*
Buddy Pairs
5 kyu
https://www.codewars.com/kata/59ccf051dcc4050f7800008f/train/go
*/
func Buddy(start, limit int) []int {
	for i := start; i <= limit; i++ {
		divisors := properDivisors(i)
		buddy := sum(divisors) - 1
		if (buddy > i) && isBuddyOf(i, buddy) {
			return []int{i, buddy}
		}
	}
	return []int{}
}

func isBuddyOf(m int, n int) bool {
	return sum(properDivisors(n))-1 == m
}

func sum(divisors map[int]bool) (result int) {
	result = 0
	for d := range divisors {
		result += d
	}
	return
}

func properDivisors(n int) (result map[int]bool) {
	result = map[int]bool{}
	if n != 1 {
		result[1] = true
	}
	for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
		if (n % i) == 0 {
			result[i] = true
		}
		if (n % (n / i)) == 0 {
			result[n/i] = true
		}
	}
	return
}

/*
Human Readable Time
5 kyu
https://www.codewars.com/kata/52685f7382004e774f0001f7/train/go
*/
func HumanReadableTime(seconds int) string {
	s := seconds % 60
	m := (seconds / 60) % 60
	h := seconds / 3600
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

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
func KingIsInCheck(board [8][8]rune) (inCheck bool) {
	inCheck = false
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			switch board[i][j] {
			case '♟':
				inCheck = checkByPawn(board, i, j)
			case '♞':
				inCheck = checkByKnight(board, i, j)
			case '♜':
				inCheck = checkByRook(board, i, j)
			case '♝':
				inCheck = checkByBishop(board, i, j)
			case '♛':
				inCheck = checkByQueen(board, i, j)
			}
			if inCheck {
				return
			}
		}
	}
	return
}

func isKing(board [8][8]rune, i int, j int) bool {
	return onField(board, i, j, '♔')
}

func isBlockedByNotKing(board [8][8]rune, i int, j int) bool {
	return !onField(board, i, j, ' ') && !onField(board, i, j, '♔')
}

func onField(board [8][8]rune, i int, j int, expected rune) bool {
	return validField(board, i, j) && board[i][j] == expected
}

func validField(board [8][8]rune, i int, j int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board) {
		return false
	} else {
		return true
	}
}

func checkByPawn(board [8][8]rune, iPawn int, jPawn int) bool {
	return isKing(board, iPawn+1, jPawn-1) || isKing(board, iPawn+1, jPawn+1)
}

func checkByKnight(board [8][8]rune, iKnight int, jKnight int) bool {
	return isKing(board, iKnight-2, jKnight+1) || isKing(board, iKnight-2, jKnight-1) || isKing(board, iKnight-1, jKnight+2) || isKing(board, iKnight+1, jKnight+2) || isKing(board, iKnight+2, jKnight+1) || isKing(board, iKnight+2, jKnight-1) || isKing(board, iKnight-1, jKnight-2) || isKing(board, iKnight+1, jKnight-2)
}

func checkByQueen(board [8][8]rune, iQueen int, jQueen int) bool {
	return checkByRook(board, iQueen, jQueen) || checkByBishop(board, iQueen, jQueen)
}

func checkByRook(board [8][8]rune, iRook int, jRook int) bool {
	return checkByRookDirection(board, iRook, jRook, Up) || checkByRookDirection(board, iRook, jRook, Right) || checkByRookDirection(board, iRook, jRook, Down) || checkByRookDirection(board, iRook, jRook, Left)
}

type StraightDirection int

const (
	Up    StraightDirection = 0
	Right StraightDirection = 1
	Down  StraightDirection = 2
	Left  StraightDirection = 3
)

func checkByRookDirection(board [8][8]rune, iRook int, jRook int, dir StraightDirection) (result bool) {
	result = false
	nextI := iRook
	nextJ := jRook
	for i := 1; ; i++ {
		switch dir {
		case Up:
			nextI -= 1
		case Down:
			nextI += 1
		case Left:
			nextJ -= 1
		case Right:
			nextJ += 1
		}
		if !validField(board, nextI, nextJ) || isBlockedByNotKing(board, nextI, nextJ) {
			result = false
			break
		}
		if isKing(board, nextI, nextJ) {
			result = true
			break
		}
	}
	return
}

func checkByBishop(board [8][8]rune, iBishop int, jBishop int) bool {
	return checkByBishopDirection(board, iBishop, jBishop, UpRight) || checkByBishopDirection(board, iBishop, jBishop, DownRight) || checkByBishopDirection(board, iBishop, jBishop, UpLeft) || checkByBishopDirection(board, iBishop, jBishop, DownLeft)
}

type DiagonalDirection int

const (
	UpRight   DiagonalDirection = 0
	DownRight DiagonalDirection = 1
	UpLeft    DiagonalDirection = 2
	DownLeft  DiagonalDirection = 3
)

func checkByBishopDirection(board [8][8]rune, iBishop int, jBishop int, dir DiagonalDirection) (result bool) {
	result = false
	nextI := iBishop
	nextJ := jBishop
	for i := 1; ; i++ {
		switch dir {
		case UpRight:
			nextI -= 1
			nextJ += 1
		case DownRight:
			nextI += 1
			nextJ += 1
		case UpLeft:
			nextI -= 1
			nextJ -= 1
		case DownLeft:
			nextI += 1
			nextJ -= 1
		}
		if !validField(board, nextI, nextJ) || isBlockedByNotKing(board, nextI, nextJ) {
			result = false
			break
		}
		if isKing(board, nextI, nextJ) {
			result = true
			break
		}
	}
	return
}
