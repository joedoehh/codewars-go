package gokyu5_test

import (
	"codewars-go/kata/gokyu5"
	"fmt"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoKyu5(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoKyu5 Suite")
}

// Fibo Akin   ----------------

func dotestLengthSupUk(n, k int, exp int) {
	var ans = gokyu5.LengthSupUk(n, k)
	Expect(ans).To(Equal(exp))
}
func dotestComp(n int, exp int) {
	var ans = gokyu5.Comp(n)
	Expect(ans).To(Equal(exp))
}

func dotestU(n int, exp int) {
	var ans = gokyu5.U(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Fibo Akin", func() {
	It("U function", func() {
		dotestU(1, 1)
		dotestU(2, 1)
		dotestU(3, 2)
		dotestU(4, 3)
		dotestU(5, 3)
		dotestU(6, 4)
		dotestU(900, 455)
		dotestU(90000, 44337)
	})
	It("should handle basic cases LengthSupUk", func() {
		dotestLengthSupUk(23, 12, 4)
		dotestLengthSupUk(50, 25, 2)
		dotestLengthSupUk(3332, 973, 1391)
	})
	It("should handle basic cases Comp", func() {
		dotestComp(23, 1)
		dotestComp(100, 22)
		dotestComp(200, 63)
		dotestComp(74626, 37128)
		dotestComp(71749, 35692)
	})
})

// Help Granny ----------------

func dotestGranny(arrFriends []string, ftwns map[string]string, h map[string]float64, exp int) {
	var ans = gokyu5.Tour(arrFriends, ftwns, h)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Granny Tour", func() {

	It("should handle basic cases", func() {

		var friends1 = []string{"A1", "A2", "A3", "A4", "A5"}
		var fTowns1 = map[string]string{"A1": "X1", "A2": "X2", "A3": "X3", "A4": "X4"}
		var dist1 = map[string]float64{"X1": 100.0, "X2": 200.0, "X3": 250.0, "X4": 300.0}
		dotestGranny(friends1, fTowns1, dist1, 889)

		friends1 = []string{"B1", "B2"}
		fTowns1 = map[string]string{"B1": "Y1", "B2": "Y2", "B3": "Y3", "B4": "Y4", "B5": "Y5"}
		dist1 = map[string]float64{"Y1": 50.0, "Y2": 70.0, "Y3": 90.0, "Y4": 110.0, "Y5": 150.0}
		dotestGranny(friends1, fTowns1, dist1, 168)

	})
})

// John & Ann ----------------

func dotestJohn(n int, exp []int) {
	var ans = gokyu5.John(n)
	Expect(ans).To(Equal(exp))
}
func dotestAnn(n int, exp []int) {
	var ans = gokyu5.Ann(n)
	Expect(ans).To(Equal(exp))
}
func dotestSumJohn(n int, exp int) {
	var ans = gokyu5.SumJohn(n)
	Expect(ans).To(Equal(exp))
}
func dotestSumAnn(n int, exp int) {
	var ans = gokyu5.SumAnn(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases John", func() {
		dotestJohn(11, []int{0, 0, 1, 2, 2, 3, 4, 4, 5, 6, 6})
	})
	It("should handle basic cases Ann", func() {
		dotestAnn(6, []int{1, 1, 2, 2, 3, 3})
	})
	It("should handle basic cases SumJohn", func() {
		dotestSumJohn(75, 1720)
	})
	It("should handle basic cases SumAnn", func() {
		dotestSumAnn(115, 4070)
	})
})

// Nico Variation ----------------

var _ = Describe("Tests Nico", func() {
	It("Sample tests", func() {
		dotestNico("crazy", "secretinformation", "cseerntiofarmit on  ")
		dotestNico("abc", "abcd", "abcd  ")
		dotestNico("ba", "1234567890", "2143658709")
		dotestNico("a", "message", "message")
		dotestNico("key", "key", "eky")
		dotestNico("abcdefgh", "abcd", "abcd    ")

	})
})

func dotestNico(key, message, expected string) {
	Expect(gokyu5.Nico(key, message)).To(Equal(expected), "With key = \"%s\", message = \"%s\"", key, message)
}

// Find The Smallest ----------------

func dotestSmallest(n int64, exp []int64) {
	var ans = gokyu5.Smallest(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Smallest", func() {

	It("should handle basic cases", func() {
		dotestSmallest(261235, []int64{126235, 2, 0})
		dotestSmallest(209917, []int64{29917, 0, 1})

	})
})

// Buddy Pairs ----------------

func arrayToString(a []int, delim string) string {
	return strings.Join(strings.Split(fmt.Sprint(a), " "), delim)
}
func dotestBuddy(start, limit int, exp string) {
	ans := arrayToString(gokyu5.Buddy(start, limit), " ")
	fmt.Printf("Expected %s\nGot %s\n", exp, ans)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Buddy Pairs", func() {
	It("should handle basic cases buddy", func() {
		dotestBuddy(48, 49, "[48 75]")
		dotestBuddy(1071625, 1103735, "[1081184 1331967]")
		dotestBuddy(57345, 90061, "[62744 75495]")
		dotestBuddy(2693, 7098, "[5775 6128]")
		dotestBuddy(6379, 8275, "[]")
	})

})

