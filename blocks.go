// Package blocks implements interactive conversion from numeric data to the block chart.
package blocks

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

var elements = [...]rune{'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'}

type blocks struct {
	m    sync.Mutex
	data []int
}

// New returns an instance of the block type. Size is scaled to the terminal.
func New() *blocks {
	_, columns := terminalSizes()
	return &blocks{
		data: make([]int, columns),
	}
}

func terminalSizes() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Print(err)
		return 0, 80
	}
	rows, columns := 0, 0
	fmt.Sscan(string(out), &rows, &columns)
	return rows, columns
}

// Reset clears the previous given data.
func (b *blocks) Reset() {
	b.m.Lock()
	defer b.m.Unlock()
	b.data = make([]int, len(b.data))
}

// SetSize changes the amount of block elements given to the output.
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

// Add saves the data for the later output and acts as sliding window if the slice is full.
func (b *blocks) Add(num int) {
	b.m.Lock()
	defer b.m.Unlock()
	b.data = b.data[1:]
	b.data = append(b.data, num)
}

// Draw prints the data as block elements and puts the carriage to the beginning of a line.
func (b *blocks) Draw() error {
	_, err := fmt.Printf("\r%s", b)
	return err
}

// String returns a string representation of data as block elements.
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
