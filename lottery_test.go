package lottery

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type DropItem struct {
	ItemID   int
	ItemName string
	DropProb int
}

func (d DropItem) Prob() int {
	return d.DropProb
}

var _ Interface = (*DropItem)(nil)

func TestLots(t *testing.T) {
	l := New(rand.New(rand.NewSource(time.Now().UnixNano())))

	dropItems := []Interface{
		DropItem{ItemID: 1, ItemName: "エリクサ", DropProb: 10},
		DropItem{ItemID: 2, ItemName: "エーテル", DropProb: 20},
		DropItem{ItemID: 3, ItemName: "ポーション", DropProb: 30},
		DropItem{ItemID: 4, ItemName: "ハズレ", DropProb: 40},
	}

	check := 2000000
	countMap := map[DropItem]int{}
	for i := 0; i < check; i++ {
		lotIdx := l.Lots(dropItems...)
		if lotIdx == -1 {
			t.Fatal("lot error")
		}

		switch d := dropItems[lotIdx].(type) {
		case DropItem:
			countMap[d]++
		}
	}

	for dropItem, count := range countMap {
		result := float64(count) / float64(check) * 100
		prob := float64(dropItem.Prob())
		// 誤差0.1チェック
		if (prob-0.1) <= result && result < (prob+0.1) {
			fmt.Printf("ok %3.5f%%(%7d) : %s\n", result, count, dropItem.ItemName)
		} else {
			t.Errorf("error %3.5f%%(%7d) : %s\n", result, count, dropItem.ItemName)
		}
	}
}

func TestLot(t *testing.T) {
	l := New(rand.New(rand.NewSource(time.Now().UnixNano())))

	check := 1000000
	prob := float64(4.0) // 4%
	count := 0
	for i := 0; i < check; i++ {
		if l.Lot(int(prob)) {
			count++
		}
	}
	result := float64(count) / float64(check) * 100

	// 誤差0.1チェック
	if (prob-0.1) <= result && result < (prob+0.1) {
		fmt.Printf("lottery ok %f%%\n", result)
	} else {
		t.Errorf("lottery error %f%%", result)
	}
}

func TestLotOf(t *testing.T) {
	l := New(rand.New(rand.NewSource(time.Now().UnixNano())))

	check := 1000000
	prob := float64(0.5) // 0.5%
	count := 0
	for i := 0; i < check; i++ {
		// 1万分率で計算
		if l.LotOf(int(prob/100*10000), 10000) {
			count++
		}
	}
	result := float64(count) / float64(check) * 100

	// 誤差0.1チェック
	if (prob-0.1) <= result && result < (prob+0.1) {
		fmt.Printf("lottery ok %f%%\n", result)
	} else {
		t.Errorf("lottery error %f%%", result)
	}
}

func TestLot_0to100(t *testing.T) {
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
