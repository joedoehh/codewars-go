package kata

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// 6 kyu -----------------------------------------------------------------------------------------

/*
Pair Of Gloves
6 kyu
https://www.codewars.com/kata/58235a167a8cb37e1a0000db/train/go
*/
func NumberOfPairs(gloves []string) (result int) {
	colorToNumber := make(map[string]int)
	for _, glove := range gloves {
		oldValue, contained := colorToNumber[glove]
		if !contained {
			oldValue = 0
		}
		oldValue++
		colorToNumber[glove] = oldValue
	}
	for _, number := range colorToNumber {
		if number%2 == 0 {
			result += number / 2
		} else if (number-1)%2 == 0 {
			result += (number - 1) / 2
		}
	}
	return
}

/*
Playing with digits
6 kyu
https://www.codewars.com/kata/5552101f47fc5178b1000050/train/go
*/
func DigPow(n, p int) int {
	pow := computePow(n, p)
	div := pow / n
	remainder := pow % n
	if remainder != 0 {
		return -1
	} else {
		return div
	}
}

func computePow(n, p int) int {
	digits := digits(n)
	power := p
	resultFloat := float64(0)
	for i := 0; i < len(digits); i++ {
		resultFloat += math.Pow(float64(digits[i]), float64(power))
		power++
	}
	return int(resultFloat)
}

func digits(n int) (result []int) {
	asString := strconv.Itoa(n)
	for _, rune := range asString {
		result = append(result, int(rune-'0'))
	}
	return
}

/*
Character with longest consecutive repetition
6 kyu
https://www.codewars.com/kata/586d6cefbcc21eed7a001155/train/go
*/
func LongestRepetition(text string) Result {
	if len(text) == 0 {
		return Result{}
	}
	textRune := []rune(text)
	currentLength := 0
	currentChar := textRune[0]
	maxLength := 0
	maxChar := currentChar
	for _, nextChar := range textRune {
		if currentChar == nextChar {
			currentLength++
		} else {
			if currentLength > maxLength {
				maxLength = currentLength
				maxChar = currentChar
			}
			currentChar = nextChar
			currentLength = 1
		}
	}
	if currentLength > maxLength {
		maxLength = currentLength
		maxChar = currentChar
	}
	return Result{maxChar, maxLength}
}

type Result struct {
	C rune // character
	L int  // count
}

/*
Bell (6 kyu)
For whom the Bell tolls
Write a function bell that will receive a positive integer and return an array. That's all splaining you receive; what needs to be done you'll have to figure out with the examples below.

	n => resulting array

	1 => [1]
	2 => [2, 2]
	3 => [3, 4, 3]
	4 => [4, 6, 6, 4]
	5 => [5, 8, 9, 8, 5]
	6 => [6, 10, 12, 12, 10, 6]
	7 => [7, 12, 15, 16, 15, 12, 7]
	8 => [8, 14, 18, 20, 20, 18, 14, 8]
	9 => [9, 16, 21, 24, 25, 24, 21, 16, 9]

10 => [10, 18, 24, 28, 30, 30, 28, 24, 18, 10]
11 => [11, 20, 27, 32, 35, 36, 35, 32, 27, 20, 11]
12 => [12, 22, 30, 36, 40, 42, 42, 40, 36, 30, 22, 12]
*/
func Bell(n int) (result []int) {
	var bellArray [][]int
	// build bell
	for i := 0; i < n; i++ {
		row := []int{}
		for j := 0; j < n-i; j++ {
			row = append(row, (j+1)*(i+1))
		}
		bellArray = append(bellArray, row)
	}
	// traverse bell
	for i := n - 1; i >= 0; i-- {
		result = append(result, bellArray[i][n-i-1])
	}
	return
}

/*
IP Validation (6 kyu)

Write an algorithm that will identify valid IPv4 addresses in dot-decimal format.
IPs should be considered valid if they consist of four octets, with values between 0 and 255, inclusive.

Valid inputs examples:
Examples of valid inputs:
1.2.3.4
123.45.67.89

Invalid input examples:
1.2.3
1.2.3.4.5
123.456.78.90
123.045.067.089
Notes:

Leading zeros (e.g. 01.02.03.04) are considered invalid

Inputs are guaranteed to be a single string
*/
func Is_valid_ip(ip string) bool {
	octets := strings.Split(ip, ".")
	if len(octets) != 4 {
		return false
	}
	for _, o := range octets {
		isValid := checkOctet(o)
		if !isValid {
			return false
		}
	}
	return true
}

func checkOctet(octet string) bool {
	u, err := strconv.ParseUint(octet, 10, 32)
	if err != nil {
		return false
	}
	if octet[0] == '0' && len(octet) > 1 {
		return false
	}
	return 0 <= u && u <= 255
}

/*
Mexican Wave (6 kyu)

Task
In this simple Kata your task is to create a function that turns a string into a Mexican Wave. You will be passed a string and you must return that string in an array where an uppercase letter is a person standing up.
Rules

 1. The input string will always be lower case but maybe empty.

 2. If the character in the string is whitespace then pass over it as if it was an empty seat

    Example

    wave("hello") => []string{"Hello", "hEllo", "heLlo", "helLo", "hellO"}

    Good luck and enjoy!
*/
func Wave(words string) (wave []string) {
	for index, rune := range words {
		if unicode.IsSpace(rune) {
			continue
		}
		wave = append(wave, standUp(words, index))
	}
	return
}

