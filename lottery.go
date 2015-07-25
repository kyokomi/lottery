package lottery

import (
	"math/rand"
	"sort"
)

type Lottery struct {
	rd *rand.Rand
}

type Interface interface {
	Prob() int
}

type lotterySort []Interface

func (s lotterySort) Len() int           { return len(s) }
func (s lotterySort) Less(i, j int) bool { return s[i].Prob() < s[j].Prob() }
func (s lotterySort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func New(rd *rand.Rand) Lottery {
	return Lottery{
		rd: rd,
	}
}

func (l Lottery) Lot(prob int) bool {
	if prob < 0 {
		return false
	}

	if prob > 100 {
		return true
	}

	return l.rd.Intn(100)+1 <= prob
}

func (l Lottery) Lots(lots ...Interface) int {
	probSum := 0
	for _, l := range lots {
		probSum += l.Prob()
	}

	if probSum <= 0 {
		return -1
	}

	randomProbability := l.rd.Intn(probSum) + 1
	tempProbability := 0

	sort.Sort(lotterySort(lots))
	for idx, l := range lots {
		tempProbability += l.Prob()

		if tempProbability >= randomProbability {
			return idx
		}
	}

	return -1
}
