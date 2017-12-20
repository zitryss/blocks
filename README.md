# blocks

[![Build Status](https://travis-ci.org/zitryss/blocks.svg?branch=master)](https://travis-ci.org/zitryss/blocks)
[![Coverage](http://gocover.io/_badge/github.com/zitryss/blocks)](https://gocover.io/github.com/zitryss/blocks)
<!---[![Go Report Card](https://goreportcard.com/badge/github.com/zitryss/blocks)](https://goreportcard.com/report/github.com/zitryss/blocks)--->

Package blocks implements interactive conversion from numeric data to the block chart. Inspired by [spark](https://github.com/holman/spark).

![Blocks Preview](https://user-images.githubusercontent.com/2380748/34184269-145fc1a0-e51f-11e7-9bd3-79e9a42528a6.gif)

## Installation
```
$ go get -u github.com/zitryss/blocks
```

## Example

```golang
package main

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/zitryss/blocks"
)

func main() {
	b := blocks.New()
	b.SetSize(40)
	for {
		cpuLoad, _ := cpu.Percent(1*time.Second, false)
		b.Add(int(cpuLoad[0]))
		b.Draw()
	}
}
```