func standUp(word string, index int) string {
	return word[:index] + strings.ToUpper(string(word[index])) + word[index+1:]
}

/*
Camel Case Again (6 kyu)

Write simple .camelCase method (camel_case function in PHP, CamelCase in C# or camelCase in Java) for strings. All words must have their first letter capitalized without spaces.

For instance:

CamelCase("hello case")      // => "HelloCase"
CamelCase("camel case word") // => "CamelCaseWord"
Don't forget to rate this kata! Thanks :)
*/
func CamelCase(s string) string {
	if s == "" {
		return ""
	}
	wordsSplit := strings.Split(s, " ")
	fmt.Println("splitted: ", wordsSplit)
	wordsUpper := []string{}
	for _, word := range wordsSplit {
		if word == "" {
			continue
		}
		wordUpper := strings.ToUpper(string(word[0])) + string(word[1:])
		wordsUpper = append(wordsUpper, wordUpper)
	}
	return strings.Join(wordsUpper, "")
}

/*
Queue (6 kyu)
There is a queue for the self-checkout tills at the supermarket. Your task is write a function to calculate the total time required for all the customers to check out!

input
customers: an array of positive integers representing the queue. Each integer represents a customer, and its value is the amount of time they require to check out.
n: a positive integer, the number of checkout tills.
output
The function should return an integer, the total time required.

Important
Please look at the examples and clarifications below, to ensure you understand the task correctly :)

Examples
queueTime([5,3,4], 1)
// should return 12
// because when there is 1 till, the total time is just the sum of the times

queueTime([10,2,3,3], 2)
// should return 10
// because here n=2 and the 2nd, 3rd, and 4th people in the
// queue finish before the 1st person has finished.

queueTime([2,3,10], 2)
// should return 12
Clarifications
There is only ONE queue serving many tills, and
The order of the queue NEVER changes, and
The front person in the queue (i.e. the first element in the array/list) proceeds to a till as soon as it becomes free.
N.B. You should assume that all the test input will be valid, as specified above.

P.S. The situation in this kata can be likened to the more-computer-science-related idea of a thread pool,
with relation to running multiple processes at the same time: https://en.wikipedia.org/wiki/Thread_pool
*/
func QueueTime(customers []int, n int) (totalTime int) {
	queues := make([]int, n)
	for i := 0; i < len(customers); i++ {
		placed := distribute(customers[i], queues)
		if !placed {
			totalTime += wait(queues)
			i--
			continue
		}
	}
	for !isEmpty(queues) {
		totalTime += wait(queues)
	}
	return
}

func distribute(time int, queues []int) (placed bool) {
	placed = false
	for i := 0; i < len(queues); i++ {
		if queues[i] == 0 {
			queues[i] = time
			placed = true
			break
		}
	}
	return
}

func wait(queues []int) (waitTime int) {
	// find smallest wait time
	waitTime = math.MaxInt32
	for _, time := range queues {
		if time <= waitTime && time != 0 {
			waitTime = time
		}
	}
	// adjust by smallest wait time and clear queue
	for index, time := range queues {
		if time == waitTime {
			queues[index] = 0
		} else if time != 0 {
			queues[index] -= waitTime
		}
	}
	return
}

func isEmpty(queues []int) (empty bool) {
	empty = true
	for _, time := range queues {
		if time > 0 {
			empty = false
			break
		}
	}
	return
}

/*
Prime (6 kyu)
Define a function that takes an integer argument and returns a logical value true or false depending on if the integer is a prime.

Per Wikipedia, a prime number ( or a prime ) is a natural number greater than 1 that has no positive divisors other than 1 and itself.

Requirements
You can assume you will be given an integer input.
You can not assume that the integer will be only positive. You may be given negative numbers as well ( or 0 ).
NOTE on performance: There are no fancy optimizations required, but still the most trivial solutions might time out.
Numbers go up to 2^31 ( or similar, depending on language ). Looping all the way up to n, or n/2, will be too slow.

Example
is_prime(1)  false
is_prime(2)  true
is_prime(-1) false
*/
func IsPrime(n int) (isPrime bool) {
	isPrime = true
	if n <= 1 {
		isPrime = false
	} else if n > 2 {
		limit := int(math.Sqrt(float64(n)))
		for i := 2; i <= limit; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
	}
	return
}

/*
Nato Alphabet Encoding (6 kyu)

Task
You'll have to translate a string to Pilot's alphabet (NATO phonetic alphabet).

Input:

If, you can read?

Output:

India Foxtrot , Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta ?

Note:

There are preloaded dictionary you can use, named NATO
The set of used punctuation is ,.!?.
Punctuation should be kept in your return string, but spaces should not.
Xray should not have a dash within.
Every word and punctuation mark should be seperated by a space ' '.
There should be no trailing whitespace
*/
func ToNato(words string) (encoded string) {
	for _, rune := range strings.ToUpper(words) {
		switch rune {
		case ',', '.', '!', '?':
			encoded += " " + string(rune)
		case ' ':
			encoded += ""
		default:
			encoded += " " + natoAlphabet[string(rune)]
		}
	}
	encoded = strings.TrimSpace(encoded)
	return
}

