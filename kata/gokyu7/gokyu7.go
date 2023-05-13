package gokyu7

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

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
	for _, v := range str {
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
