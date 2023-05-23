package gokyu5

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
Pick peaks
5 kyu
https://www.codewars.com/kata/5279f6fe5ab7f447890006a7/train/go
*/
type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) (result PosPeaks) {
	fmt.Printf("Peak Array: %v\n", array)
	delta := []string{}
	// derive ups´n´downs
	for i := 1; i < len(array); i++ {
		if array[i] > array[i-1] {
			delta = append(delta, "+")
		} else if array[i] < array[i-1] {
			delta = append(delta, "-")
		} else {
			delta = append(delta, "=")
		}

	}
	// collect peaks
	result.Pos = make([]int, 0)
	result.Peaks = make([]int, 0)
	posPlateau := -1
	peaksPlateau := -1
	for i := 1; i < len(delta); i++ {
		if delta[i-1] == "+" && delta[i] == "=" {
			posPlateau = i
			peaksPlateau = array[i]
		} else if delta[i-1] == "+" && delta[i] == "-" {
			result.Pos = append(result.Pos, i)
			result.Peaks = append(result.Peaks, array[i])
		} else if delta[i-1] == "=" && delta[i] == "-" {
			if posPlateau == -1 {
				continue
			}
			result.Pos = append(result.Pos, posPlateau)
			result.Peaks = append(result.Peaks, peaksPlateau)
			posPlateau = -1
			peaksPlateau = -1
		}
	}
	return
}

/*
Coding squared string
5 kyu
https://www.codewars.com/kata/56fcc393c5957c666900024d/train/go
*/
func Code(s string) (coded string) {
	fmt.Printf("Code INPUT: %q\n", s)
	if s == "" {
		return ""
	}
	// determine square size and padding missing chars
	n := squareSize(s, true)
	s += strings.Repeat("\v", n*n-len(s))
	// square string
	square := make([][]byte, n)
	for i := range square {
		square[i] = make([]byte, n)
		for j := range square[i] {
			square[i][j] = s[i*n+j]
		}
	}
	// rotate square 90 degree clockwise
	squareRotated := make([][]byte, n)
	for i := range square {
		squareRotated[i] = make([]byte, n)
		for j := range square[i] {
			squareRotated[i][j] = square[n-j-1][i]
		}
	}
	// convert square to code string
	for i := range square {
		for j := range square[i] {
			coded += string(squareRotated[i][j])
		}
		if i != n-1 {
			coded += "\n"
		}
	}
	return
}

func Decode(s string) (decoded string) {
	fmt.Printf("Decode: %q\n", s)
	if s == "" {
		return ""
	}
	// determine square size (remove \n)
	s = strings.ReplaceAll(s, "\n", "")
	n := squareSize(s, false)
	// square string
	square := make([][]byte, n)
	for i := range square {
		square[i] = make([]byte, n)
		for j := range square[i] {
			square[i][j] = s[i*n+j]
		}
	}
	// rotate square 90 degree counter-clockwise
	squareRotated := make([][]byte, n)
	for i := range square {
		squareRotated[i] = make([]byte, n)
		for j := range square[i] {
			squareRotated[i][j] = square[j][n-i-1]
		}
	}
	// convert square to decode string
	for i := range square {
		for j := range square[i] {
			decoded += string(squareRotated[i][j])
		}
	}
	decoded = strings.ReplaceAll(decoded, "\v", "")
	return
}

func squareSize(s string, ceil bool) int {
	sqrt := math.Sqrt(float64(len(s)))
	if ceil {
		return int(sqrt) + 1
	} else {
		return int(sqrt)
	}
}

/*
Fibo akin
5 kyu
https://www.codewars.com/kata/5772382d509c65de7e000982/train/go
*/
func LengthSupUk(n, k int) (count int) {
	U(n)
	count = 0
	for i := 1; i <= n; i++ {
		if U(i) >= k {
			count += 1
		}
	}
	return
}

func Comp(n int) (count int) {
	U(n)
	count = 0
	for i := 2; i <= n; i++ {
		if U(i) < U(i-1) {
			count += 1
		}
	}
	return
}

var uCache map[int]int = map[int]int{1: 1, 2: 1}

func U(n int) (u int) {
	u, exists := uCache[n]
	if !exists {
		u = computeU(n)
		uCache[n] = u
	}
	return
}

func computeU(n int) int {
	series := []int{1, 1}
	for i := 2; i < n; i++ {
		index1 := i - series[i-1]
		index2 := i - series[i-2]
		newValue := series[index1] + series[index2]
		series = append(series, newValue)
		uCache[i+1] = newValue
	}
	return series[n-1]
}