// Human Readable Times ----------------

var _ = Describe("Human Readable Times", func() {
	It("Sample tests", func() {
		Expect(gokyu5.HumanReadableTime(0)).To(Equal("00:00:00"))
		Expect(gokyu5.HumanReadableTime(59)).To(Equal("00:00:59"))
		Expect(gokyu5.HumanReadableTime(60)).To(Equal("00:01:00"))
		Expect(gokyu5.HumanReadableTime(90)).To(Equal("00:01:30"))
		Expect(gokyu5.HumanReadableTime(3599)).To(Equal("00:59:59"))
		Expect(gokyu5.HumanReadableTime(3600)).To(Equal("01:00:00"))
		Expect(gokyu5.HumanReadableTime(45296)).To(Equal("12:34:56"))
		Expect(gokyu5.HumanReadableTime(86399)).To(Equal("23:59:59"))
		Expect(gokyu5.HumanReadableTime(86400)).To(Equal("24:00:00"))
		Expect(gokyu5.HumanReadableTime(359999)).To(Equal("99:59:59"))
	})
})

// Product of consecutive Fib numbers ----------------

var _ = Describe("consecutive fib numbers", func() {

	It("should handle basic cases", func() {
		dotestProductFib(4895, [3]uint64{55, 89, 1})
		dotestProductFib(5895, [3]uint64{89, 144, 0})
		dotestProductFib(74049690, [3]uint64{6765, 10946, 1})
		dotestProductFib(84049690, [3]uint64{10946, 17711, 0})
		dotestProductFib(193864606, [3]uint64{10946, 17711, 1})
	})
})

func dotestProductFib(prod uint64, exp [3]uint64) {
	var ans = gokyu5.ProductFib(prod)
	Expect(ans).To(Equal(exp))
}

// prime in numbers ----------------------------------

var _ = Describe("ConvertFracts", func() {
	It("prime in numbers tests", func() {
		dotestPrimeFactors(7775460, "(2**2)(3**3)(5)(7)(11**2)(17)")
		dotestPrimeFactors(79340, "(2**2)(5)(3967)")

	})
})

func dotestPrimeFactors(n int, exp string) {
	var ans = gokyu5.PrimeFactors(n)
	Expect(ans).To(Equal(exp))
}

// rgb conversion ----------------------------------

var _ = Describe("RGB conversion", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(gokyu5.RGB(0, 0, 0)).To(Equal("000000"))
		Expect(gokyu5.RGB(1, 2, 3)).To(Equal("010203"))
		Expect(gokyu5.RGB(255, 255, 255)).To(Equal("FFFFFF"))
		Expect(gokyu5.RGB(254, 253, 252)).To(Equal("FEFDFC"))
		Expect(gokyu5.RGB(-20, 275, 125)).To(Equal("00FF7D"))
	})
})

// king is in check ----------------------------------
const BASE string = "|---|---|---|---|---|---|---|---|"

func stringifyBoard(board [8][8]rune) string {
	arr := make([]string, 9)
	for i, row := range board {
		arr[i] = fmt.Sprintf("%s\n| %c | %c | %c | %c | %c | %c | %c | %c |", BASE, row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7])
	}
	arr[8] = BASE
	return strings.Join(arr, "\n")
}

func dotest(board [8][8]rune, expected bool) {
	if gokyu5.KingIsInCheck(board) == expected {
		Expect(true).To(BeTrue())
	} else {
		Expect(!expected).To(Equal(expected), "With board\n%s", stringifyBoard(board))
	}
}

var _ = Describe("King Check Tests", func() {
	It("Should work with a check by pawn", func() {
		dotest([8][8]rune{
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', '♟', ' ', ' ', ' ', ' '},
			{' ', ' ', '♔', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		}, true)
	})
	It("Should work with a check by bishop", func() {
		dotest([8][8]rune{
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', '♝'},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{'♔', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		}, true)
	})
	It("Should work with a check by rook", func() {
		dotest([8][8]rune{
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', '♔', ' ', ' ', '♜', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		}, true)
	})
	It("Should work with a check by knight", func() {
		dotest([8][8]rune{
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', '♔', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{'♞', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		}, true)
	})
	It("Should work with a check by queen", func() {
		dotest([8][8]rune{
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', '♛', ' ', '♔', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		}, true)
	})
	It("Should work with a king alone", func() {
		dotest([8][8]rune{
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', '♔', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		}, false)
	})
	It("Should work with no check", func() {
		dotest([8][8]rune{
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{'♛', ' ', ' ', '♞', ' ', ' ', ' ', '♛'},
			{' ', ' ', ' ', '♔', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		}, false)
	})
	It("Should work with a piece blocking another's line of sight", func() {
		dotest([8][8]rune{
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{'♜', ' ', '♝', '♔', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		}, false)
	})
})
