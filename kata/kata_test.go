package kata_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"codewars-go/kata"
)

var _ = Describe("Kata RockPaperScissors", func() {
	It("works as expected", func() {
		Expect(kata.Rps("rock", "scissors")).To(Equal("Player 1 won!"))
		Expect(kata.Rps("scissors", "rock")).To(Equal("Player 2 won!"))
		Expect(kata.Rps("rock", "rock")).To(Equal("Draw!"))
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
	actual := kata.Invert(append([]int{}, a...))
	if len(expected) == 0 {
		Expect(actual).To(BeEmpty(), "With arr = %v", a)
	} else {
		Expect(actual).To(Equal(expected), "With arr = %v", a)
	}
}

var _ = Describe("Kata EvenOrOdd", func() {
	It("works as expected", func() {
		Expect(kata.EvenOrOdd(0)).To(Equal("Even"))
		Expect(kata.EvenOrOdd(1)).To(Equal("Odd"))
		Expect(kata.EvenOrOdd(2)).To(Equal("Even"))
		Expect(kata.EvenOrOdd(-1)).To(Equal("Odd"))
		Expect(kata.EvenOrOdd(-2)).To(Equal("Even"))
	})
})

var _ = Describe("Kata OppositesAttract", func() {
	It("works as expected", func() {
		Expect(kata.LoveFunc(1, 4)).To(Equal(true))
		Expect(kata.LoveFunc(2, 2)).To(Equal(false))
		Expect(kata.LoveFunc(0, 1)).To(Equal(true))
		Expect(kata.LoveFunc(0, 0)).To(Equal(false))
	})
})
