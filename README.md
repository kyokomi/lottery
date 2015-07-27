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
