package gokyu7_test

import (
	"codewars-go/kata/gokyu7"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoKyu5(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoKyu7 Suite")
}

var _ = Describe("Kata Printer Errors", func() {
	It("Kata Printer Errors", func() {
		Expect(gokyu7.PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz")).To(Equal("3/56"))
		Expect(gokyu7.PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz")).To(Equal("6/60"))
		Expect(gokyu7.PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu")).To(Equal("11/65"))
	})
})

var _ = Describe("Kata NextSquare", func() {
	It("fixed tests", func() {
		Expect(gokyu7.FindNextSquare(int64(121))).To(Equal(int64(144)))
		Expect(gokyu7.FindNextSquare(int64(625))).To(Equal(int64(676)))
		Expect(gokyu7.FindNextSquare(int64(319225))).To(Equal(int64(320356)))
		Expect(gokyu7.FindNextSquare(int64(15241383936))).To(Equal(int64(15241630849)))
		Expect(gokyu7.FindNextSquare(int64(155))).To(Equal(int64(-1)))
	})
})

var _ = Describe("Kata Two To One", func() {
	It("fixed tests", func() {
		Expect(gokyu7.TwoToOne("aretheyhere", "yestheyarehere")).To(Equal("aehrsty"))
		Expect(gokyu7.TwoToOne("loopingisfunbutdangerous", "lessdangerousthancoding")).To(Equal("abcdefghilnoprstu"))
	})
})

var _ = Describe("Kata GetSum", func() {
	It("Sample tests", func() {
		Expect(gokyu7.GetSum(0, 1)).To(Equal(1))
		Expect(gokyu7.GetSum(1, 2)).To(Equal(3))
		Expect(gokyu7.GetSum(5, -1)).To(Equal(14))
		Expect(gokyu7.GetSum(505, 4)).To(Equal(127759))
		Expect(gokyu7.GetSum(321, 123)).To(Equal(44178))
		Expect(gokyu7.GetSum(0, -1)).To(Equal(-1))
		Expect(gokyu7.GetSum(-50, 0)).To(Equal(-1275))
		Expect(gokyu7.GetSum(-1, -5)).To(Equal(-15))
		Expect(gokyu7.GetSum(-5, -5)).To(Equal(-5))
		Expect(gokyu7.GetSum(-505, 4)).To(Equal(-127755))
		Expect(gokyu7.GetSum(-321, 123)).To(Equal(-44055))
		Expect(gokyu7.GetSum(0, 0)).To(Equal(0))
		Expect(gokyu7.GetSum(-5, -1)).To(Equal(-15))
		Expect(gokyu7.GetSum(5, 1)).To(Equal(15))
		Expect(gokyu7.GetSum(-17, -17)).To(Equal(-17))
		Expect(gokyu7.GetSum(17, 17)).To(Equal(17))
	})
})

var _ = Describe("Kata DNA", func() {
	It("Basic Tests", func() {
		Expect(gokyu7.DNAStrand("AAAA")).To(Equal("TTTT"))
		Expect(gokyu7.DNAStrand("ATTGC")).To(Equal("TAACG"))
		Expect(gokyu7.DNAStrand("GTAT")).To(Equal("CATA"))
	})
})

var _ = Describe("Kata Shortest Word", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(gokyu7.FindShort("bitcoin take over the world maybe who knows perhaps")).To(Equal(3))
		Expect(gokyu7.FindShort("turns out random test cases are easier than writing out basic ones")).To(Equal(3))
		Expect(gokyu7.FindShort("Let's travel abroad shall we")).To(Equal(2))
	})
})

var _ = Describe("Kata Jaden Casing", func() {
	It("should work for sample test cases", func() {
		Expect(gokyu7.ToJadenCase("most trees are blue")).To(Equal("Most Trees Are Blue"))
		Expect(gokyu7.ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own.")).To(Equal("All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own."))
		Expect(gokyu7.ToJadenCase("When I die. then you will realize")).To(Equal("When I Die. Then You Will Realize"))
		Expect(gokyu7.ToJadenCase("Jonah Hill is a genius")).To(Equal("Jonah Hill Is A Genius"))
		Expect(gokyu7.ToJadenCase("Dying is mainstream")).To(Equal("Dying Is Mainstream"))
	})
})

var _ = Describe("Kata Mumbling", func() {

	It("should handle basic cases", func() {
		doTestAccum("ZpglnRxqenU", "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu")
		doTestAccum("NyffsGeyylB", "N-Yy-Fff-Ffff-Sssss-Gggggg-Eeeeeee-Yyyyyyyy-Yyyyyyyyy-Llllllllll-Bbbbbbbbbbb")
		doTestAccum("MjtkuBovqrU", "M-Jj-Ttt-Kkkk-Uuuuu-Bbbbbb-Ooooooo-Vvvvvvvv-Qqqqqqqqq-Rrrrrrrrrr-Uuuuuuuuuuu")
		doTestAccum("EvidjUnokmM", "E-Vv-Iii-Dddd-Jjjjj-Uuuuuu-Nnnnnnn-Oooooooo-Kkkkkkkkk-Mmmmmmmmmm-Mmmmmmmmmmm")
		doTestAccum("HbideVbxncC", "H-Bb-Iii-Dddd-Eeeee-Vvvvvv-Bbbbbbb-Xxxxxxxx-Nnnnnnnnn-Cccccccccc-Ccccccccccc")
	})
})

func doTestAccum(s string, exp string) {
	var ans = gokyu7.Accum(s)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Kata GetMiddle", func() {
	It("Tests", func() {
		Expect(gokyu7.GetMiddle("test")).To(Equal("es"))
		Expect(gokyu7.GetMiddle("testing")).To(Equal("t"))
		Expect(gokyu7.GetMiddle("middle")).To(Equal("dd"))
		Expect(gokyu7.GetMiddle("A")).To(Equal("A"))
	})
})

var _ = Describe("Kata High And Low", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(gokyu7.HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")).To(Equal("42 -9"))
	})
	It("test 2", func() {
		Expect(gokyu7.HighAndLow("1 2 3")).To(Equal("3 1"))
	})
})

