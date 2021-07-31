package matrix_multiplication

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

var tests = []struct {
	a              [][]int
	b              [][]int
	expectedResult [][]int
}{
	{
		[][]int{
			{1, 2},
			{3, 4},
		},
		[][]int{
			{5, 6},
			{7, 8},
		},
		[][]int{
			{19, 22},
			{43, 50},
		},
	},
	{
		[][]int{
			{1, 2},
			{3, 4},
		},
		[][]int{
			{5, 6, 7},
			{8, 9, 10},
		},
		[][]int{
			{21, 24, 27},
			{47, 54, 61},
		},
	},
}

func TestMul(t *testing.T) {
	for idx, test := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			result := Mul(test.a, test.b)
			if !reflect.DeepEqual(result, test.expectedResult) {
				t.Fatalf("expected: %v, got: %v", test.expectedResult, result)
			}
		})
	}
}

func TestMulRows(t *testing.T) {
	for idx, test := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			result := MulRow(test.a, test.b)
			if !reflect.DeepEqual(result, test.expectedResult) {
				t.Fatalf("expected: %v, got: %v", test.expectedResult, result)
			}
		})
	}
}

func TestStrassen(t *testing.T) {
	for i := 1; i <= 10; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			a, b := randomMatrix(256)
			expectedResult := Mul(a, b)

			result := Strassen(a, b)

			if !reflect.DeepEqual(result, expectedResult) {
				t.Fatalf("expected: %v, got: %v", expectedResult, result)
			}
		})
	}
}

// go test -run=XXX -bench=. -benchtime=20s
// pkg: matrix_multiplication
// cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
// BenchmarkMul1024x1024x1024/Mul-8                       5        4406788215 ns/op
// BenchmarkMul1024x1024x1024/MulRow-8                   13        1663144705 ns/op
// BenchmarkMul1024x1024x1024/Strassen-8                 18        1262086644 ns/op
func BenchmarkMul1024x1024x1024(b *testing.B) {
	benchmarkMul(1024, b)
}

var result [][]int

func benchmarkMul(width int, bench *testing.B) {
	var r [][]int
	a, b := randomMatrix(width)

	bench.Run(
		"Mul",
		func(bench *testing.B) {
			for i := 0; i < bench.N; i++ {
				r = Mul(a, b)
			}
		},
	)

	bench.Run(
		"MulRow",
		func(bench *testing.B) {
			for i := 0; i < bench.N; i++ {
				r = MulRow(a, b)
			}
		},
	)

	bench.Run(
		"Strassen",
		func(bench *testing.B) {
			for i := 0; i < bench.N; i++ {
				r = Strassen(a, b)
			}
		},
	)

	result = r
}

func randomMatrix(width int) ([][]int, [][]int) {
	a := make([][]int, width)
	for i := range a {
		a[i] = make([]int, width)
		for j := range a[i] {
			a[i][j] = rand.Intn(1000)
		}
	}

	b := make([][]int, width)
	for i, _ := range b {
		b[i] = make([]int, width)
		for j := range b[i] {
			b[i][j] = rand.Intn(1000)
		}
	}
	return a, b
}
