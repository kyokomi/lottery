package lottery

import (
	"math/rand"
)

type Lottery struct {
	rd *rand.Rand
}

type Interface interface {
	Lot() int
}

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

	return l.rd.Intn(100) + 1 <= prob
}