package lottery_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/kyokomi/lottery"
)

type DropItem struct {
	ItemName string
	DropProb int
}

func (d DropItem) Prob() int {
	return d.DropProb
}

var _ lottery.Interface = (*DropItem)(nil)

type Trap struct {
	TrapName string
	prob     int
}

func (t Trap) Prob() int {
	return t.prob
}

var _ lottery.Interface = (*Trap)(nil)

func TestLots(t *testing.T) {
	l := lottery.New(rand.New(rand.NewSource(time.Now().UnixNano())))

	dropItems := []lottery.Interface{
		DropItem{ItemName: "エリクサ", DropProb: 5},
		DropItem{ItemName: "エーテル", DropProb: 10},
		DropItem{ItemName: "ポーション", DropProb: 20},
		DropItem{ItemName: "ハズレ", DropProb: 50},
		Trap{TrapName: "地雷", prob: 5},
		Trap{TrapName: "トラバサミ", prob: 10},
	}

	check := 2000000
	countMap := map[lottery.Interface]int{}
	for i := 0; i < check; i++ {
		lotIdx := l.Lots(dropItems...)
		if lotIdx == -1 {
			t.Fatal("lot error")
		}

		switch d := dropItems[lotIdx].(type) {
		case DropItem, Trap:
			countMap[d]++
		}
	}

	for item, count := range countMap {
		result := float64(count) / float64(check) * 100
		prob := float64(item.Prob())

		name := ""
		switch t := item.(type) {
		case DropItem:
			name = t.ItemName
		case Trap:
			name = t.TrapName
		}

		// 0.1 check
		if (prob-0.1) <= result && result < (prob+0.1) {
			fmt.Printf("ok %3.5f%%(%7d) : %s\n", result, count, name)
		} else {
			t.Errorf("error %3.5f%%(%7d) : %s\n", result, count, name)
		}
	}
}

func TestLot(t *testing.T) {
	l := lottery.New(rand.New(rand.NewSource(time.Now().UnixNano())))

	check := 1000000
	prob := float64(4.0) // 4%
	count := 0
	for i := 0; i < check; i++ {
		if l.Lot(int(prob)) {
			count++
		}
	}
	result := float64(count) / float64(check) * 100

	// 0.1 check
	if (prob-0.1) <= result && result < (prob+0.1) {
		fmt.Printf("lottery ok %f%%\n", result)
	} else {
		t.Errorf("lottery error %f%%", result)
	}
}

func TestLotOf(t *testing.T) {
	l := lottery.New(rand.New(rand.NewSource(time.Now().UnixNano())))

	check := 1000000
	prob := float64(0.5) // 0.5%
	count := 0
	for i := 0; i < check; i++ {
		// 10000 minutes rate
		if l.LotOf(int(prob/100*10000), 10000) {
			count++
		}
	}
	result := float64(count) / float64(check) * 100

	// 0.1 check
	if (prob-0.1) <= result && result < (prob+0.1) {
		fmt.Printf("lottery ok %f%%\n", result)
	} else {
		t.Errorf("lottery error %f%%", result)
	}
}

func TestLot_0to100(t *testing.T) {
	l := lottery.New(rand.New(rand.NewSource(time.Now().UnixNano())))

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

func TestLots_error(t *testing.T) {
	l := lottery.New(rand.New(rand.NewSource(time.Now().UnixNano())))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	probMock := lottery.NewMockInterface(ctrl)
	probMock.EXPECT().Prob().Return(0)
	probMock.EXPECT().Prob().Return(0)

	idx := l.Lots(probMock, probMock)
	if idx != -1 {
		t.Errorf("lots idx error %d != %d", idx, -1)
	}
}

func TestLotMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	lotMock := lottery.NewMockLottery(ctrl)
	lotMock.EXPECT().Lot(1).Return(true)
	lotMock.EXPECT().Lot(2).Return(false)

	if !lotMock.Lot(1) {
		t.Errorf("mock error")
	}

	if lotMock.Lot(2) {
		t.Errorf("mock error")
	}
}

func TestLotOfMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	lotMock := lottery.NewMockLottery(ctrl)
	lotMock.EXPECT().LotOf(1, 100).Return(true)
	lotMock.EXPECT().LotOf(2, 100).Return(false)

	if !lotMock.LotOf(1, 100) {
		t.Errorf("mock error")
	}

	if lotMock.LotOf(2, 100) {
		t.Errorf("mock error")
	}
}

func TestLotsMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	probMock1 := lottery.NewMockInterface(ctrl)
	probMock2 := lottery.NewMockInterface(ctrl)

	lotMock := lottery.NewMockLottery(ctrl)
	lotMock.EXPECT().Lots(probMock1, probMock2).Return(1)

	if lotMock.Lots(probMock1, probMock2) != 1 {
		t.Errorf("mock error")
	}
}
