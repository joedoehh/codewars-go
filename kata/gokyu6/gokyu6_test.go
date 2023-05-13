package gokyu6_test

import (
	"codewars-go/kata/gokyu6"
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoKyu5(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoKyu6 Suite")
}

// ----------------------------------

var _ = Describe("Transpose", func() {
	It("Sample tests", func() {
		Expect(gokyu6.Transpose([][]int{{1}})).To(Equal([][]int{{1}}))
		Expect(gokyu6.Transpose([][]int{{1, 2, 3}})).To(Equal([][]int{{1}, {2}, {3}}))
		Expect(gokyu6.Transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})).To(Equal([][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}))
		Expect(gokyu6.Transpose([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}, {0, 1, 0}, {1, 0, 0}})).To(Equal([][]int{{1, 0, 0, 0, 1}, {0, 1, 0, 1, 0}, {0, 0, 1, 0, 0}}))
	})
})

var _ = Describe("Encrypt", func() {
	It("Sample Tests", func() {
		Expect(gokyu6.EncryptThis("")).Should(Equal(""))
		Expect(gokyu6.EncryptThis("A wise old owl lived in an oak")).Should(Equal("65 119esi 111dl 111lw 108dvei 105n 97n 111ka"))
		Expect(gokyu6.EncryptThis("The more he saw the less he spoke")).Should(Equal("84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp"))
		Expect(gokyu6.EncryptThis("The less he spoke the more he heard")).Should(Equal("84eh 108sse 104e 115eokp 116eh 109ero 104e 104dare"))
		Expect(gokyu6.EncryptThis("Why can we not all be like that wise old bird")).Should(Equal("87yh 99na 119e 110to 97ll 98e 108eki 116tah 119esi 111dl 98dri"))
		Expect(gokyu6.EncryptThis("Thank you Piotr for all your help")).Should(Equal("84kanh 121uo 80roti 102ro 97ll 121ruo 104ple"))
	})
})

var _ = Describe("Number Of Pairs", func() {
	It("should pass basic tests", func() {
		Expect(gokyu6.NumberOfPairs([]string{"red", "red"})).To(Equal(1))
		Expect(gokyu6.NumberOfPairs([]string{"red", "green", "blue"})).To(Equal(0))
		Expect(gokyu6.NumberOfPairs([]string{"gray", "black", "purple", "purple", "gray", "black"})).To(Equal(3))
		Expect(gokyu6.NumberOfPairs([]string{})).To(Equal(0))
		Expect(gokyu6.NumberOfPairs([]string{"red", "green", "blue", "blue", "red", "green", "red", "red", "red"})).To(Equal(4))
	})
})

var _ = Describe("DigPow", func() {

	It("should handle basic cases", func() {
		doTestDigPow(89, 1, 1)
		doTestDigPow(92, 1, -1)
	})
})

