package topn_test

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"

	"github.com/dhowden/topn"
)

func TestMaxInts(t *testing.T) {
	tests := []struct {
		in  []int
		n   int
		out []int
	}{
		{
			in:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			n:   1,
			out: []int{10},
		},
		{
			in:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			n:   2,
			out: []int{10, 9},
		},
		{
			in:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			n:   3,
			out: []int{10, 9, 8},
		},
		{
			in:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			n:   4,
			out: []int{10, 9, 8, 7},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := topn.MaxInts(test.in, test.n)
			if !reflect.DeepEqual(got, test.out) {
				t.Errorf(" = %v, expected %v", got, test.out)
			}
		})
	}
}

func BenchmarkSortInts(b *testing.B) {
	n := 100000
	m := 12
	in := rand.Perm(n)

	for i := 0; i < b.N; i++ {
		sort.Ints(in)
		_ = in[:m]
	}
}

func BenchmarkMaxInts(b *testing.B) {
	n := 100000
	m := 12
	in := rand.Perm(n)

	for i := 0; i < b.N; i++ {
		_ = topn.MaxInts(in, m)
	}
}