var natoAlphabet = map[string]string{
	"A": "Alfa", "B": "Bravo", "C": "Charlie", "D": "Delta", "E": "Echo",
	"F": "Foxtrot", "G": "Golf", "H": "Hotel", "I": "India", "J": "Juliett",
	"K": "Kilo", "L": "Lima", "M": "Mike", "N": "November", "O": "Oscar",
	"P": "Papa", "Q": "Quebec", "R": "Romeo", "S": "Sierra", "T": "Tango",
	"U": "Uniform", "V": "Victor", "W": "Whiskey", "X": "X-ray", "Y": "Yankee",
	"Z": "Zulu",
}

/*
Tower Builder (6 kyu)

Build Tower
Build a pyramid-shaped tower, as an array/list of strings, given a positive integer number of floors. A tower block is represented with "*" character.

For example, a tower with 3 floors looks like this:

[

	"  *  ",
	" *** ",
	"*****"

]
And a tower with 6 floors looks like this:

[

	"     *     ",
	"    ***    ",
	"   *****   ",
	"  *******  ",
	" ********* ",
	"***********"

]
Go challenge Build Tower Advanced once you have finished this :)
*/
func TowerBuilder(nFloors int) (tower []string) {
	tower, _ = towerBuilder(1, nFloors, 1)
	// reverse tower (base is on top)
	for i, j := 0, len(tower)-1; i < j; i, j = i+1, j-1 {
		tower[i], tower[j] = tower[j], tower[i]
	}
	return tower
}

func towerBuilder(level, nFloors, numberOfStars int) (levels []string, baseSize int) {
	if level == nFloors {
		baseSize = numberOfStars
		levels = append(levels, strings.Repeat("*", baseSize))
		return
	} else {
		// get levels below
		levels, baseSize = towerBuilder(level+1, nFloors, numberOfStars+2)
		// add current level
		paddingLength := (baseSize - numberOfStars) / 2
		levelString := strings.Repeat(" ", paddingLength) + strings.Repeat("*", numberOfStars) + strings.Repeat(" ", paddingLength)
		levels = append(levels, levelString)
		return
	}
}

/*
Find Unique Number (6 kyu)

There is an array with some numbers. All numbers are equal except for one. Try to find it!

findUniq([ 1, 1, 1, 2, 1, 1 ]) === 2
findUniq([ 0, 0, 0.55, 0, 0 ]) === 0.55
It’s guaranteed that array contains at least 3 numbers.

The tests contain some very huge arrays, so think about performance.

This is the first kata in series:

Find the unique number (this kata)
Find the unique string
Find The Unique
*/
func FindUniq(arr []float32) (unique float32) {
	// count occurences
	occurences := map[float32]int{}
	for _, next := range arr {
		occurences[next]++
	}
	// return occurence
	for key, value := range occurences {
		if value == 1 {
			unique = key
			break
		}
	}
	return
}

/*
Tortoise Race (6 kyu)

Two tortoises named A and B must run a race. A starts with an average speed of 720 feet per hour. Young B knows she runs faster than A, and furthermore has not finished her cabbage.

When she starts, at last, she can see that A has a 70 feet lead but B's speed is 850 feet per hour. How long will it take B to catch A?

More generally: given two speeds v1 (A's speed, integer > 0) and v2 (B's speed, integer > 0) and a lead g (integer > 0) how long will it take B to catch A?

The result will be an array [hour, min, sec] which is the time needed in hours, minutes and seconds (round down to the nearest second) or a string in some languages.

If v1 >= v2 then return nil, nothing, null, None or {-1, -1, -1} for C++, C, Go, Nim, Pascal, COBOL, Erlang, [-1, -1, -1] for Perl,[] for Kotlin or "-1 -1 -1" for others.

Examples:
(form of the result depends on the language)

race(720, 850, 70) => [0, 32, 18] or "0 32 18"
race(80, 91, 37)   => [3, 21, 49] or "3 21 49"
Note:
See other examples in "Your test cases".

In Fortran - as in any other language - the returned string is not permitted to contain any redundant trailing whitespace: you can use dynamically allocated character strings.

** Hints for people who don't know how to convert to hours, minutes, seconds:

Tortoises don't care about fractions of seconds
Think of calculation by hand using only integers (in your code use or simulate integer division)
or Google: "convert decimal time to hours minutes seconds"
*/
func Race(v1, v2, g int) [3]int {
	// special case: no race - tortoise 1 i faster anyway
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}
	// simulate second by second (until race is won)
	v1PosFeet := float64(g)
	v1FeetPerS := float64(v1) / 60.0 / 60.0
	v2PosFeet := float64(0)
	v2FeetPerS := float64(v2) / 60.0 / 60.0
	second := 0
	for i := 0; ; i++ {
		if v2PosFeet >= v1PosFeet {
			second = i
			break
		}
		v1PosFeet += v1FeetPerS
		v2PosFeet += v2FeetPerS
	}
	// transfer result
	hours := second / 3600
	minutes := (second % 3600) / 60
	seconds := second%60 - 1
	if seconds < 0 {
		seconds = 0
	}
	return [3]int{hours, minutes, seconds}
}

