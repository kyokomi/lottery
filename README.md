lottery
==============

[![Circle CI](https://circleci.com/gh/kyokomi/lottery.svg?style=svg)](https://circleci.com/gh/kyokomi/lottery)
[![Coverage Status](https://coveralls.io/repos/kyokomi/lottery/badge.svg?branch=master&service=github)](https://coveralls.io/github/kyokomi/lottery?branch=master)
[![GoDoc](https://godoc.org/github.com/kyokomi/lottery?status.svg)](https://godoc.org/github.com/kyokomi/lottery)

lottery is rand base wrapper library

## Install

```
$ go get github.com/kyokomi/lottery
```

## Usage

```
import "github.com/kyokomi/lottery"
```

## Example

### Simple lottery

```go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kyokomi/lottery"
)

func main() {
	l := lottery.New(rand.New(rand.NewSource(time.Now().UnixNano())))

	check := 1000000
	prob := 4
	count := 0
	for i := 0; i < check; i++ {
		if l.Lot(prob) {
			count++
		}
	}
	fmt.Println(float32(count) / float32(check) * 100, "%")
}
```

### Multiple lottery

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
	"log"

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

func main() {
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
			log.Fatalln("lot error")
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
			log.Fatalln("error %3.5f%%(%7d) : %s\n", result, count, name)
		}
	}
}
```

![](https://cloud.githubusercontent.com/assets/1456047/9173541/ddf0ac12-3fb5-11e5-8cc0-8e0e39c078a4.png)

## Auther

[kyokomi](http://github.com/kyokomi)

## License

[MIT](https://github.com/kyokomi/lottery/blob/master/LICENSE)