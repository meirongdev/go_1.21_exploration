package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var buf bytes.Buffer
	for i := 0; i < 4; i++ {
		available := buf.Available()
		fmt.Printf("available: %d\n", available)
		b := buf.AvailableBuffer()
		b = strconv.AppendInt(b, int64(i), 10)
		b = append(b, ' ')
		buf.Write(b)
	}
	os.Stdout.Write(buf.Bytes())
	available := buf.Available()
	fmt.Printf("\navailable: %d\n", available)

	// available: 0
	// available: 62
	// available: 60
	// available: 58
	// 0 1 2 3
	// available: 56
}