/*
Camel Case (6 kyu)

Complete the method/function so that it converts dash/underscore delimited words into camel casing.
The first word within the output should be capitalized only if the original word was capitalized
(known as Upper Camel Case, also often referred to as Pascal case). The next words should be always
capitalized.

Examples
"the-stealth-warrior" gets converted to "theStealthWarrior"

"The_Stealth_Warrior" gets converted to "TheStealthWarrior"

"The_Stealth-Warrior" gets converted to "TheStealthWarrior"
*/
func ToCamelCase(s string) string {
	wordsSplit := strings.FieldsFunc(s, splitCamelCase)
	wordsUpper := []string{}
	for i, word := range wordsSplit {
		var wordUpper string
		if i == 0 {
			// dont uppercase first word
			wordUpper = word
		} else {
			// uppercase all other words
			wordUpper = strings.ToUpper(string(word[0])) + string(word[1:])
		}
		wordsUpper = append(wordsUpper, wordUpper)

	}
	return strings.Join(wordsUpper, "")
}

func splitCamelCase(r rune) bool {
	return r == '-' || r == '_'
}

/*
Count Bits (6 kyu)

Write a function that takes an integer as input, and returns the number of bits that are equal to one in the binary representation of that number.
You can guarantee that input is non-negative.

Example: The binary representation of 1234 is 10011010010, so the function should return 5 in this case
*/
func CountBits(nr uint) (countedBits int) {
	for _, v := range fmt.Sprintf("%b", nr) {
		if v == '1' {
			countedBits++
		}
	}
	return
}

/*
Create Phone Number (6 kyu)

Write a function that accepts an array of 10 integers (between 0 and 9), that returns a string of those numbers in the form of a phone number.

Example
CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})  // returns "(123) 456-7890"
The returned format must be correct in order to complete this challenge.

Don't forget the space after the closing parentheses!
*/
func CreatePhoneNumber(numbers [10]uint) string {
	return fmt.Sprintf("(%v) %v-%v", makeNumber(numbers, 0, 3), makeNumber(numbers, 3, 3), makeNumber(numbers, 6, 4))
}

func makeNumber(numbers [10]uint, start, length int) (substringNr string) {
	for i := start; i < length+start; i++ {
		substringNr += fmt.Sprintf("%v", numbers[i])
	}
	return
}

/*
Find the odd int (6 kyu)

Given an array of integers, find the one that appears an odd number of times.

There will always be only one integer that appears an odd number of times.

Examples
[7] should return 7, because it occurs 1 time (which is odd).
[0] should return 0, because it occurs 1 time (which is odd).
[1,1,2] should return 2, because it occurs 1 time (which is odd).
[0,1,0,1,0] should return 0, because it occurs 3 times (which is odd).
[1,2,2,3,3,3,4,3,3,3,2,2,1] should return 4, because it appears 1 time (which is odd).
*/
func FindOdd(seq []int) (occurence int) {
	// count occurences
	intToOccurence := map[int]int{}
	for _, nextInt := range seq {
		intToOccurence[nextInt] += 1
	}
	// return first odd occurence
	for key, value := range intToOccurence {
		if value%2 != 0 {
			occurence = key
			break
		}
	}
	return
}

/*
Stop gninnipS My sdroW! (6 kyu)

Write a function that takes in a string of one or more words, and returns the same string,
but with all five or more letter words reversed (Just like the name of this Kata).

Strings passed in will consist of only letters and spaces. Spaces will be included only when more than one word is present.

Examples:

spinWords( "Hey fellow warriors" ) => returns "Hey wollef sroirraw"
spinWords( "This is a test") => returns "This is a test"
spinWords( "This is another test" )=> returns "This is rehtona test"
*/
func SpinWords(str string) string {
	var words []string
	for _, word := range strings.Split(str, " ") {
		if len(word) >= 5 {
			words = append(words, spin(word))
		} else {
			words = append(words, word)
		}
	}
	return strings.Join(words, " ")
}

func spin(word string) (spinned string) {
	for i := len(word) - 1; i >= 0; i-- {
		spinned += string(word[i])
	}
	return
}

/*
Multiples of 3 and 5 (6 kyu)

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Finish the solution so that it returns the sum of all the multiples of 3 or 5 below the number passed in.
Additionally, if the number is negative, return 0 (for languages that do have them).

Note: If the number is a multiple of both 3 and 5, only count it once.

Courtesy of projecteuler.net (Problem 1)
*/
func Multiple3And5(number int) (sum int) {
	// special case handling
	if number <= 0 {
		sum = 0
		return
	}
	// enumerate and sum
	for i := 1; i < number; i++ {
		if multiple(i, 3) || multiple(i, 5) {
			sum += i
		}
	}
	return
}

func multiple(number, divisor int) bool {
	return number%divisor == 0
}

