package blocks

import (
	"errors"
	"fmt"
	"sync"
)

var elements = [...]rune{'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'}

type blocks struct {
	m    sync.Mutex
	data []int
}

func New() *blocks {
	return &blocks{
		data: make([]int, 80),
	}
}

func (b *blocks) Reset() {
	b.m.Lock()
	defer b.m.Unlock()
	b.data = make([]int, len(b.data))
}

func (b *blocks) SetSize(size int) error {
	if size < 1 {
		return errors.New("size should be a positive number")
	}
	b.m.Lock()
	defer b.m.Unlock()
	tmp := make([]int, size)
	min := minInt(size, len(b.data))
	for i := 1; i <= min; i++ {
		tmp[len(tmp)-i] = b.data[len(b.data)-i]
	}
	b.data = tmp
	return nil
}

func (b *blocks) Add(num int) {
	b.m.Lock()
	defer b.m.Unlock()
	b.data = b.data[1:]
	b.data = append(b.data, num)
}

func (b *blocks) Draw() error {
	_, err := fmt.Printf("\r%s", b)
	return err
}

func (b *blocks) String() string {
	b.m.Lock()
	defer b.m.Unlock()
	output := make([]rune, len(b.data))
	min, max := minMax(b.data)
	if min != max {
		for i, value := range b.data {
			index := ((value - min) * (len(elements) - 1)) / (max - min)
			output[i] = elements[index]
		}
	} else {
		for i := range b.data {
			output[i] = elements[0]
		}
	}
	return string(output)
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minMax(s []int) (int, int) {
	min := s[0]
	max := s[0]
	for _, value := range s {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
