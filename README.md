lottery
==============

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