/*
Millipede of words (6 kyu)

The set of words is given. Words are joined if the last letter of one word and the first letter of another word are the same.
Return true if all words of the set can be combined into one word. Each word can and must be used only once. Otherwise return false.

Input
Array of 3 to 7 words of random length. No capital letters.

Example true
Set: excavate, endure, desire, screen, theater, excess, night.
Millipede: desirE EndurE ExcavatE ExcesS ScreeN NighT Theater.

Example false
Set: trade, pole, view, grave, ladder, mushroom, president.
Millipede: presidenT Trade.
*/
func Millipede(words []string) bool {
	return millipede("", words, 0, len(words))
}

func millipede(sentence string, words []string, wordsUsed, totalNrOfWords int) (result bool) {
	if len(words) == 0 && wordsUsed == totalNrOfWords {
		return true
	}
	for _, word := range words {
		if fits(sentence, word) {
			result = millipede(sentence+" "+word, dropFirstOccurence(word, words), wordsUsed+1, totalNrOfWords)
			if result {
				return
			}
		}
	}
	return false
}

func dropFirstOccurence(word string, words []string) (result []string) {
	droppingMode := true
	for _, nextWord := range words {
		if word == nextWord {
			if droppingMode {
				droppingMode = false
			} else {
				result = append(result, nextWord)
			}
		} else {
			result = append(result, nextWord)
		}
	}
	return
}

func fits(sentence, word string) bool {
	if sentence == "" {
		return true
	}
	sentenceRune := []rune(sentence)
	wordRune := []rune(word)
	return sentenceRune[len(sentenceRune)-1] == wordRune[0]
}

// 7 kyu -----------------------------------------------------------------------------------------

/*
Printer Errors

In a factory a printer prints labels for boxes. For one kind of boxes the printer has to use colors which,
for the sake of simplicity, are named with letters from a to m.

The colors used by the printer are recorded in a control string. For example a "good" control string would be
aaabbbbhaijjjm meaning that the printer used three times color a, four times color b, one time color h then one time color a...

Sometimes there are problems: lack of colors, technical malfunction and a "bad" control string is produced
e.g. aaaxbbbbyyhwawiwjjjwwm with letters not from a to m.

You have to write a function printer_error which given a string will return the error rate of the printer as a
string representing a rational whose numerator is the number of errors and the denominator the length of the control string.
Don't reduce this fraction to a simpler expression.

The string has a length greater or equal to one and contains only letters from a to z.

Examples:
s="aaabbbbhaijjjm"
printer_error(s) => "0/14"

s="aaaxbbbbyyhwawiwjjjwwm"
printer_error(s) => "8/22"
*/
func PrinterError(s string) string {
	errors := 0
	for _, c := range s {
		if !strings.Contains("abcdedfghijklm", string(c)) {
			errors += 1
		}
	}
	return fmt.Sprint(errors, "/", len(s))
}

/*
Perfect Square

You might know some pretty large perfect squares. But what about the NEXT one?

Complete the findNextSquare method that finds the next integral perfect square after the one passed as a parameter.
Recall that an integral perfect square is an integer n such that sqrt(n) is also an integer.

If the parameter is itself not a perfect square then -1 should be returned. You may assume the parameter is non-negative.

Examples:(Input --> Output)

121 --> 144
625 --> 676
114 --> -1 since 114 is not a perfect square
*/
func FindNextSquare(sq int64) int64 {
	sqrt := math.Sqrt(float64(sq))
	sqFloor := math.Floor(sqrt)
	if float64(sq) != math.Pow(sqFloor, 2) {
		// special case: check if sq is not perfect square
		return -1
	} else {
		// return next perfect square
		return int64(math.Pow(sqFloor+1.0, 2))
	}
}

/*
Two to One

Take 2 strings s1 and s2 including only letters from a to z.
Return a new sorted string, the longest possible, containing distinct letters - each taken only once - coming from s1 or s2.

Examples:
a = "xyaabbbccccdefww"
b = "xxxxyyyyabklmopq"
longest(a, b) -> "abcdefklmopqwxy"

a = "abcdefghijklmnopqrstuvwxyz"
longest(a, a) -> "abcdefghijklmnopqrstuvwxyz"
*/
func TwoToOne(s1 string, s2 string) (result string) {
	for _, c := range "abcdefghijklmnopqrstuvwxyz" {
		if strings.Contains(s1, string(c)) || strings.Contains(s2, string(c)) {
			result += string(c)
		}
	}
	return
}

/*
Sum Of Numbers

Given two integers a and b, which can be positive or negative, find the sum of all the integers between and including them and return it. If the two numbers are equal return a or b.

Note: a and b are not ordered!

Examples (a, b) --> output (explanation)
(1, 0) --> 1 (1 + 0 = 1)
(1, 2) --> 3 (1 + 2 = 3)
(0, 1) --> 1 (0 + 1 = 1)
(1, 1) --> 1 (1 since both are same)
(-1, 0) --> -1 (-1 + 0 = -1)
(-1, 2) --> 2 (-1 + 0 + 1 + 2 = 2)

Your function should only return a number, not the explanation about how you get that number.
*/
func GetSum(a, b int) (sum int) {
	// order interval
	if a > b {
		a, b = b, a
	}
	// iterate and sum up
	for i := a; i <= b; i++ {
		sum += i
	}
	return
}

