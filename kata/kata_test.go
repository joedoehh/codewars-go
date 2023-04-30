package kata_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"codewars-go/kata"
)

var _ = Describe("Kata Millipede Example tests", func() {
	It("Should return true", func() {
		Expect(kata.Millipede([]string{"excavate", "endure", "screen", "desire", "theater", "excess", "night"})).To(Equal(true))
	})
	It("Should return false", func() {
		Expect(kata.Millipede([]string{"trade", "pole", "view", "grave", "ladder", "mushroom", "president"})).To(Equal(false))
	})
	It("Should return false", func() {
		Expect(kata.Millipede([]string{"extract", "cycle", "trade", "thesis", "effort", "thesis"})).To(Equal(false))
	})
})

var _ = Describe("Kata Millipede Fixed tests", func() {
	It("Five words true", func() {
		Expect(kata.Millipede([]string{"screen", "desire", "theater", "excess", "night"})).To(Equal(true))
	})
	It("One letter words true", func() {
		Expect(kata.Millipede([]string{"a", "b", "v", "z", "x", "r", "e"})).To(Equal(false))
	})
	It("One letter words true", func() {
		Expect(kata.Millipede([]string{"east", "e", "e", "t", "t", "e", "time"})).To(Equal(true))
	})
	It("One more test", func() {
		Expect(kata.Millipede([]string{"no", "dog", "on", "good"})).To(Equal(false))
	})
})

var _ = Describe("Kata Sum Of Odd Numbers", func() {
	It("Testing for 1", func() { Expect(kata.RowSumOddNumbers(1)).To(Equal(1)) })
	It("Testing for 2", func() { Expect(kata.RowSumOddNumbers(2)).To(Equal(8)) })
	It("Testing for 13", func() { Expect(kata.RowSumOddNumbers(13)).To(Equal(2197)) })
	It("Testing for 19", func() { Expect(kata.RowSumOddNumbers(19)).To(Equal(6859)) })
	It("Testing for 41", func() { Expect(kata.RowSumOddNumbers(41)).To(Equal(68921)) })
	It("Testing for 42", func() { Expect(kata.RowSumOddNumbers(42)).To(Equal(74088)) })
	It("Testing for 74", func() { Expect(kata.RowSumOddNumbers(74)).To(Equal(405224)) })
	It("Testing for 86", func() { Expect(kata.RowSumOddNumbers(86)).To(Equal(636056)) })
	It("Testing for 93", func() { Expect(kata.RowSumOddNumbers(93)).To(Equal(804357)) })
	It("Testing for 101", func() { Expect(kata.RowSumOddNumbers(101)).To(Equal(1030301)) })
})

var _ = Describe("Kata Cats and shelves", func() {
	It("Fixed tests", func() {
		Expect(kata.Cats(1, 5)).To(Equal(2))
		Expect(kata.Cats(1, 1)).To(Equal(0))
		Expect(kata.Cats(2, 5)).To(Equal(1))
		Expect(kata.Cats(2, 4)).To(Equal(2), "Mew")
	})
})

var _ = Describe("Kata Factorial", func() {
	It("basic tests", func() {
		tests_arr := [...][2]int{
			{0, 1},
			{1, 1},
			{4, 24},
			{7, 5040},
			{17, 355687428096000},
		}
		for _, v := range tests_arr {
			Expect(kata.Factorial(v[0])).To(Equal(v[1]))
		}
	})
})

var _ = Describe("Kata Another card game", func() {
	It("Example tests", func() {
		Expect(kata.Game([4]int{2, 5, 8, 11}, [4]int{1, 4, 7, 10}, [4]int{0, 3, 6, 9})).To(Equal(true))
		Expect(kata.Game([4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}, [4]int{0, 9, 10, 11})).To(Equal(false))
		Expect(kata.Game([4]int{1, 4, 5, 9}, [4]int{2, 6, 7, 8}, [4]int{0, 3, 10, 11})).To(Equal(true))
		Expect(kata.Game([4]int{1, 3, 5, 6}, [4]int{0, 4, 7, 8}, [4]int{2, 9, 10, 11})).To(Equal(false))
		Expect(kata.Game([4]int{8, 9, 10, 11}, [4]int{0, 4, 5, 7}, [4]int{1, 2, 3, 6})).To(Equal(true))
	})
})

var _ = Describe("Kata Returning Strings", func() {
	It("should return greeting for Ryan", func() {
		Expect(kata.Greet("Ryan")).To(Equal("Hello, Ryan how are you doing today?"))
	})
})

var _ = Describe("Kata Smallest Integer", func() {
	It("should work for sample tests", func() {
		Expect(Expect(kata.SmallestIntegerFinder([]int{34, 15, 88, 2})).To(Equal(2)))
		Expect(Expect(kata.SmallestIntegerFinder([]int{34, -345, -1, 100})).To(Equal(-345)))
	})
})

var _ = Describe("Kata String to Array", func() {
	It("Sample tests", func() {
		Expect(kata.StringToArray("Robin Singh")).To(Equal([]string{"Robin", "Singh"}))
		Expect(kata.StringToArray("CodeWars")).To(Equal([]string{"CodeWars"}))
		Expect(kata.StringToArray("I love arrays they are my favorite")).To(Equal([]string{"I", "love", "arrays", "they", "are", "my", "favorite"}))
		Expect(kata.StringToArray("1 2 3")).To(Equal([]string{"1", "2", "3"}))
	})
})

var _ = Describe("Kata XOR", func() {
	It("basic tests", func() {
		Expect(kata.Xor(false, false)).To(Equal(false))
		Expect(kata.Xor(true, false)).To(Equal(true))
		Expect(kata.Xor(false, true)).To(Equal(true))
		Expect(kata.Xor(true, true)).To(Equal(false))
	})
	It("nested tests", func() {
		Expect(kata.Xor(false, kata.Xor(false, false))).To(Equal(false))
		Expect(kata.Xor(kata.Xor(true, false), false)).To(Equal(true))
		Expect(kata.Xor(kata.Xor(true, true), false)).To(Equal(false))
		Expect(kata.Xor(true, kata.Xor(true, true))).To(Equal(true))
		Expect(kata.Xor(kata.Xor(false, false), kata.Xor(false, false))).To(Equal(false))
		Expect(kata.Xor(kata.Xor(false, false), kata.Xor(false, true))).To(Equal(true))
		Expect(kata.Xor(kata.Xor(true, false), kata.Xor(false, false))).To(Equal(true))
		Expect(kata.Xor(kata.Xor(true, false), kata.Xor(true, false))).To(Equal(false))
		Expect(kata.Xor(kata.Xor(true, true), kata.Xor(true, false))).To(Equal(true))
		Expect(kata.Xor(kata.Xor(true, kata.Xor(true, true)), kata.Xor(kata.Xor(true, true), false))).To(Equal(true))
	})
})

var _ = Describe("Kata Opposite Number", func() {
	It("should invert number", func() {
		Expect(kata.Opposite(1)).To(Equal(-1))
	})
})

var _ = Describe("Kata Positive Sum", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(kata.PositiveSum([]int{1, 2, 3, 4, 5})).To(Equal(15))
		Expect(kata.PositiveSum([]int{1, -2, 3, 4, 5})).To(Equal(13))
		Expect(kata.PositiveSum([]int{})).To(Equal(0))
		Expect(kata.PositiveSum([]int{-1, -2, -3, -4, -5})).To(Equal(0))
		Expect(kata.PositiveSum([]int{-1, 2, 3, 4, -5})).To(Equal(9))
	})
})

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
