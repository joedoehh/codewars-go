package gokyu5_test

import (
	"codewars-go/kata/gokyu5"
	"fmt"
	"math"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoKyu5(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoKyu5 Suite")
}

// count ip adresses ----------------
var _ = Describe("Test IP4 adresses", func() {

	It("should handle basic cases", func() {
		Expect(gokyu5.IpsBetween("150.0.0.0", "150.0.0.1")).To(Equal(1))
		Expect(gokyu5.IpsBetween("10.0.0.0", "10.0.0.50")).To(Equal(50))
		Expect(gokyu5.IpsBetween("20.0.0.10", "20.0.1.0")).To(Equal(246))
		Expect(gokyu5.IpsBetween("150.0.0.0", "150.0.0.0")).To(Equal(0))
	})
})

// int32 to IPv4 ----------------

var _ = Describe("Tests IPv4", func() {
	It("Sample tests", func() {
		Expect(gokyu5.Int32ToIp(math.MaxUint32)).To(Equal("255.255.255.255"))
		Expect(gokyu5.Int32ToIp(2149583361)).To(Equal("128.32.10.1"))
		Expect(gokyu5.Int32ToIp(2154959208)).To(Equal("128.114.17.104"))
		Expect(gokyu5.Int32ToIp(0)).To(Equal("0.0.0.0"))

	})
})

// last digit of a large number ----------------

var _ = Describe("Tests Large Number", func() {
	It("Sample tests", func() {
		Expect(gokyu5.LastDigit("4", "1")).To(Equal(4))
		Expect(gokyu5.LastDigit("4", "2")).To(Equal(6))
		Expect(gokyu5.LastDigit("9", "7")).To(Equal(9))
		Expect(gokyu5.LastDigit("10", "10000000000")).To(Equal(0))
		Expect(gokyu5.LastDigit("1606938044258990275541962092341162602522202993782792835301376", "2037035976334486086268445688409378161051468393665936250636140449354381299763336706183397376")).To(Equal(6))
		Expect(gokyu5.LastDigit("3715290469715693021198967285016729344580685479654510946723", "68819615221552997273737174557165657483427362207517952651")).To(Equal(7))
	})
})

// List Squared ----------------

func dotestSquared(m, n int, exp [][]int) {
	var ans = gokyu5.ListSquared(m, n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("List Squared", func() {

	It("should handle basic cases", func() {
		dotestSquared(1, 250, [][]int{{1, 1}, {42, 2500}, {246, 84100}})
		dotestSquared(250, 500, [][]int{{287, 84100}})
		dotestSquared(300, 600, [][]int{})
	})
})

// Weight For Weight ----------------

func dotestOrderWeight(s string, exp string) {
	var ans = gokyu5.OrderWeight(s)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests OrderWeight", func() {

	It("should handle basic cases", func() {
		dotestOrderWeight("103 123 4444 99 2000", "2000 103 123 4444 99")
		dotestOrderWeight("2000 10003 1234000 44444444 9999 11 11 22 123", "11 11 2000 10003 22 123 1234000 44444444 9999")
		dotestOrderWeight("", "")
	})

})

// Direction Reduction ----------------

func dotestDirReduc(arr []string, exp []string) {
	var ans = gokyu5.DirReduc(arr)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests DirReduc", func() {

	It("should handle basic cases", func() {
		var a = []string{"NORTH", "SOUTH", "EAST", "WEST"}
		dotestDirReduc(a, []string{})
		a = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}
		dotestDirReduc(a, []string{"WEST"})
		a = []string{"NORTH", "WEST", "SOUTH", "EAST"}
		dotestDirReduc(a, []string{"NORTH", "WEST", "SOUTH", "EAST"})
		a = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}
		dotestDirReduc(a, []string{"NORTH"})
		a = []string{"WEST", "WEST", "EAST", "EAST", "WEST", "SOUTH", "NORTH", "SOUTH"}
		dotestDirReduc(a, []string{"WEST", "SOUTH"})
	})
})

// Best Travel ----------------

func dotestTravel(t, k int, ls []int, exp int) {
	var ans = gokyu5.ChooseBestSum(t, k, ls)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Best Travel", func() {

	It("should handle basic cases", func() {
		var ts = []int{50, 55, 56, 57, 58}
		dotestTravel(163, 3, ts, 163)
		ts = []int{50}
		dotestTravel(163, 3, ts, -1)
		ts = []int{91, 74, 73, 85, 73, 81, 87}
		dotestTravel(230, 3, ts, 228)
		dotestTravel(331, 2, ts, 178)
		dotestTravel(331, 4, ts, 331)
		dotestTravel(331, 5, ts, -1)
	})
})

// Peaks ----------------

func dotestPeaks(array []int, expected gokyu5.PosPeaks) {
	var ans = gokyu5.PickPeaks(array)
	if len(expected.Pos) == 0 {
		if len(ans.Pos) != 0 || len(ans.Peaks) != 0 {
			Expect(ans).To(Equal(expected))
		} else {
			Expect(true).To(BeTrue())
		}
	} else {
		Expect(ans).To(Equal(expected))
	}
}

var _ = Describe("Test Example", func() {

	It("should support finding peaks, despite the plateau", func() {
		dotestPeaks(
			[]int{2, 1, 3, 2, 2, 2, 2, 1},
			gokyu5.PosPeaks{Pos: []int{2}, Peaks: []int{3}},
		)
	})

	It("should support finding peaks", func() {
		dotestPeaks(
			[]int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1},
			gokyu5.PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		)
	})

	It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
		dotestPeaks(
			[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3},
			gokyu5.PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		)
	})

	It("should support finding peaks; if the peak is a plateau, it should only return the position of the first element of the plateau", func() {
		dotestPeaks(
			[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1},
			gokyu5.PosPeaks{Pos: []int{3, 7, 10}, Peaks: []int{6, 3, 2}},
		)
	})

	It("should support finding peaks; if the peak is a plateau, it should only return the position of the first element of the plateau", func() {
		dotestPeaks(
			[]int{2, 1, 3, 1, 2, 2, 2, 2, 1},
			gokyu5.PosPeaks{Pos: []int{2, 4}, Peaks: []int{3, 2}},
		)
	})

	It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
		dotestPeaks(
			[]int{2, 1, 3, 1, 2, 2, 2, 2},
			gokyu5.PosPeaks{Pos: []int{2}, Peaks: []int{3}},
		)
	})
})

// Coding squared String ----------------

func dotestCode(a1 string, exp string) {
	var ans = gokyu5.Code(a1)
	Expect(ans).To(Equal(exp))
}
func dotestDecode(a1 string, exp string) {
	var ans = gokyu5.Decode(a1)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Squared Code, Decode", func() {

	It("should handle basic cases", func() {
		var d = "abcdefg"
		var s = "gda\n\veb\n\vfc"
		dotestCode(d, s)
		dotestDecode(s, d)

		d = "I.was.going.fishing.that.morning.at.ten.o'clock"
		s = "c.nhsoI\nltiahi.\noentinw\ncng.nga\nk..mg.s\n\voao.f.\n\v'trtig"
		dotestCode(d, s)
		dotestDecode(s, d)

		d = "Process terminated with status 0 (0 minute(s), 6 second(s))"
		s = "s t setP\n)se(tder\n)e(0a ro\n\vcs twmc\n\vo)muiie\n\vn,istns\n\vd n has\n\v(6u0 t "
		dotestCode(d, s)
		dotestDecode(s, d)

		d = ""
		s = ""
		dotestCode(d, s)
		dotestDecode(s, d)

	})
})

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