/*
DNA

Deoxyribonucleic acid (DNA) is a chemical found in the nucleus of cells and carries the "instructions" for the development and functioning of living organisms.

If you want to know more: http://en.wikipedia.org/wiki/DNA

In DNA strings, symbols "A" and "T" are complements of each other, as "C" and "G".
Your function receives one side of the DNA (string, except for Haskell); you need to return the other complementary side.
DNA strand is never empty or there is no DNA at all (again, except for Haskell).

More similar exercise are found here: http://rosalind.info/problems/list-view/ (source)

Example: (input --> output)

"ATTGC" --> "TAACG"
"GTAT" --> "CATA"
*/
func DNAStrand(dna string) (result string) {
	for _, v := range dna {
		switch string(v) {
		case "A":
			result += "T"
		case "T":
			result += "A"
		case "C":
			result += "G"
		case "G":
			result += "C"
		default:
			result += string(v)
		}
	}
	return
}

/*
Shortest Word

Simple, given a string of words, return the length of the shortest word(s).

String will never be empty and you do not need to account for different data types.
*/
func FindShort(s string) (result int) {
	result = math.MaxInt32
	for _, word := range strings.Split(s, " ") {
		if len(word) <= result {
			result = len(word)
		}
	}
	return
}

/*
Jaden Casing

Jaden Smith, the son of Will Smith, is the star of films such as The Karate Kid (2010) and After Earth (2013).
Jaden is also known for some of his philosophy that he delivers via Twitter. When writing on Twitter,
he is known for almost always capitalizing every word. For simplicity, you'll have to capitalize each word,
check out how contractions are expected to be in the example below.

Your task is to convert strings to how they would be written by Jaden Smith.
The strings are actual quotes from Jaden Smith, but they are not capitalized in the same way he originally typed them.

Example:

Not Jaden-Cased: "How can mirrors be real if our eyes aren't real"
Jaden-Cased:     "How Can Mirrors Be Real If Our Eyes Aren't Real"
Link to Jaden's former Twitter account @officialjaden via archive.org
*/
func ToJadenCase(str string) string {
	var words []string
	for _, word := range strings.Split(str, " ") {
		wordJadenCased := strings.ToUpper(string(word[0])) + string(word[1:])
		words = append(words, wordJadenCased)
	}
	return strings.Join(words, " ")
}

/*
Mumbling

This time no story, no theory. The examples below show you how to write function accum:

Examples:
accum("abcd") -> "A-Bb-Ccc-Dddd"
accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
accum("cwAt") -> "C-Ww-Aaa-Tttt"
The parameter of accum is a string which includes only letters from a..z and A..Z.
*/
func Accum(s string) (result string) {
	length := 1
	for j, v := range s {
		for i := 0; i < length; i++ {
			if i == 0 {
				result += strings.ToUpper(string(v))
			} else {
				result += strings.ToLower(string(v))
			}
		}
		if j+1 < len(s) {
			result += "-"
		}
		length += 1

	}
	return
}

/*
Middle Character

You are going to be given a word. Your job is to return the middle character of the word.
If the word's length is odd, return the middle character.
If the word's length is even, return the middle 2 characters.

#Examples:

Kata.getMiddle("test") should return "es"

Kata.getMiddle("testing") should return "t"

Kata.getMiddle("middle") should return "dd"

Kata.getMiddle("A") should return "A"

#Input

A word (string) of length 0 < str < 1000 (In javascript you may get slightly more than 1000 in some test cases due to an error in the test cases). You do not need to test for this. This is only here to tell you that you do not need to worry about your solution timing out.

#Output

The middle character(s) of the word represented as a string.
*/
func GetMiddle(s string) (middle string) {
	pos := len(s) / 2
	if len(s)%2 != 0 {
		middle = s[pos : pos+1]
	} else {
		middle = s[pos-1 : pos+1]
	}
	return
}

/*
Highest Lowest

In this little assignment you are given a string of space separated numbers, and have to return the highest and lowest number.

Examples
HighAndLow("1 2 3 4 5")  // return "5 1"
HighAndLow("1 2 -3 4 5") // return "5 -3"
HighAndLow("1 9 3 4 -5") // return "9 -5"
Notes
All numbers are valid Int32, no need to validate them.
There will always be at least one number in the input string.
Output string must be two numbers separated by a single space, and highest number is first.
*/
func HighAndLow(in string) string {
	lowest := math.MaxInt32
	highest := math.MinInt32
	for _, s := range strings.Split(in, " ") {
		intValue, err := strconv.Atoi(s)
		if err == nil {
			if intValue <= lowest {
				lowest = intValue
			}
			if intValue >= highest {
				highest = intValue
			}
		}
	}
	return fmt.Sprint(highest) + " " + fmt.Sprint(lowest)
}

/*
Disemvowel Trolls

Trolls are attacking your comment section!

A common way to deal with this situation is to remove all of the vowels from the trolls' comments, neutralizing the threat.

Your task is to write a function that takes a string and return a new string with all vowels removed.

For example, the string "This website is for losers LOL!" would become "Ths wbst s fr lsrs LL!".

Note: for this kata y isn't considered a vowel.
*/
func Disemvowel(comment string) (disemvoweld string) {
	for _, v := range comment {
		lowerV := strings.ToLower(string(v))
		switch lowerV {
		case "a", "e", "i", "o", "u":
		default:
			disemvoweld = disemvoweld + string(v)
		}
	}
	return disemvoweld
}