/*
Help your granny
5 kyu
https://www.codewars.com/kata/5536a85b6ed4ee5a78000035/train/go
*/
func Tour(arrFriends []string, ftwns map[string]string, h map[string]float64) int {
	distance := 0.
	lastDistance := -1.
	for _, nextFriend := range arrFriends {
		// get distance to town (ans skip unknown towns)
		nextTwn, known := ftwns[nextFriend]
		if !known {
			continue
		}
		nextDistance, known := h[nextTwn]
		if !known {
			continue
		}
		// add distance: special case first distance, compute otherwise
		if lastDistance < 0 {
			distance += nextDistance
		} else {
			distance += math.Sqrt(nextDistance*nextDistance - lastDistance*lastDistance)
		}
		lastDistance = nextDistance
	}
	distance += lastDistance
	return int(distance)
}

/*
John/ann Sign Up Codewars
5 kyu
https://www.codewars.com/kata/57591ef494aba64d14000526/train/go
*/
func Ann(n int) []int {
	_, a := createKataSeries(n, []int{0}, []int{1})
	return a
}

func John(n int) []int {
	j, _ := createKataSeries(n, []int{0}, []int{1})
	return j
}

func SumJohn(n int) (result int) {
	result = 0
	for _, v := range John(n) {
		result += v
	}
	return
}

func SumAnn(n int) (result int) {
	result = 0
	for _, v := range Ann(n) {
		result += v
	}
	return
}

func createKataSeries(nLimit int, j, a []int) (jSeries, aSeries []int) {
	for n := 1; n < nLimit; n++ {
		newJ := n - a[j[n-1]]
		j = append(j, newJ)
		newA := n - j[a[n-1]]
		a = append(a, newA)
	}
	jSeries = j
	aSeries = a
	return
}

/*
Basic Nico Variation
5 kyu
https://www.codewars.com/kata/5968bb83c307f0bb86000015/train/go
*/
func Nico(key, message string) (result string) {
	keyArray := createKey(key)
	if len(message)%len(key) != 0 {
		padding := len(key)*((len(message)/len(key))+1) - len(message)
		for i := 0; i <= padding; i++ {
			message += " "
		}
	}
	for i := len(key); i <= len(message); i += len(key) {
		result += encode(message[i-len(key):i], keyArray)
	}
	return
}

func encode(s string, key []byte) (result string) {
	posMap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		posMap[key[i]] = s[i]
	}
	for i := 0; i < len(key); i++ {
		result += string(posMap[byte(i+1)])
	}
	return
}

func createKey(keyString string) (key []byte) {
	key = make([]byte, len(keyString))
	s := strings.Split(keyString, "")
	sort.Strings(s)
	keyMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		keyMap[s[i]] = i + 1
	}
	for i := 0; i < len(keyString); i++ {
		nextString := string(keyString[i])
		key[i] = byte(keyMap[nextString])
	}
	return
}

/*
Find The Smallest
5 kyu
https://www.codewars.com/kata/573992c724fc289553000e95/train/go
*/
func Smallest(n int64) []int64 {
	smallest := n
	fromIndexSmallest := int64(0)
	toIndexSmallest := int64(0)
	digits := digits(n)
	for fromIndex := 0; fromIndex < len(digits); fromIndex++ {
		for toIndex := 0; toIndex < len(digits); toIndex++ {
			if fromIndex == toIndex {
				continue
			}
			permutated := shuffle(digits, fromIndex, toIndex)
			if permutated < smallest {
				smallest = permutated
				fromIndexSmallest = int64(fromIndex)
				toIndexSmallest = int64(toIndex)
			}
		}
	}
	return []int64{smallest, fromIndexSmallest, toIndexSmallest}
}

func shuffle(digits []int64, from int, to int) int64 {
	asString := ""
	tmp := digits[from]
	toIndex := 0
	fromIndex := 0
	for toIndex < len(digits) {
		if fromIndex == from {
			fromIndex++
		}
		if toIndex == to {
			asString += strconv.FormatInt(tmp, 10)
			fromIndex--
		} else {
			asString += strconv.FormatInt(digits[fromIndex], 10)
		}
		toIndex++
		fromIndex++
	}
	returnValue, _ := strconv.Atoi(asString)
	return int64(returnValue)
}

func digits(n int64) (result []int64) {
	result = []int64{}
	nAsString := strconv.FormatInt(n, 10)
	for i := 0; i < len(nAsString); i++ {
		result = append(result, int64(nAsString[i]-48))
	}
	return
}

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