func doTestDigPow(n, p int, exp int) {
	var ans = gokyu6.DigPow(n, p)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Longest Repition", func() {
	It("it should work with the sample tests", func() {
		Expect(gokyu6.LongestRepetition("aaaabb")).Should(Equal(gokyu6.Result{'a', 4}))
		Expect(gokyu6.LongestRepetition("bbbaaabaaaa")).Should(Equal(gokyu6.Result{'a', 4}))
		Expect(gokyu6.LongestRepetition("cbdeuuu900")).Should(Equal(gokyu6.Result{'u', 3}))
		Expect(gokyu6.LongestRepetition("abbbbb")).Should(Equal(gokyu6.Result{'b', 5}))
		Expect(gokyu6.LongestRepetition("aabb")).Should(Equal(gokyu6.Result{'a', 2}))
		Expect(gokyu6.LongestRepetition("ba")).Should(Equal(gokyu6.Result{'b', 1}))
		Expect(gokyu6.LongestRepetition("")).Should(Equal(gokyu6.Result{}))
	})
})

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		Expect(gokyu6.Bell(1)).To(Equal([]int{1}))
		Expect(gokyu6.Bell(2)).To(Equal([]int{2, 2}))
		Expect(gokyu6.Bell(3)).To(Equal([]int{3, 4, 3}))
		Expect(gokyu6.Bell(4)).To(Equal([]int{4, 6, 6, 4}))
		Expect(gokyu6.Bell(5)).To(Equal([]int{5, 8, 9, 8, 5}))
		Expect(gokyu6.Bell(6)).To(Equal([]int{6, 10, 12, 12, 10, 6}))
		Expect(gokyu6.Bell(7)).To(Equal([]int{7, 12, 15, 16, 15, 12, 7}))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 12.255.56.1 is correct", func() {
		Expect(gokyu6.Is_valid_ip("12.255.56.1")).To(Equal(true))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that '' is uncorrect", func() {
		Expect(gokyu6.Is_valid_ip("")).To(Equal(false))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that abc.def.ghi.jkl is uncorrect", func() {
		Expect(gokyu6.Is_valid_ip("abc.def.ghi.jkl")).To(Equal(false))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 123.456.789.0 is uncorrect", func() {
		Expect(gokyu6.Is_valid_ip("123.456.789.0")).To(Equal(false))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 12.34.56 is uncorrect", func() {
		Expect(gokyu6.Is_valid_ip("12.34.56")).To(Equal(false))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 12.34.56 .1 is uncorrect", func() {
		Expect(gokyu6.Is_valid_ip("12.34.56 .1")).To(Equal(false))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 12.34.56.-1 is uncorrect", func() {
		Expect(gokyu6.Is_valid_ip("12.34.56.-1")).To(Equal(false))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 123.045.067.089 is uncorrect", func() {
		Expect(gokyu6.Is_valid_ip("123.045.067.089")).To(Equal(false))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 127.1.1.0 is correct", func() {
		Expect(gokyu6.Is_valid_ip("127.1.1.0")).To(Equal(true))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 0.0.0.0 is correct", func() {
		Expect(gokyu6.Is_valid_ip("0.0.0.0")).To(Equal(true))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 0.34.82.53 is correct", func() {
		Expect(gokyu6.Is_valid_ip("0.34.82.53")).To(Equal(true))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 192.168.1.300 is uncorrect", func() {
		Expect(gokyu6.Is_valid_ip("192.168.1.300")).To(Equal(false))
	})
})

var _ = Describe("Kata Wave:", func() {
	It("should return the correct values", func() {
		dotestWave(" x yz", []string{" X yz", " x Yz", " x yZ"})
		dotestWave("abc", []string{"Abc", "aBc", "abC"})
		dotestWave("abc", []string{"Abc", "aBc", "abC"})
		dotestWave(" ab  c", []string{" Ab  c", " aB  c", " ab  C"})
		dotestWave("", []string{})
		dotestWave("z", []string{"Z"})
		dotestWave("a a a a a", []string{"A a a a a", "a A a a a", "a a A a a", "a a a A a", "a a a a A"})
		dotestWave("aaaaa", []string{"Aaaaa", "aAaaa", "aaAaa", "aaaAa", "aaaaA"})
		dotestWave("                                                           ", []string{})
		dotestWave(" a  b     c  defghijk l  m no pqrs tuvwx yz     ", []string{" A  b     c  defghijk l  m no pqrs tuvwx yz     ", " a  B     c  defghijk l  m no pqrs tuvwx yz     ", " a  b     C  defghijk l  m no pqrs tuvwx yz     ", " a  b     c  Defghijk l  m no pqrs tuvwx yz     ", " a  b     c  dEfghijk l  m no pqrs tuvwx yz     ", " a  b     c  deFghijk l  m no pqrs tuvwx yz     ", " a  b     c  defGhijk l  m no pqrs tuvwx yz     ", " a  b     c  defgHijk l  m no pqrs tuvwx yz     ", " a  b     c  defghIjk l  m no pqrs tuvwx yz     ", " a  b     c  defghiJk l  m no pqrs tuvwx yz     ", " a  b     c  defghijK l  m no pqrs tuvwx yz     ", " a  b     c  defghijk L  m no pqrs tuvwx yz     ", " a  b     c  defghijk l  M no pqrs tuvwx yz     ", " a  b     c  defghijk l  m No pqrs tuvwx yz     ", " a  b     c  defghijk l  m nO pqrs tuvwx yz     ", " a  b     c  defghijk l  m no Pqrs tuvwx yz     ", " a  b     c  defghijk l  m no pQrs tuvwx yz     ", " a  b     c  defghijk l  m no pqRs tuvwx yz     ", " a  b     c  defghijk l  m no pqrS tuvwx yz     ", " a  b     c  defghijk l  m no pqrs Tuvwx yz     ", " a  b     c  defghijk l  m no pqrs tUvwx yz     ", " a  b     c  defghijk l  m no pqrs tuVwx yz     ", " a  b     c  defghijk l  m no pqrs tuvWx yz     ", " a  b     c  defghijk l  m no pqrs tuvwX yz     ", " a  b     c  defghijk l  m no pqrs tuvwx Yz     ", " a  b     c  defghijk l  m no pqrs tuvwx yZ     "})
	})
})