/*
Vowel Count

Return the number (count) of vowels in the given string.

We will consider a, e, i, o, u as vowels for this Kata (but not y).

The input string will only consist of lower case letters and/or spaces.
*/
func GetCount(str string) (count int) {
	characters := []rune(str)
	for _, v := range characters {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			count += 1
		}
	}
	return count
}

/*
Sum Of Odd Numbers
Given the triangle of consecutive odd numbers:

	          1
	       3     5
	    7     9    11
	13    15    17    19

21    23    25    27    29
...
Calculate the sum of the numbers in the nth row of this triangle (starting at index 1) e.g.: (Input --> Output)

1 -->  1
2 --> 3 + 5 = 8
*/
func RowSumOddNumbers(n int) int {
	return rowSumOddNumbers(1, n, 1, []int{n})
}

func rowSumOddNumbers(currentLevel, targetLevel, lastOddNumber int, row []int) (rowSum int) {
	if currentLevel == targetLevel {
		for _, v := range row {
			rowSum += v
		}
		return
	} else {
		nextRow := []int{}
		nextOddNumber := lastOddNumber
		for i := 0; i < len(row)+1; i++ {
			nextOddNumber = nextOddNumber + 2
			nextRow = append(nextRow, nextOddNumber)
		}
		return rowSumOddNumbers(currentLevel+1, targetLevel, nextOddNumber, nextRow)
	}
}

/*
Cats and Shelves

Description
An infinite number of shelves are arranged one above the other in a staggered fashion.
The cat can jump either one or three shelves at a time:
from shelf i to shelf i+1 or i+3 (the cat cannot climb on the shelf directly above its head), according to the illustration:

	┌────────┐
	│-6------│
	└────────┘

┌────────┐
│------5-│
└────────┘  ┌─────► OK!

	│    ┌────────┐
	│    │-4------│
	│    └────────┘

┌────────┐  │
│------3-│  │
BANG!────┘  ├─────► OK!

	▲  |\_/|  │    ┌────────┐
	│ ("^-^)  │    │-2------│
	│ )   (   │    └────────┘

┌─┴─┴───┴┬──┘
│------1-│
└────────┘
Input
Start and finish shelf numbers (always positive integers, finish no smaller than start)

Task
Find the minimum number of jumps to go from start to finish

Example
Start 1, finish 5, then answer is 2 (1 => 4 => 5 or 1 => 2 => 5)
*/
func Cats(start, finish int) int {
	var initialPlatforms []int
	initialPlatforms = append(initialPlatforms, start)
	return moveCat(0, finish, initialPlatforms)
}

func moveCat(level, finish int, currentPlatforms []int) int {
	if contains(finish, currentPlatforms) {
		return level
	} else {
		var nextCurrentPlatforms []int
		for _, nextPlatform := range currentPlatforms {
			if nextPlatform+1 <= finish && !contains(nextPlatform+1, nextCurrentPlatforms) {
				nextCurrentPlatforms = append(nextCurrentPlatforms, nextPlatform+1)
			}
			if nextPlatform+3 <= finish && !contains(nextPlatform+3, nextCurrentPlatforms) {
				nextCurrentPlatforms = append(nextCurrentPlatforms, nextPlatform+3)
			}
		}
		return moveCat(level+1, finish, nextCurrentPlatforms)
	}
}

func contains(platform int, currentPlatforms []int) bool {
	if currentPlatforms == nil {
		return false
	}
	for _, nextPlatform := range currentPlatforms {
		if nextPlatform == platform {
			return true
		}
	}
	return false
}

/*
Factorial

Your task is to write function factorial.

https://en.wikipedia.org/wiki/Factorial
*/
func Factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}

/*
Another Card Game
Twelve cards with grades from 0 to 11 randomly divided among 3 players: Frank, Sam, and Tom, 4 cards each. The game consists of 4 rounds.
The goal of the round is to move by the card with the most points.

In round 1, the first player who has a card with 0 points, takes the first turn, and he starts with that card.

Then the second player (queue - Frank -> Sam -> Tom -> Frank, etc.) can move with any of his cards
(each card is used only once per game, and there are no rules that require players to make only the best moves).

The third player makes his move after the second player, and he sees the previous moves.

The winner of the previous round then makes the first move in the next round with any remaining card.

The player who wins 2 rounds first, wins the game.

Task
Return true if Frank has a chance of winning the game.
Return false if Frank has no chance.

Input
3 arrays of 4 unique numbers in each (numbers in array are sorted in ascending order). Input is always valid, no need to check.

Example
Round 1: Frank 2 5 8 11, Sam 1 4 7 10, Tom 0 3 6 9.

	Tom has to start from 0.
	Frank is risking nothing and goes 2.
	Sam gets lucky and wins round by throwing 4.

Round 2: Frank 5 8 11, Sam 1 7 10, Tom 3 6 9.

	Sam starts from 1.
	Tom goes 3,
	Frank wins with 5.

Round 3: Frank 8 11, Sam 7 10, Tom 6 9.

	Frank starts from 11 and wins the round either way.

Frank is the first to win 2 rounds and therefore wins the game, the answer is true.

One more example
Frank 0 1 2 3, Sam 6 7 8 11, Tom 4 5 9 10.

	With these cards Frank has no chance, the answer is false.

Tip
Players can actually play DUMB moves, especially Sam and Tom.
*/
func Game(frank, sam, tom [4]int) bool {
	beatingHandSam := beatingHand(frank, sam)
	beatingHandTom := beatingHand(frank, tom)
	return len(beatingHandSam) >= 2 && len(beatingHandTom) >= 2
}

