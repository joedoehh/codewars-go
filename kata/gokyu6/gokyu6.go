package gokyu6

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

/*
Matrix Transpose
6 kyu
https://www.codewars.com/kata/52fba2a9adcd10b34300094c/train/go
*/
func Transpose(matrix [][]int) (result [][]int) {
	xl := len(matrix[0])
	yl := len(matrix)
	result = make([][]int, xl)
	for i := range result {
		result[i] = make([]int, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = matrix[j][i]
		}
	}
	return result
}

/*
Encrypt this!
6 kyu
https://www.codewars.com/kata/5848565e273af816fb000449/train/go
*/
func EncryptThis(text string) string {
	wordsSplit := strings.Split(text, " ")
	wordsEncrypted := []string{}
	for _, word := range wordsSplit {
		wordsEncrypted = append(wordsEncrypted, encryptWord(word))
	}
	return strings.Join(wordsEncrypted, " ")
}

func encryptWord(text string) string {
	runes := []rune(text)
	if len(text) == 0 {
		return ""
	} else if len(runes) == 1 {
		return toAscii(runes[0])
	} else if len(runes) == 2 {
		return toAscii(runes[0]) + string(runes[1])
	} else if len(runes) == 3 {
		return toAscii(runes[0]) + string(runes[2]) + string(runes[1])
	} else {
		return toAscii(runes[0]) + string(runes[len(runes)-1]) + string(runes[2:len(runes)-1]) + string(runes[1])
	}
}
func toAscii(r rune) string {
	return strconv.Itoa(int(r))
}

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
	return u <= 255
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
