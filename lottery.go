// Package lottery math/randベースの抽選ライブラリ
package lottery

import (
	"math/rand"
	"sort"
)

// Lottery math/randのwrapper
type Lottery struct {
	rd *rand.Rand
}

// Interface 複数の抽選対象を扱う際のインターフェースを提供します
type Interface interface {
	Prob() int
}

// lotterySort 確率が低い順に並び替えるsortインタフェースの実装
type lotterySort []Interface

func (s lotterySort) Len() int           { return len(s) }
func (s lotterySort) Less(i, j int) bool { return s[i].Prob() < s[j].Prob() }
func (s lotterySort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// New return lottery library
func New(rd *rand.Rand) Lottery {
	return Lottery{
		rd: rd,
	}
}

// Lot 0〜100で抽選した結果を返します
func (l Lottery) Lot(prob int) bool {
	return l.LotOf(prob, 100)
}

// LotOf 0〜totalProbで抽選した結果を返します
func (l Lottery) LotOf(prob int, totalProb int) bool {
	if prob < 0 {
		return false
	}

	if prob > totalProb {
		return true
	}

	return l.rd.Intn(totalProb)+1 <= prob
}

// Lots lottery.Interfaceを実装した複数の抽選対象から1件抽選した、indexを返します
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