func beatingHand(player1, player2 [4]int) (result []int) {
	player2StartIndex := 0
	for _, player1Card := range player1 {
		for j, player2Card := range player2 {
			if j < player2StartIndex {
				continue
			}
			if j >= player2StartIndex && player1Card > player2Card {
				result = append(result, player1Card)
				player2StartIndex += 1
				break
			}
		}
	}
	return
}

/*
Greet: Returning Strings

Make a function that will return a greeting statement that uses an input; your program should return, "Hello, <name> how are you doing today?".

[Make sure you type the exact thing I wrote or the program may not execute properly]
*/
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}

// 8 kyu -----------------------------------------------------------------------------------------

/*
Smallest Integer

Given an array of integers your solution should find the smallest integer.

For example:

Given [34, 15, 88, 2] your solution will return 2
Given [34, -345, -1, 100] your solution will return -345

You can assume, for the purpose of this kata, that the supplied array will not be empty.
*/
func SmallestIntegerFinder(numbers []int) (smallest int) {
	for i, v := range numbers {
		if i == 0 || v < smallest {
			smallest = v
		}
	}
	return
}

/*
Convert string to array

Write a function to split a string and convert it into an array of words.

Examples (Input ==> Output):
"Robin Singh" ==> ["Robin", "Singh"]

"I love arrays they are my favorite" ==> ["I", "love", "arrays", "they", "are", "my", "favorite"]
*/
func StringToArray(str string) []string {
	return strings.Split(str, " ")
}

/*
XOR

Exclusive "or" (xor) Logical Operator
Overview
In some scripting languages like PHP, there exists a logical operator (e.g. &&, ||, and, or, etc.) called the "Exclusive Or" (hence the name of this Kata).
The exclusive or evaluates two booleans. It then returns true if exactly one of the two expressions are true, false otherwise. For example:

false xor false == false // since both are false
true xor false == true // exactly one of the two expressions are true
false xor true == true // exactly one of the two expressions are true
true xor true == false // Both are true.  "xor" only returns true if EXACTLY one of the two expressions evaluate to true.
Task
Since we cannot define keywords in Javascript (well, at least I don't know how to do it),
your task is to define a function xor(a, b) where a and b are the two expressions to be evaluated.
Your xor function should have the behaviour described above, returning true if exactly one of the two expressions evaluate to true, false otherwise.
*/
func Xor(a, b bool) bool {
	return (a && !b) || (!a && b)
}

/*
Opposite Number

Very simple, given an integer or a floating-point number, find its opposite.

Examples:

1: -1
14: -14
-34: 34
*/
func Opposite(value int) int {
	return -value
}

/*
PositiveSum
You get an array of numbers, return the sum of all of the positives ones.

Example [1,-4,7,12] => 1 + 7 + 12 = 20

Note: if there is nothing to sum, the sum is default to 0.
*/
func PositiveSum(numbers []int) (sum int) {
	for _, n := range numbers {
		if n > 0 {
			sum += n
		}
	}
	return
}

/*
Rock Paper Scissors
Let's play! You have to return which player won! In case of a draw return Draw!.

Examples(Input1, Input2 --> Output):

"scissors", "paper" --> "Player 1 won!"
"scissors", "rock" --> "Player 2 won!"
"paper", "paper" --> "Draw!"
*/
func Rps(p1, p2 string) (result string) {
	if p1 == p2 {
		result = "Draw!"
	} else if (p1 == "scissors" && p2 == "paper") || (p1 == "paper" && p2 == "rock") || (p1 == "rock" && p2 == "scissors") {
		result = "Player 1 won!"
	} else {
		result = "Player 2 won!"
	}
	return
}

/*
Given a set of numbers, return the additive inverse of each. Each positive becomes negatives, and the negatives become positives.

invert([1,2,3,4,5]) == [-1,-2,-3,-4,-5]
invert([1,-2,3,-4,5]) == [-1,2,-3,4,-5]
invert([]) == []

You can assume that all values are integers. Do not mutate the input array/list.
*/
func Invert(arr []int) (result []int) {
	for _, nr := range arr {
		result = append(result, -1*nr)
	}
	return
}

/*
Timmy & Sarah think they are in love, but around where they live, they will only know once they pick a flower each.
If one of the flowers has an even number of petals and the other has an odd number of petals it means they are in love.

Write a function that will take the number of petals of each flower and return true if they are in love and false if they aren't.
*/
func LoveFunc(flower1, flower2 int) bool {
	return flower1%2 != flower2%2
}

/*
Create a function that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.
*/
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