var _ = Describe("Kata Disemvowel", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(gokyu7.Disemvowel("This website is for losers LOL!")).To(Equal("Ths wbst s fr lsrs LL!"))
	})
})

var _ = Describe("Kata Vowel Count", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(gokyu7.GetCount("abracadabra")).To(Equal(5))
	})
})

var _ = Describe("Kata Sum Of Odd Numbers", func() {
	It("Testing for 1", func() { Expect(gokyu7.RowSumOddNumbers(1)).To(Equal(1)) })
	It("Testing for 2", func() { Expect(gokyu7.RowSumOddNumbers(2)).To(Equal(8)) })
	It("Testing for 13", func() { Expect(gokyu7.RowSumOddNumbers(13)).To(Equal(2197)) })
	It("Testing for 19", func() { Expect(gokyu7.RowSumOddNumbers(19)).To(Equal(6859)) })
	It("Testing for 41", func() { Expect(gokyu7.RowSumOddNumbers(41)).To(Equal(68921)) })
	It("Testing for 42", func() { Expect(gokyu7.RowSumOddNumbers(42)).To(Equal(74088)) })
	It("Testing for 74", func() { Expect(gokyu7.RowSumOddNumbers(74)).To(Equal(405224)) })
	It("Testing for 86", func() { Expect(gokyu7.RowSumOddNumbers(86)).To(Equal(636056)) })
	It("Testing for 93", func() { Expect(gokyu7.RowSumOddNumbers(93)).To(Equal(804357)) })
	It("Testing for 101", func() { Expect(gokyu7.RowSumOddNumbers(101)).To(Equal(1030301)) })
})

var _ = Describe("Kata Cats and shelves", func() {
	It("Fixed tests", func() {
		Expect(gokyu7.Cats(1, 5)).To(Equal(2))
		Expect(gokyu7.Cats(1, 1)).To(Equal(0))
		Expect(gokyu7.Cats(2, 5)).To(Equal(1))
		Expect(gokyu7.Cats(2, 4)).To(Equal(2), "Mew")
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
			Expect(gokyu7.Factorial(v[0])).To(Equal(v[1]))
		}
	})
})

var _ = Describe("Kata Another card game", func() {
	It("Example tests", func() {
		Expect(gokyu7.Game([4]int{2, 5, 8, 11}, [4]int{1, 4, 7, 10}, [4]int{0, 3, 6, 9})).To(Equal(true))
		Expect(gokyu7.Game([4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}, [4]int{0, 9, 10, 11})).To(Equal(false))
		Expect(gokyu7.Game([4]int{1, 4, 5, 9}, [4]int{2, 6, 7, 8}, [4]int{0, 3, 10, 11})).To(Equal(true))
		Expect(gokyu7.Game([4]int{1, 3, 5, 6}, [4]int{0, 4, 7, 8}, [4]int{2, 9, 10, 11})).To(Equal(false))
		Expect(gokyu7.Game([4]int{8, 9, 10, 11}, [4]int{0, 4, 5, 7}, [4]int{1, 2, 3, 6})).To(Equal(true))
	})
})

var _ = Describe("Kata Returning Strings", func() {
	It("should return greeting for Ryan", func() {
		Expect(gokyu7.Greet("Ryan")).To(Equal("Hello, Ryan how are you doing today?"))
	})
})
