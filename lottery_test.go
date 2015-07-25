package lottery

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestLot(t *testing.T) {
	l := New(rand.New(rand.NewSource(time.Now().UnixNano())))

	check := 1000000
	prob := 4 // 4%
	count := 0
	for i := 0; i < check; i++ {
		if l.Lot(prob) {
			count++
		}
	}
	result := float64(count) / float64(check) * 100

	// 約4% check
	if 3.9 <= result && result < 4.1 {
		fmt.Printf("lottery ok %f%%\n", result)
	} else {
		t.Errorf("lottery error %f%%", result)
	}
}

func TestLot_0_100(t *testing.T) {
	l := New(rand.New(rand.NewSource(time.Now().UnixNano())))

	testCases := []struct {
		prob   int
		result bool
	}{
		{prob: 120, result: true},
		{prob: 100, result: true},
		{prob: 0, result: false},
		{prob: -1, result: false},
	}

	for _, testCase := range testCases {
		if l.Lot(testCase.prob) != testCase.result {
			t.Errorf("lottery error not %f%%", testCase.prob)
		}
	}
}