func dotestWave(s string, expected []string) {
	actual := gokyu6.Wave(s)
	if len(expected) == 0 {
		Expect(actual).To(BeEmpty(), "with words = \"%s\"", s)
	} else {
		Expect(actual).To(Equal(expected), "with words = \"%s\"", s)
	}
}

var _ = Describe("Kata Camel Case", func() {
	t := [...][2]string{
		{"test case", "TestCase"},
		{"camel case method", "CamelCaseMethod"},
		{"say hello ", "SayHello"},
		{" camel case word", "CamelCaseWord"},
		{"", ""},
		{"Test case", "TestCase"},
		{" camel case word", "CamelCaseWord"},
	}

	for _, v := range t {
		It(fmt.Sprintf("Testing input \"%s\"", v[0]), func() { Expect(gokyu6.CamelCase(v[0])).To(Equal(v[1])) })
	}
})

var _ = Describe("Kata Queue", func() {
	It("Sample tests", func() {
		Expect(gokyu6.QueueTime([]int{}, 1)).To(Equal(0))
		Expect(gokyu6.QueueTime([]int{1, 2, 3, 4}, 1)).To(Equal(10))
		Expect(gokyu6.QueueTime([]int{2, 2, 3, 3, 4, 4}, 2)).To(Equal(9))
		Expect(gokyu6.QueueTime([]int{1, 2, 3, 4, 5}, 100)).To(Equal(5))
	})
})

var _ = Describe("Kata Prime", func() {
	It("Basic tests", func() {
		doTestPrime(0, false)
		doTestPrime(1, false)
		doTestPrime(2, true)
		doTestPrime(73, true)
		doTestPrime(75, false)
		doTestPrime(-1, false)
	})

	It("Test prime", func() {
		doTestPrime(3, true)
		doTestPrime(5, true)
		doTestPrime(7, true)
		doTestPrime(41, true)
		doTestPrime(5099, true)
	})

	It("Test not prime", func() {
		doTestPrime(4, false)
		doTestPrime(6, false)
		doTestPrime(8, false)
		doTestPrime(9, false)
		doTestPrime(45, false)
		doTestPrime(-5, false)
		doTestPrime(-8, false)
		doTestPrime(-41, false)
	})
})

func doTestPrime(n int, expected bool) {
	Expect(gokyu6.IsPrime(n)).To(Equal(expected), "With n = %d", n)
}

var _ = Describe("Kata Nato", func() {
	It("Should return a correctly translated string", func() {
		Expect(gokyu6.ToNato("If you can read")).To(Equal("India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta"))
		Expect(gokyu6.ToNato("Did not see that coming")).To(Equal("Delta India Delta November Oscar Tango Sierra Echo Echo Tango Hotel Alfa Tango Charlie Oscar Mike India November Golf"))
		Expect(gokyu6.ToNato("go for it!")).To(Equal("Golf Oscar Foxtrot Oscar Romeo India Tango !"))
	})
})

var _ = Describe("Kata Tower Builder", func() {
	It("Fixed tests", func() {
		Expect(gokyu6.TowerBuilder(1)).To(Equal([]string{"*"}))
		Expect(gokyu6.TowerBuilder(2)).To(Equal([]string{" * ", "***"}))
		Expect(gokyu6.TowerBuilder(3)).To(Equal([]string{"  *  ", " *** ", "*****"}))
	})
})

