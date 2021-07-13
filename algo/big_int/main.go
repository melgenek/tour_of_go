package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//a, _ := new(big.Int).SetString("3141592653589793238462643383279502884197169399375105820974944592", 0)
	//b, _ := new(big.Int).SetString("2718281828459045235360287471352662497757247093699959574966967627", 0)
	// 8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184
	//c := new(big.Int).Mul(a, b)
	//
	//fmt.Println(c)

	fmt.Println(karatsuba(BigInt{"3141592653589793238462643383279502884197169399375105820974944592", true}, BigInt{"2718281828459045235360287471352662497757247093699959574966967627", true}))

	assert("1", func() bool {
		return BigInt{"100", true}.sum(BigInt{"100", true}) == BigInt{"200", true}
	})
	assert("2", func() bool {
		return BigInt{"200", false}.sum(BigInt{"100", false}) == BigInt{"300", false}
	})
	assert("3", func() bool {
		return BigInt{"200", true}.sum(BigInt{"100", false}) == BigInt{"100", true}
	})
	assert("4", func() bool {
		return BigInt{"200", false}.sum(BigInt{"100", true}) == BigInt{"100", false}
	})
	assert("5", func() bool {
		return BigInt{"100", true}.sum(BigInt{"200", false}) == BigInt{"100", false}
	})
	assert("6", func() bool {
		return BigInt{"100", false}.sum(BigInt{"200", true}) == BigInt{"100", true}
	})

	assert("7", func() bool {
		return BigInt{"200", true}.subtract(BigInt{"100", true}) == BigInt{"100", true}
	})
	assert("8", func() bool {
		return BigInt{"200", false}.subtract(BigInt{"100", false}) == BigInt{"100", false}
	})
	assert("9", func() bool {
		return BigInt{"200", true}.subtract(BigInt{"100", false}) == BigInt{"300", true}
	})
	assert("10", func() bool {
		return BigInt{"200", false}.subtract(BigInt{"100", true}) == BigInt{"300", false}
	})
	assert("11", func() bool {
		return BigInt{"100", true}.subtract(BigInt{"200", false}) == BigInt{"300", true}
	})
	assert("12", func() bool {
		return BigInt{"100", false}.subtract(BigInt{"200", true}) == BigInt{"300", false}
	})
	assert("13", func() bool {
		return BigInt{"50194912", true}.subtract(BigInt{"8537238", true}) == BigInt{"41657674", true}
	})
	assert("14", func() bool {
		return BigInt{"41657674", true}.subtract(BigInt{"16699468", true}) == BigInt{"24958206", true}
	})
}

func assert(name string, f func() bool) {
	if !f() {
		panic(name)
	}
}

type BigInt struct {
	value    string
	positive bool
}

func (i BigInt) String() string {
	if i.positive {
		return i.value
	} else {
		return "-" + i.value
	}
}

func (i BigInt) length() int {
	return len(i.value)
}

func (i BigInt) toInt() int {
	fmt.Printf("|%v|\n", i)
	res, _ := strconv.Atoi(i.value)
	return res
}

func (i BigInt) shiftLeft(n int) BigInt {
	return BigInt{i.value + strings.Repeat("0", n), i.positive}
}

func (i BigInt) shiftRight(n int) BigInt {
	return BigInt{strings.Repeat("0", n) + i.value, i.positive}
}

func (a BigInt) sum(b BigInt) BigInt {
	if a.positive == b.positive {
		return sum(a, b)
	} else {
		return subtract(a, b)
	}
}

func sum(a BigInt, b BigInt) BigInt {
	maxLength := max(a.length(), b.length())
	paddedA := a.shiftRight(maxLength - a.length())
	paddedB := b.shiftRight(maxLength - b.length())

	buf := make([]uint8, maxLength)
	var carry uint8 = 0

	for i := maxLength - 1; i >= 0; i-- {
		next := (paddedA.value[i] - '0') + (paddedB.value[i] - '0') + carry
		buf[i] = next%10 + '0'
		carry = next / 10
	}

	if carry != 0 {
		return BigInt{string(carry+'0') + string(buf), a.positive}
	} else {
		return BigInt{string(buf), a.positive}
	}
}

func (a BigInt) biggerThan(b BigInt) bool {
	result := false
	if a.length() > b.length() {
		result = true
	} else if a.length() == b.length() {
		for i := 0; i < a.length(); i++ {
			if a.value[i] > b.value[i] {
				result = true
				break
			} else if a.value[i] < b.value[i] {
				break
			}
		}
	}
	return result
}

func (a BigInt) subtract(b BigInt) BigInt {
	if a.positive == b.positive {
		return subtract(a, b)
	} else {
		return sum(a, b)
	}
}

func subtract(a BigInt, b BigInt) BigInt {
	if b.biggerThan(a) {
		a, b = b, a
	}

	maxLength := max(a.length(), b.length())
	paddedA := a.shiftRight(maxLength - a.length())
	paddedB := b.shiftRight(maxLength - b.length())

	buf := make([]uint8, maxLength)
	var carry int8 = 0

	for i := maxLength - 1; i >= 0; i-- {
		next := int8(paddedA.value[i]-'0') - int8(paddedB.value[i]-'0') - carry
		if next >= 0 {
			buf[i] = uint8(next) + '0'
			carry = 0
		} else {
			buf[i] = uint8(next+10) + '0'
			carry = 1
		}
	}

	return BigInt{string(trimLeadingZero(buf)), a.positive}
}

func mult(a BigInt, b BigInt) BigInt {
	if b.biggerThan(a) {
		a, b = b, a
	}

	total := BigInt{"0", a.positive == b.positive}
	for i := b.length() - 1; i >= 0; i-- {
		buf := make([]uint8, a.length())
		var carry uint8 = 0

		for j := a.length() - 1; j >= 0; j-- {
			next := (a.value[j]-'0')*(b.value[i]-'0') + carry
			buf[j] = next%10 + '0'
			carry = next / 10
		}

		shift := b.length() - 1 - i

		if carry != 0 {
			total = total.sum(BigInt{string(carry+'0') + string(buf), total.positive}.shiftLeft(shift))
		} else {
			total = total.sum(BigInt{string(buf), total.positive}.shiftLeft(shift))
		}
	}

	return total
}

func karatsuba(first BigInt, second BigInt) BigInt {
	if first.length() <= 10 || second.length() <= 10 {
		return mult(first, second)
	} else {
		minLength := min(first.length(), second.length())
		half := minLength / 2

		a := BigInt{first.value[:(first.length() - half)], true}
		b := BigInt{first.value[(first.length() - half):], true}

		c := BigInt{second.value[:(second.length() - half)], true}
		d := BigInt{second.value[(second.length() - half):], true}

		p1 := karatsuba(a, c)
		p2 := karatsuba(a.sum(b), c.sum(d))
		p3 := karatsuba(b, d)

		result := p1.shiftLeft(half * 2).sum(p2.subtract(p1).subtract(p3).shiftLeft(half)).sum(p3)

		if first.positive != second.positive {
			return BigInt{result.value, !result.positive}
		} else {
			return result
		}

	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func trimLeadingZero(buf []uint8) []uint8 {
	firstDigitWithValue := 0
	for idx, value := range buf {
		if value != '0' || idx == len(buf)-1 {
			firstDigitWithValue = idx
			break
		}
	}
	return buf[firstDigitWithValue:]
}
