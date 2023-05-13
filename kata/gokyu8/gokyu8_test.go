package gokyu8_test

import (
	"codewars-go/kata/gokyu8"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoKyu5(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoKyu8 Suite")
}

var _ = Describe("Kata Smallest Integer", func() {
	It("should work for sample tests", func() {
		Expect(Expect(gokyu8.SmallestIntegerFinder([]int{34, 15, 88, 2})).To(Equal(2)))
		Expect(Expect(gokyu8.SmallestIntegerFinder([]int{34, -345, -1, 100})).To(Equal(-345)))
	})
})

var _ = Describe("Kata String to Array", func() {
	It("Sample tests", func() {
		Expect(gokyu8.StringToArray("Robin Singh")).To(Equal([]string{"Robin", "Singh"}))
		Expect(gokyu8.StringToArray("CodeWars")).To(Equal([]string{"CodeWars"}))
		Expect(gokyu8.StringToArray("I love arrays they are my favorite")).To(Equal([]string{"I", "love", "arrays", "they", "are", "my", "favorite"}))
		Expect(gokyu8.StringToArray("1 2 3")).To(Equal([]string{"1", "2", "3"}))
	})
})

var _ = Describe("Kata XOR", func() {
	It("basic tests", func() {
		Expect(gokyu8.Xor(false, false)).To(Equal(false))
		Expect(gokyu8.Xor(true, false)).To(Equal(true))
		Expect(gokyu8.Xor(false, true)).To(Equal(true))
		Expect(gokyu8.Xor(true, true)).To(Equal(false))
	})
	It("nested tests", func() {
		Expect(gokyu8.Xor(false, gokyu8.Xor(false, false))).To(Equal(false))
		Expect(gokyu8.Xor(gokyu8.Xor(true, false), false)).To(Equal(true))
		Expect(gokyu8.Xor(gokyu8.Xor(true, true), false)).To(Equal(false))
		Expect(gokyu8.Xor(true, gokyu8.Xor(true, true))).To(Equal(true))
		Expect(gokyu8.Xor(gokyu8.Xor(false, false), gokyu8.Xor(false, false))).To(Equal(false))
		Expect(gokyu8.Xor(gokyu8.Xor(false, false), gokyu8.Xor(false, true))).To(Equal(true))
		Expect(gokyu8.Xor(gokyu8.Xor(true, false), gokyu8.Xor(false, false))).To(Equal(true))
		Expect(gokyu8.Xor(gokyu8.Xor(true, false), gokyu8.Xor(true, false))).To(Equal(false))
		Expect(gokyu8.Xor(gokyu8.Xor(true, true), gokyu8.Xor(true, false))).To(Equal(true))
		Expect(gokyu8.Xor(gokyu8.Xor(true, gokyu8.Xor(true, true)), gokyu8.Xor(gokyu8.Xor(true, true), false))).To(Equal(true))
	})
})

var _ = Describe("Kata Opposite Number", func() {
	It("should invert number", func() {
		Expect(gokyu8.Opposite(1)).To(Equal(-1))
	})
})

var _ = Describe("Kata Positive Sum", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(gokyu8.PositiveSum([]int{1, 2, 3, 4, 5})).To(Equal(15))
		Expect(gokyu8.PositiveSum([]int{1, -2, 3, 4, 5})).To(Equal(13))
		Expect(gokyu8.PositiveSum([]int{})).To(Equal(0))
		Expect(gokyu8.PositiveSum([]int{-1, -2, -3, -4, -5})).To(Equal(0))
		Expect(gokyu8.PositiveSum([]int{-1, 2, 3, 4, -5})).To(Equal(9))
	})
})

var _ = Describe("Kata RockPaperScissors", func() {
	It("works as expected", func() {
		Expect(gokyu8.Rps("rock", "scissors")).To(Equal("Player 1 won!"))
		Expect(gokyu8.Rps("scissors", "rock")).To(Equal("Player 2 won!"))
		Expect(gokyu8.Rps("rock", "rock")).To(Equal("Draw!"))
	})
})

var _ = Describe("Kata InvertValues", func() {
	It("works as expected", func() {
		dotest_invertvalues([]int{1, 2, 3, 4, 5}, []int{-1, -2, -3, -4, -5})
		dotest_invertvalues([]int{1, -2, 3, -4, 5}, []int{-1, 2, -3, 4, -5})
		dotest_invertvalues(nil, nil)
		dotest_invertvalues([]int{0}, []int{0})
	})
})

func dotest_invertvalues(a []int, expected []int) {
	actual := gokyu8.Invert(append([]int{}, a...))
	if len(expected) == 0 {
		Expect(actual).To(BeEmpty(), "With arr = %v", a)
	} else {
		Expect(actual).To(Equal(expected), "With arr = %v", a)
	}
}

var _ = Describe("Kata EvenOrOdd", func() {
	It("works as expected", func() {
		Expect(gokyu8.EvenOrOdd(0)).To(Equal("Even"))
		Expect(gokyu8.EvenOrOdd(1)).To(Equal("Odd"))
		Expect(gokyu8.EvenOrOdd(2)).To(Equal("Even"))
		Expect(gokyu8.EvenOrOdd(-1)).To(Equal("Odd"))
		Expect(gokyu8.EvenOrOdd(-2)).To(Equal("Even"))
	})
})

var _ = Describe("Kata OppositesAttract", func() {
	It("works as expected", func() {
		Expect(gokyu8.LoveFunc(1, 4)).To(Equal(true))
		Expect(gokyu8.LoveFunc(2, 2)).To(Equal(false))
		Expect(gokyu8.LoveFunc(0, 1)).To(Equal(true))
		Expect(gokyu8.LoveFunc(0, 0)).To(Equal(false))
	})
})
