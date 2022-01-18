package quicksort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/go-test/deep"
)

type TestCase struct {
	Name     string
	Sorted   []int
	Unsorted []int
}

func TestQuicksortBasic(t *testing.T) {
	doTestQuicksort(t, []TestCase{
		{
			Name:     "One Element",
			Sorted:   []int{0},
			Unsorted: []int{0},
		},
		{
			Name:     "Different Numbers",
			Sorted:   []int{-67, 34, 81, 91},
			Unsorted: []int{81, 34, -67, 91},
		},
		{
			Name:     "Same Numbers",
			Sorted:   []int{1, 1, 1, 1},
			Unsorted: []int{1, 1, 1, 1},
		},
	})
}

func TestQuicksortGenerated(t *testing.T) {
	cases := []TestCase{}
	const numberOfCases = 1000
	for i := 0; i < numberOfCases; i++ {
		length := rand.Intn(100) + 50
		unsorted := make([]int, length)
		for j := 0; j < len(unsorted); j++ {
			unsorted[j] = rand.Intn(1000) - 500
		}
		testCase := TestCase{
			Name:     fmt.Sprintf("Generated Test Case %v (length %v)", i, length),
			Unsorted: unsorted,
		}
		cases = append(cases, testCase)
	}
	doTestQuicksort(t, cases)
}

func doTestQuicksort(t *testing.T, cases []TestCase) {
	for _, testCase := range cases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.Sorted == nil {
				testCase.Sorted = make([]int, len(testCase.Unsorted))
				copy(testCase.Sorted, testCase.Unsorted)
				sort.Ints(testCase.Sorted)
			}
			diff := deep.Equal(testCase.Sorted, testCase.Unsorted)
			if diff != nil {
				t.Fail()
				for _, msg := range diff {
					t.Log(msg)
				}
			}
		})
	}
}