var _ = Describe("Kata FindUniq", func() {
	It("should work for some basic cases", func() {
		Expect(gokyu6.FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0})).To(Equal(float32(2)))
		Expect(gokyu6.FindUniq([]float32{0, 0, 0.55, 0, 0})).To(Equal(float32(0.55)))
	})
})

var _ = Describe("Kata CamelCase", func() {
	It("should handle basic cases", func() {
		doTestCamelCase("", "")
		doTestCamelCase("The_Stealth_Warrior", "TheStealthWarrior")
		doTestCamelCase("the-Stealth-Warrior", "theStealthWarrior")
	})
})

func doTestCamelCase(str, exp string) {
	fmt.Println("input:", str)
	var ans = gokyu6.ToCamelCase(str)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Kata CountBits", func() {
	It("basic tests", func() {
		Expect(gokyu6.CountBits(0)).To(Equal(0))
		Expect(gokyu6.CountBits(4)).To(Equal(1))
		Expect(gokyu6.CountBits(7)).To(Equal(3))
		Expect(gokyu6.CountBits(9)).To(Equal(2))
		Expect(gokyu6.CountBits(10)).To(Equal(2))
	})
})

var _ = Describe("Kata CreatePhoneNumber", func() {
	It("basic test", func() {
		Expect(gokyu6.CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})).To(Equal("(123) 456-7890"))
	})
})

var _ = Describe("Kata FindOdd", func() {
	arr := [...][]int{
		{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5},
		{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5},
		{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5},
		{10},
		{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1},
		{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10},
	}
	sol := [...]int{5, -1, 5, 10, 10, 1}
	for i, v := range arr {
		It(fmt.Sprintf("Testing input %v", v), func() { Expect(gokyu6.FindOdd(v)).To(Equal(sol[i])) })
	}
})

var _ = Describe("Kata SpinWords", func() {
	It("should test that the solution returns the correct value for single word inputs", func() {
		Expect(gokyu6.SpinWords("Welcome")).To(Equal("emocleW"))
		Expect(gokyu6.SpinWords("to")).To(Equal("to"))
		Expect(gokyu6.SpinWords("CodeWars")).To(Equal("sraWedoC"))
	})
	It("should test that the solution returns the correct value for multiple word outputs", func() {
		Expect(gokyu6.SpinWords("Hey fellow warriors")).To(Equal("Hey wollef sroirraw"))
		Expect(gokyu6.SpinWords("Burgers are my favorite fruit")).To(Equal("sregruB are my etirovaf tiurf"))
		Expect(gokyu6.SpinWords("Pizza is the best vegetable")).To(Equal("azziP is the best elbategev"))
	})
})

var _ = Describe("Kata Multipliers 3 and 5", func() {

	It("should handle basic cases", func() {
		Expect(gokyu6.Multiple3And5(10)).To(Equal(23))
	})
})

var _ = Describe("Kata Millipede Example tests", func() {
	It("Should return true", func() {
		Expect(gokyu6.Millipede([]string{"excavate", "endure", "screen", "desire", "theater", "excess", "night"})).To(Equal(true))
	})
	It("Should return false", func() {
		Expect(gokyu6.Millipede([]string{"trade", "pole", "view", "grave", "ladder", "mushroom", "president"})).To(Equal(false))
	})
	It("Should return false", func() {
		Expect(gokyu6.Millipede([]string{"extract", "cycle", "trade", "thesis", "effort", "thesis"})).To(Equal(false))
	})
})

var _ = Describe("Kata Millipede Fixed tests", func() {
	It("Five words true", func() {
		Expect(gokyu6.Millipede([]string{"screen", "desire", "theater", "excess", "night"})).To(Equal(true))
	})
	It("One letter words true", func() {
		Expect(gokyu6.Millipede([]string{"a", "b", "v", "z", "x", "r", "e"})).To(Equal(false))
	})
	It("One letter words true", func() {
		Expect(gokyu6.Millipede([]string{"east", "e", "e", "t", "t", "e", "time"})).To(Equal(true))
	})
	It("One more test", func() {
		Expect(gokyu6.Millipede([]string{"no", "dog", "on", "good"})).To(Equal(false))
	})
})
