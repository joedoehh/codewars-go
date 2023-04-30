package kata

import (
	"fmt"
	"strings"
)

/*
Millipede of words

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
