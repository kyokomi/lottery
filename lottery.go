// Package lottery math/rand based lottery library.
package lottery

import (
	"math/rand"
	"sort"
)

//go:generate mockgen -package lottery -source lottery.go -destination lottery_mock.go

// Lottery math/rand wrapper.
type Lottery interface {
	Lot(prob int) bool
	LotOf(prob int, totalProb int) bool
	Lots(lots ...Interface) int
}

type lottery struct {
	rd *rand.Rand
}

// Interface provide an interface to handle multiple lottery object.
type Interface interface {
	Prob() int
}

// lotterySort probability is low order sort interface.
type lotterySort []Interface

func (s lotterySort) Len() int           { return len(s) }
func (s lotterySort) Less(i, j int) bool { return s[i].Prob() < s[j].Prob() }
func (s lotterySort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// New return lottery library.
func New(rd *rand.Rand) Lottery {
	return &lottery{
		rd: rd,
	}
}

// Lot the result of lottery at 0-100 to return.
func (l lottery) Lot(prob int) bool {
	return l.LotOf(prob, 100)
}

// LotOf the result of lottery at specified value to return.
func (l lottery) LotOf(prob int, totalProb int) bool {
	if prob < 0 {
		return false
	}

	if prob > totalProb {
		return true
	}

	return l.rd.Intn(totalProb)+1 <= prob
}

// Lots the result index of One lottery from multiple lottery object to return.
func (l lottery) Lots(lots ...Interface) int {
	probSum := 0
	for _, l := range lots {
		probSum += l.Prob()
	}

	if probSum <= 0 {
		return -1
	}

	randomProbability := l.rd.Intn(probSum) + 1
	tempProbability := 0
	lotIndex := -1

	sort.Sort(lotterySort(lots))

	for idx, l := range lots {
		tempProbability += l.Prob()

		if tempProbability >= randomProbability {
			lotIndex = idx
			break
		}
	}
	return lotIndex
}
