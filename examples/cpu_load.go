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
